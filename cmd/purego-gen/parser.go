package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

type Type string

var (
	// ref: https://github.com/ebitengine/purego/blob/v0.8.0-alpha/func.go#L39-L59
	types = []Type{
		"string", "byte", "rune", "bool", "uintptr",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"struct", "func", "[]T", "[N]T", "*T", "unsafe.Pointer",
		"float32", "float64",
	}
	typesLookup           = make(map[Type]struct{}, len(types))
	unsupportedParamTypes = map[Type]struct{}{
		//"struct": {},
	}
	unsupportedReturnTypes = map[Type]struct{}{
		//"struct": {},
	}
	noGenParamTypes = map[string]struct{}{
		"float32": {},
		"float64": {},
		"struct":  {},
	}
	noGenReturnTypes = map[string]struct{}{
		"float32": {},
		"float64": {},
		"func":    {}, // TODO: could do with purego.NewCallback
		"struct":  {},
	}
	existingPlatforms = []string{"windows", "darwin", "linux", "freebsd"}
)

func init() {
	// Fill types lookup table
	for _, t := range types {
		typesLookup[t] = struct{}{}
	}
}

type Import struct {
	Name string
	Path string
}

type FuncArg struct {
	Name     string
	OrigType string
	Type     Type
}

type Func struct {
	Library    *Library
	Name       string
	Symbol     string
	ParamArgs  []*FuncArg
	ReturnArgs []*FuncArg
	// NeedsRegisterFunc tells if the generated code will need purego.RegisterFunc instead
	// of the syscall declaration (e.g: if floats, structs are used)
	NeedsRegisterFunc bool
}

type state struct {
	// Necessary in order not to re-parse the same directive more than once
	handledComments map[token.Pos]struct{}
	CurrentLibrary  *Library
	NextSymbol      string
}

func (p *Parser) error(n ast.Node, msg string) error {
	return &genError{
		level: "error",
		pos:   p.fset.Position(n.Pos()),
		msg:   msg,
	}
}

func (p *Parser) warn(n ast.Node, msg string) error {
	return &genError{
		level: "warn",
		pos:   p.fset.Position(n.Pos()),
		msg:   msg,
	}
}

func (p *Parser) parseOrigType(t ast.Expr) string {
	switch tt := t.(type) {
	case *ast.StarExpr:
		return "*" + p.parseOrigType(tt.X)
	case *ast.SelectorExpr:
		return p.parseOrigType(tt.X) + "." + tt.Sel.Name
	case *ast.ArrayType:
		if tt.Len == nil {
			return "[]" + p.parseOrigType(tt.Elt)
		}
		return fmt.Sprintf("[%s]", p.parseOrigType(tt.Len)) + p.parseOrigType(tt.Elt)
	case *ast.BasicLit:
		return tt.Value
	case *ast.FuncType:
		ret := "func("
		if tt.Params != nil {
			for i, param := range tt.Params.List {
				_type := p.parseOrigType(param.Type)
				if len(param.Names) > 0 {
					for j, n := range param.Names {
						ret += n.Name + " " + _type
						if j < len(param.Names)-1 {
							ret += ", "
						}
					}
				} else {
					ret += _type
				}
				if i < len(tt.Params.List)-1 {
					ret += ", "
				}
			}
		}
		ret += ")"
		if tt.Results != nil {
			if len(tt.Results.List) > 1 {
				p.errors = append(p.errors, p.error(t, "function as a param must have at most 1 return value"))
				return ""
			}
			if len(tt.Results.List) == 1 {
				ret += " " + p.parseOrigType(tt.Results.List[0].Type)
			}
		}
		return ret
	case *ast.StructType:
		return "struct"
	case *ast.MapType, *ast.ChanType, *ast.InterfaceType:
		p.errors = append(p.errors, p.error(t, "unsupported type"))
		return ""
	case fmt.Stringer:
		return tt.String()
	default:
		p.errors = append(p.errors, p.error(t, "unknown type"))
		return ""
	}
}

func (p *Parser) parseType(t ast.Expr) Type {
	switch tt := t.(type) {
	case *ast.StarExpr:
		return Type("*T")
	case *ast.SelectorExpr:
		return Type("unsafe.Pointer")
	case *ast.ArrayType:
		if tt.Len == nil {
			return Type("[]T")
		}
		return Type("[N]T")
	case *ast.MapType, *ast.ChanType, *ast.InterfaceType:
		p.errors = append(p.errors, p.error(t, fmt.Sprintf("unsupported type: %v", tt)))
		return ""
	case *ast.StructType:
		return Type("struct")
	case *ast.FuncType:
		return Type("func")
	case fmt.Stringer:
		ret := Type(tt.String())
		switch ret {
		case "float32", "float64":
			p.warnings = append(p.warnings, p.warn(t, fmt.Sprintf("unsupported float type: %v", ret)))
		}
		return ret
	default:
		p.errors = append(p.errors, p.error(t, fmt.Sprintf("unknown type: %v", tt)))
		return ""
	}
}

