package ch14

import (
	"errors"
	"testing"
)

func TestPanicVsExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("recovered from ", err)
		}
	}()
	t.Log("Start")
	panic(errors.New("get some wrong"))
	// os.Exit(-1)

}

func TestRevise(t *testing.T) {

}
