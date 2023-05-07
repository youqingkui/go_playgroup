package ch11

import (
	"testing"
)

type Programmer interface {
	WriteHelloWord() string
}

type GoProgeammer struct {
}

func (e *GoProgeammer) WriteHelloWord() string {
	return "fmt.Println(\"hello word\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgeammer)
	t.Log(p.WriteHelloWord())
}
