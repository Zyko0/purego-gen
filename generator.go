package puregogen

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
)

const (
	puregoQual = "github.com/ebitengine/purego"
)

type Generator struct {
	filename string
	pkg      string
	imports  []*Import
	types    map[string]Type
	funcs    []*Func
	errors   []error

	platforms        map[string]struct{}
	symbolsByLibrary map[*Library][]string
	symbolByFunc     map[string]string
}

func argCallName(arg *FuncArg) string {
	return "_" + arg.Name
}

func (g *Generator) appendArgsConv(codes []jen.Code, arg *FuncArg) []jen.Code {
	_type := string(arg.Type)
	if t, ok := g.types[_type]; ok {
		_type = string(t)
	}
	stmt := jen.Id(argCallName(arg)).Op(":=")
	rname := jen.Id(arg.Name)
	switch _type {
	case "uintptr":
		codes = append(codes, stmt.Add(rname))
	case "*T":
		codes = append(codes, stmt.Uintptr().Parens(
			jen.Qual("unsafe", "Pointer").Parens(rname),
		))
	case "bool":
		codes = append(codes, stmt.Uintptr().Parens(jen.Id("0")))
		codes = append(codes, jen.If(rname).Block(
			jen.Id(argCallName(arg)).Op("=").Id("1"),
		))
	case "[]T", "[N]T":
		codes = append(codes, stmt.Uintptr().Parens(
			jen.Qual("unsafe", "Pointer").Parens(
				jen.Id("&").Add(rname.Index(jen.Id("0"))),
			),
		))
	case "string":
		// Append a null byte to the string if necessary (to convert to a C string)
		codes = append(codes, jen.If(
			jen.Qual("strings", "HasSuffix").Call(
				jen.Id(arg.Name),
				jen.Id("\"\\x00\""),
			).Block(
				jen.Id(arg.Name).Op("+=").Id("\"\\x00\""),
			),
		))
		codes = append(codes, stmt.
			Uintptr().
			Parens(
				jen.Qual("unsafe", "Pointer").Parens(jen.Op("&").
					Index().Byte().
					Parens(jen.Id(arg.Name)).
					Index(jen.Id("0")),
				),
			),
		)
		codes = append(codes, jen.Defer().Qual("runtime", "KeepAlive").Call(jen.Id(argCallName(arg))))
	case "func":
		codes = append(codes, stmt.Add(
			jen.Qual(puregoQual, "NewCallback").Call(rname),
		))
	default:
		switch {
		case strings.HasPrefix(_type, "uint"),
			strings.HasPrefix(_type, "int"),
			_type == "unsafe.Pointer",
			_type == "rune",
			_type == "byte":
			codes = append(codes, stmt.Uintptr().Parens(rname))
		default:
			g.errors = append(g.errors, fmt.Errorf("unsupported type: %s", _type))
		}
	}
	return codes
}

func (g *Generator) appendRetsConv(codes []jen.Code, ret *FuncArg) ([]jen.Code, bool) {
	_type := string(ret.Type)
	if t, ok := g.types[_type]; ok {
		_type = string(t)
	}
	stmt := jen.Id("_" + ret.Name).Op(":=")
	rname := jen.Id(ret.Name)
	switch _type {
	case "uintptr":
		return codes, false
	case "*T":
		codes = append(codes,
			stmt.Parens(
				jen.Id(ret.OrigType),
			).Parens(
				jen.Op("*").Parens(
					jen.Op("*").Qual("unsafe", "Pointer"),
				).Parens(
					jen.Qual("unsafe", "Pointer").Parens(
						jen.Id("&").Add(rname),
					),
				),
			),
		)
	case "bool":
		codes = append(codes, stmt.Id(ret.Name).Op("==").Id("1"))
	case "[]T", "[N]T":
		codes = append(codes, stmt.Uintptr().Parens(
			jen.Qual("unsafe", "Pointer").Parens(
				jen.Id("&").Add(rname.Index(jen.Id("0"))),
			),
		))
	case "string":
		codes = append(codes, jen.If(
			jen.Qual("strings", "HasSuffix").Call(
				jen.Id(ret.Name),
				jen.Id("\"\\x00\""),
			).Block(
				jen.Id(ret.Name).Op("+=").Id("\"\\x00\""),
			),
		))
		codes = append(codes,
			stmt.Id("&").Parens(
				jen.Op("*").Parens(
					jen.Op("*").Index().Id("byte"),
				).Parens(
					jen.Qual("unsafe", "Pointer").Parens(
						jen.Op("&").Id(ret.Name),
					),
				),
			).Index(jen.Id("0")),
		)
		codes = append(codes, jen.Defer().Qual("runtime", "KeepAlive").Call(jen.Id(argCallName(ret))))
	case "func":
		codes = append(codes, stmt.Add(
			jen.Qual(puregoQual, "NewCallBack").Call(rname),
		))
	default:
		switch {
		case strings.HasPrefix(_type, "uint"),
			strings.HasPrefix(_type, "int"),
			_type == "unsafe.Pointer",
			_type == "rune",
			_type == "byte":
			codes = append(codes, stmt.Id(ret.OrigType).Parens(rname))
		default:
			g.errors = append(g.errors, fmt.Errorf("unsupported type: %s", _type))
			return codes, false
		}
	}
	return codes, true
}