func (p *Parser) parseFuncArgs(fields *ast.FieldList) []*FuncArg {
	if fields == nil {
		return nil
	}

	var args []*FuncArg
	var unqual int
	for _, fl := range fields.List {
		_type := p.parseType(fl.Type)
		if len(p.errors) > 0 {
			return nil
		}
		origType := p.parseOrigType(fl.Type)
		if len(fl.Names) > 0 {
			for i := 0; i < len(fl.Names); i++ {
				args = append(args, &FuncArg{
					Name:     fl.Names[i].Name,
					OrigType: origType,
					Type:     _type,
				})
			}
		} else {
			args = append(args, &FuncArg{
				Name:     fmt.Sprintf("a%d", unqual),
				OrigType: origType,
				Type:     _type,
			})
			unqual++
		}
	}

	return args
}

func (p *Parser) parseFunc(vs *ast.ValueSpec) *Func {
	switch vt := vs.Type.(type) {
	case *ast.FuncType:
		if len(vs.Values) > 0 {
			p.errors = append(p.errors, p.error(vs, "function should not have a value"))
			return nil
		}
		if len(vs.Names) > 1 {
			p.errors = append(p.errors, p.error(vs, "multi function declarations not supported"))
			return nil
		}
		if p.state.CurrentLibrary == nil {
			p.errors = append(p.errors, p.error(vs, "cannot link the function due to no previous 'library' directive"))
			return nil
		}

		fn := &Func{
			Library: p.state.CurrentLibrary,
			Name:    vs.Names[0].Name,
			// Symbol equals to the function name by default
			Symbol: vs.Names[0].Name,
		}
		// Symbol override directive
		if p.state.NextSymbol != "" {
			fn.Symbol = p.state.NextSymbol
			// Clear the symbol for the next function
			p.state.NextSymbol = ""
		}
		// Params
		args := p.parseFuncArgs(vt.Params)
		if len(p.errors) > 0 {
			return nil
		}
		fn.ParamArgs = append(fn.ParamArgs, args...)
		// Return value
		args = p.parseFuncArgs(vt.Results)
		if len(p.errors) > 0 {
			return nil
		}
		if len(args) > 3 {
			p.errors = append(p.errors, p.error(vs, "functions must not have more than 3 return values"))
			return nil
		}

		for _, a := range args {
			if _, ok := unsupportedReturnTypes[Type(a.OrigType)]; ok {
				p.errors = append(p.errors, p.error(vt.Params, fmt.Sprintf("unsupported function return type: %s", a.Type)))
				return nil
			}
		}
		fn.ReturnArgs = append(fn.ReturnArgs, args...)
		return fn
	default:
		return nil
	}
}

// Some comments might not be parsed
// See: https://github.com/golang/go/issues/20744
func (p *Parser) parseDirective(c *ast.Comment) {
	// Skip comment if already handled before
	if _, ok := p.state.handledComments[c.Pos()]; ok {
		return
	}
	p.state.handledComments[c.Pos()] = struct{}{}

	text := strings.TrimLeft(c.Text, " \t/")
	if strings.HasPrefix(text, "puregogen:") {
		cmd, ok := strings.CutPrefix(text, "puregogen:")
		if !ok || cmd == "" {
			p.warnings = append(p.warnings, p.warn(c, "no option defined for the directive (ignored)"))
			return
		}
		args := strings.Split(cmd, " ")
		switch args[0] {
		case "library":
			if len(args) <= 1 {
				p.warnings = append(p.warnings, p.warn(c, "no arguments specified for 'library' directive (ignored)"))
				return
			}
			var alias string
			paths := map[string]string{}
			for _, arg := range args[1:] {
				kv := strings.Split(arg, "=")
				if len(kv) <= 1 {
					p.errors = append(p.errors, p.error(c, "directive argument should be of the format key=value"))
					return
				}
				switch {
				case strings.HasPrefix(kv[0], "path"):
					kargs := strings.Split(kv[0], ":")
					if kargs[0] != "path" {
						p.errors = append(p.errors, p.error(c, "malformated 'path' option for directive 'library': '"+kargs[0]+"'"))
						return
					}
					if len(kargs) > 1 {
						// Set the path for a specific OS
						paths[strings.TrimSpace(kargs[1])] = kv[1]
						if kv[1] == "" {
							p.errors = append(p.errors, p.error(c, "'path' must not be empty for 'library' directive"))
							return
						}
					} else {
						// If no OS specified set it for all OS
						for _, platform := range existingPlatforms {
							paths[platform] = kv[1]
						}
					}
				case kv[0] == "alias":
					alias = kv[1]
				default:
					p.warnings = append(p.warnings, p.warn(c, "unknown option for directive 'library': '"+kv[0]+"' (ignored)"))
				}
			}

			alias = strings.TrimSpace(alias)
			if alias == "" {
				// Build an alias based on the file name
				var path string
				for _, pt := range paths {
					if pt != "" {
						path = pt
						break
					}
				}
				_, file := filepath.Split(path)
				alias = strings.TrimSuffix(file, filepath.Ext(file))
			}
			l, ok := p.libraries[alias]
			// If the library has not been registered yet, add a reference
			if !ok {
				l = &Library{
					Alias:    alias,
					PathByOS: map[string]string{},
				}
				p.libraries[alias] = l
			}
			// Update the paths by os
			for platform, pt := range paths {
				l.PathByOS[platform] = pt
			}
			// Activate the current library
			p.state.CurrentLibrary = l
		case "function":
			if len(args) <= 1 {
				p.warnings = append(p.warnings, p.warn(c, "no arguments specified for 'function' directive (ignored)"))
				return
			}
			for _, arg := range args[1:] {
				kv := strings.Split(arg, "=")
				if len(kv) <= 1 {
					p.errors = append(p.errors, p.error(c, "directive argument should be of the format key=value"))
					return
				}
				switch kv[0] {
				case "symbol":
					// If the current next symbol hasn't been consumed by a function declaration, warn the user
					if p.state.NextSymbol != "" && p.state.NextSymbol != kv[1] {
						p.warnings = append(p.warnings,
							p.warn(c, "symbol redefined but previous symbol has not been consumed by a func decl: '"+p.state.NextSymbol+"'"),
						)
					}
					p.state.NextSymbol = kv[1]
				default:
					p.warnings = append(p.warnings, p.warn(c, "unknown option for directive 'function': "+kv[0]+" (ignored)"))
				}
			}
		}
	}
}

