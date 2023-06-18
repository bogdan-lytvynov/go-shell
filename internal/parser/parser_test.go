package parser

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/bogdan-lytvynov/go-shell/internal/lexer"
  "github.com/bogdan-lytvynov/go-shell/internal/common"
)

func TestSimpleCommand(t *testing.T) {
  t.Skip()
  items := []lexer.Item{
    lexer.NewItem(lexer.Symbol, "echo"),
    lexer.NewItem(lexer.Symbol, "hello"),
    lexer.NewItem(lexer.Symbol, "world"),
    lexer.NewItem(lexer.EOF, ""),
  }
  p := New(items)
  
  c := p.Parse()
  assert.Equal(t, c.Type(), common.CmdCell) 
  assert.Equal(t, common.NewCmd([]string{"echo", "hello", "world"}), c)
}


func TestSimplePipe(t *testing.T) {
  items := []lexer.Item{
    lexer.NewItem(lexer.Symbol, "cat"),
    lexer.NewItem(lexer.Symbol, "./f.txt"),

    lexer.NewItem(lexer.Pipe, "|"),

    lexer.NewItem(lexer.Symbol, "wc"),

    lexer.NewItem(lexer.EOF, ""),
  }

  p := New(items)
  
  c := p.Parse()
  assert.Equal(t, c.Type(), common.PipelineCell) 
  assert.Equal(t, common.NewPipeline([]common.Cmd{
    common.NewCmd([]string{"cat", "./f.txt"}),
    common.NewCmd([]string{"wc"}),
  }), c)
}


func TestList(t *testing.T) {
  items := []lexer.Item{
    lexer.NewItem(lexer.Symbol, "cat"),
    lexer.NewItem(lexer.Symbol, "./f.txt"),

    lexer.NewItem(lexer.Semicolon, ";"),

    lexer.NewItem(lexer.Symbol, "echo"),
    lexer.NewItem(lexer.Symbol, "foo"),

    lexer.NewItem(lexer.EOF, ""),
    
  }

  p := New(items)
  
  c := p.Parse()
  assert.Equal(t, c.Type(), common.ListCell) 
  assert.Equal(t, common.NewList([]common.ListItem{
    common.NewListItem(
      common.NewCmd([]string{"cat", "./f.txt"}),
      common.Next,
    ),
    common.NewListItem(
      common.NewCmd([]string{"echo", "foo"}),
      common.None,
    ),
  }), c)
}
