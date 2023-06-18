package engine

import (
  "fmt"
  "os/exec"

  "github.com/bogdan-lytvynov/go-shell/internal/common"
  "github.com/bogdan-lytvynov/go-shell/internal/lexer"
  "github.com/bogdan-lytvynov/go-shell/internal/parser"
)

type Engine struct {
}

func New() Engine {
  return Engine{}
}

func (e Engine) Exec(s string) {
  l := lexer.New(s)
  items := l.Tokenize()
  fmt.Println(items)
  p := parser.New(items)
  c := p.Parse()

  switch c.Type() {
  case common.CmdCell:
    e.execCmd(c.(common.Cmd))
  }
}

func (e Engine) execCmd(c common.Cmd) {
  cmd := exec.Command(c.Cmd(), c.Args()...)
  stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
    fmt.Println(err)
	}
  fmt.Printf("%s\n", stdoutStderr)
}