type Parser struct {
	filename   string
	extrafiles []string
	fset       *token.FileSet
	libraries  map[string]*Library
	warnings   []error
	errors     []error

	state state
}

func NewParser(filename string, extrafiles ...string) *Parser {
	return &Parser{
		filename:   filename,
		extrafiles: extrafiles,
		libraries:  map[string]*Library{},

		state: state{
			handledComments: map[token.Pos]struct{}{},
		},
	}
}

func (p *Parser) Parse() (*Generator, error) {
	g := &Generator{
		filename: p.filename,
		types:    map[string]Type{},
	}

	for _, filename := range append([]string{p.filename}, p.extrafiles...) {
		p.fset = token.NewFileSet()
		p.filename = filename
		root, err := parser.ParseFile(p.fset, p.filename, nil, parser.ParseComments)
		if err != nil {
			return nil, errors.New("parser: " + err.Error())
		}
		// Set package
		g.pkg = root.Name.Name

		// Traverse tree
		ast.Inspect(root, func(n ast.Node) bool {
			if err != nil {
				return false
			}
			if n == nil {
				return true
			}
			switch tn := n.(type) {
			case *ast.ImportSpec:
				imp := &Import{
					Path: strings.Trim(tn.Path.Value, "\""),
				}
				if tn.Name != nil {
					imp.Name = tn.Name.Name
				}
				g.imports = append(g.imports, imp)
			case *ast.TypeSpec:
				var _type Type
				_type = p.parseType(tn.Type)
				if len(p.errors) > 0 {
					return false
				}
				g.types[tn.Name.Name] = _type
			case *ast.ValueSpec:
				// If comment is right above a function variable
				if tn.Doc != nil {
					for _, c := range tn.Doc.List {
						p.parseDirective(c)
					}
					if len(p.errors) > 0 {
						return false
					}
				}
				// if comment is inside the var block at a random location
				if tn.Comment != nil {
					for _, c := range tn.Comment.List {
						p.parseDirective(c)
					}
					if len(p.errors) > 0 {
						return false
					}
				}
				fn := p.parseFunc(tn)
				if len(p.errors) > 0 {
					return false
				}
				if fn != nil {
					g.funcs = append(g.funcs, fn)
				}
			case *ast.Comment:
				p.parseDirective(tn)
				if len(p.errors) > 0 {
					return false
				}
			case *ast.CommentGroup:
				for _, c := range tn.List {
					p.parseDirective(c)
				}
				if len(p.errors) > 0 {
					return false
				}
			default:

				// Nope
			}
			return true
		})
	}
	if len(p.errors) > 0 {
		return nil, fmt.Errorf("parser: %d error(s) encountered", len(p.errors))
	}

	// Resolve types definitions recursively
	var conflict = true
	for conflict {
		conflict = false
		for name, _type := range g.types {
			if typeDef, ok := g.types[string(_type)]; ok {
				conflict = true
				g.types[name] = typeDef
			}
		}
	}
	// Mark func as non-generable if types don't allow it (floats, structs)
	for _, fn := range g.funcs {
		for _, arg := range fn.ParamArgs {
			if _, ok := noGenParamTypes[string(arg.OrigType)]; ok {
				fn.NeedsRegisterFunc = true
				break
			}
			if _, ok := noGenParamTypes[string(g.types[string(arg.Type)])]; ok {
				fn.NeedsRegisterFunc = true
				break
			}
		}
		for _, arg := range fn.ReturnArgs {
			if _, ok := noGenReturnTypes[string(arg.OrigType)]; ok {
				fn.NeedsRegisterFunc = true
				break
			}
			if _, ok := noGenReturnTypes[string(g.types[string(arg.Type)])]; ok {
				fn.NeedsRegisterFunc = true
				break
			}
		}
	}

	return g, nil
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) Warnings() []error {
	return p.warnings
}