type File struct {
	Filename string
	Content  string
}

const (
	implFileSuffix = "_impl"
	libFileSuffix  = "_lib"
)

func symbolVarName(symbol string) string {
	return "_addr_" + symbol
}

func libHndVarName(lib string) string {
	return "_hnd_" + lib
}

func (g *Generator) appendLoaderFiles(files []*File) []*File {
	// TODO: handle multiple files
	/*for _, p := range g.platforms {
		f := jen.NewFile(g.pkg)
		// OS build directive
		f.Commentf("go:build %s", p)
		// Declarations
		f.Var().Block(
			jen.Commentf("Library %s")
		)
	}*/

	return files
}

func outputFilename(filename, suffix string) string {
	d, f := filepath.Split(filename)
	ext := filepath.Ext(f)
	return d + strings.TrimSuffix(f, ext) + suffix + ext
}

func (g *Generator) Generate(linkOpenLib bool) ([]*File, error) {
	var files []*File
	var usedPlatforms []string

	g.platforms = map[string]struct{}{}
	g.symbolsByLibrary = map[*Library][]string{}
	g.symbolByFunc = map[string]string{}

	// Initialize platforms, libraries and symbols
	for _, fn := range g.funcs {
		if _, ok := g.symbolsByLibrary[fn.Library]; !ok {
			g.symbolsByLibrary[fn.Library] = []string{fn.Symbol}
		} else {
			g.symbolsByLibrary[fn.Library] = append(
				g.symbolsByLibrary[fn.Library], fn.Symbol,
			)
		}
		g.symbolByFunc[fn.Name] = fn.Symbol
		for platform := range fn.Library.PathByOS {
			g.platforms[platform] = struct{}{}
		}
	}
	if len(g.platforms) == 0 {
		return nil, fmt.Errorf("generate: no OS specified")
	}
	for p := range g.platforms {
		usedPlatforms = append(usedPlatforms, p)
	}

	// Functions implementations file
	f := jen.NewFile(g.pkg)
	buildComment := "//go:build " + strings.Join(usedPlatforms, " || ")
	f.Comment(buildComment)
	// Imports
	for _, imp := range g.imports {
		if imp.Name != "" {
			f.ImportAlias(imp.Path, imp.Name)
		}
	}
	// Link unexported functions from purego if single file is specified
	linkOpenLib = true
	var initBody jen.Statement
	if linkOpenLib {
		// Linking
		f.Commentf("//go:linkname openLibrary github.com/ebitengine/purego.openLibrary")
		f.Func().Id("openLibrary").Params(jen.Id("name").String()).Params(
			jen.Uintptr(), jen.Error(),
		)
		f.Commentf("//go:linkname loadSymbol github.com/ebitengine/purego.loadSymbol")
		f.Func().Id("loadSymbol").
			Params(
				jen.Id("handle").Uintptr(),
				jen.Id("name").String(),
			).
			Params(
				jen.Uintptr(), jen.Error(),
			)
		// Library handles
		initBody.Add(jen.Var().Err().Error())
		initBody.Add(jen.Var().Id("path").String())
		initBody.Line()
		for l, symbols := range g.symbolsByLibrary {
			initBody.Comment(l.Alias)
			initBody.Add(jen.Switch(jen.Qual("runtime", "GOOS")).Block(
				*jen.Do(func(cases *jen.Statement) {
					for p := range g.platforms {
						cases.Add(jen.Case(jen.Id("\"" + p + "\"")).
							Id("path").Op("=").Id("\"" + l.PathByOS[p] + "\""),
						)
					}
					cases.Add(jen.Default().Add(
						jen.Panic(
							jen.Id("\"os not supported: \"").
								Op("+").
								Qual("runtime", "GOOS"),
						),
					))
				})...,
			))

			initBody.Add(jen.List(
				jen.Id(libHndVarName(l.Alias)),
				jen.Id("err"),
			).Op("=").Id("openLibrary").Call(jen.Id("path")))
			initBody.If(jen.Err().Op("!=").Nil().Block(
				jen.Panic(jen.Id("\"cannot openLibrary: \"").Op("+").Id("path")),
			))
			// Symbols handles
			initBody.Commentf("Symbols %s", l.Alias)
			for _, symbol := range symbols {
				initBody.Add(jen.List(
					jen.Id(symbolVarName(symbol)),
					jen.Id("err"),
				).Op("=").Id("loadSymbol").Call(
					jen.Id(libHndVarName(l.Alias)),
					jen.Id("\""+symbol+"\"")),
				)
				initBody.If(jen.Err().Op("!=").Nil().Block(
					jen.Panic(jen.Id("\"cannot loadSymbol: " + symbol + "\"")),
				))
			}
			initBody.Line()
		}
	}
	// Library and symbols pointers
	f.Var().Defs(
		jen.Comment("Library handles"),
		jen.Do(func(s *jen.Statement) {
			for l := range g.symbolsByLibrary {
				s.Id(libHndVarName(l.Alias)).Uintptr()
			}
		}),
		jen.Comment("Symbols"),
		jen.Do(func(s *jen.Statement) {
			for l, symbols := range g.symbolsByLibrary {
				s.Comment(l.Alias).Line()
				for _, symbol := range symbols {
					s.Id(symbolVarName(symbol)).Uintptr().Line()
				}
			}
		}),
	)
	// Init
	var funcs []jen.Code
	for _, fn := range g.funcs {
		var params []jen.Code
		var callParams []jen.Code
		var returnTypes []jen.Code
		var funcBody []jen.Code
		// Arguments conversions to uintptr
		callParams = append(callParams, jen.Id(symbolVarName(fn.Symbol)))
		for _, arg := range fn.ParamArgs {
			params = append(params,
				jen.Id(arg.Name).Id(string(arg.OrigType)),
			)
			callParams = append(callParams, jen.Id(argCallName(arg)))
			funcBody = g.appendArgsConv(funcBody, arg)
		}

		// Syscall
		call := jen.Qual(puregoQual, "SyscallN")
		outValues := [3]string{"_", "_", "_"}
		if len(fn.ReturnArgs) > 0 {
			for i, a := range fn.ReturnArgs {
				outValues[i] = "_r" + string('0'+i)
				returnTypes = append(returnTypes, jen.Id(a.OrigType))
			}
			call = jen.Id(outValues[0]).Op(",").
				Id(outValues[1]).Op(",").
				Id(outValues[2]).Op(":=").Add(
				call,
			)
		}
		call = call.Call(callParams...)
		funcBody = append(funcBody, call)
		// Return
		var returnValues []jen.Code
		for i, arg := range fn.ReturnArgs {
			arg.Name = outValues[i]
			var conv bool
			funcBody, conv = g.appendRetsConv(funcBody, arg)
			if conv {
				returnValues = append(returnValues, jen.Id("_"+arg.Name))
			} else {
				returnValues = append(returnValues, jen.Id(arg.Name))
			}
		}
		if len(returnValues) > 0 {
			funcBody = append(funcBody, jen.Return(returnValues...))
		}
		// Append func
		funcs = append(funcs,
			jen.Id(fn.Name).Op("=").Func().
				Params(params...).
				Params(returnTypes...).
				Block(funcBody...),
		)
	}
	// Errors
	if len(g.errors) > 0 {
		return nil, fmt.Errorf("generate: %d error(s) encountered", len(g.errors))
	}
	f.Func().Id("init").Params().Block(
		//jen.Commentf("// Functions - %s", "").Add( // TODO: add comments
		append(initBody, funcs...)...,
	//),
	)
	// Append functions implementations file
	files = append(files, &File{
		Filename: outputFilename(g.filename, implFileSuffix),
		Content:  f.GoString(),
	})

	// Libraries loaders files
	// TODO: additional files per OS here

	return files, nil
}

func (g *Generator) Errors() []error {
	return g.errors
}
