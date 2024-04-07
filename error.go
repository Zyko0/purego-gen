package puregogen

import (
	"fmt"
	"go/token"
)

type genError struct {
	level string
	pos   token.Position
	msg   string
}

func (e *genError) Error() string {
	return fmt.Sprintf(
		"%s: %s:(l.%d:%d): %s",
		e.level, e.pos.Filename, e.pos.Line, e.pos.Column, e.msg,
	)
}
