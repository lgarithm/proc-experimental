package builtin

import (
	"bytes"
	"testing"

	"github.com/lgarithm/proc-experimental/execution"
	"github.com/lgarithm/proc-experimental/iostream"
)

func Test_echo(t *testing.T) {
	p := Echo(`pong`)
	out := &bytes.Buffer{}
	stdpipe := &iostream.StdWriters{
		Stdout: out,
	}
	execution.Run(p, stdpipe)
	want := "pong\n"
	if got := out.String(); got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
