package lexer

import (
	"testing"
  "github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
  l := New("echo 123\n")
  items := l.Tokenize()
  assert.Equal(t, Item{t: Symbol, value: "echo"}, items[0])
  assert.Equal(t, Item{t: Symbol, value: "123"}, items[1])
  assert.Equal(t, Item{t: EOF, value: ""}, items[2])
}

func TestSequentialCommands(t *testing.T) {
  l := New("echo 123; cat patch\n")
  items := l.Tokenize()
  assert.Equal(t, Item{t: Symbol, value: "echo"}, items[0])
  assert.Equal(t, Item{t: Symbol, value: "123"}, items[1])
  assert.Equal(t, Item{t: Semicolon, value: ";"}, items[2])
  assert.Equal(t, Item{t: Symbol, value: "cat"}, items[3])
  assert.Equal(t, Item{t: Symbol, value: "patch"}, items[4])
  assert.Equal(t, Item{t: EOF, value: ""}, items[5])
}

func TestListWithAndOperator(t *testing.T) {
  l := New("echo 123 && cat patch\n")
  items := l.Tokenize()
  assert.Equal(t, Item{t: Symbol, value: "echo"}, items[0])
  assert.Equal(t, Item{t: Symbol, value: "123"}, items[1])
  assert.Equal(t, Item{t: DoubleAmpersand, value: "&&"}, items[2])
  assert.Equal(t, Item{t: Symbol, value: "cat"}, items[3])
  assert.Equal(t, Item{t: Symbol, value: "patch"}, items[4])
  assert.Equal(t, Item{t: EOF, value: ""}, items[5])
}

func TestListWithOrOperator(t *testing.T) {
  l := New("echo 123 || cat patch\n")
  items := l.Tokenize()
  assert.Equal(t, Item{t: Symbol, value: "echo"}, items[0])
  assert.Equal(t, Item{t: Symbol, value: "123"}, items[1])
  assert.Equal(t, Item{t: DoubleVerticalBar, value: "||"}, items[2])
  assert.Equal(t, Item{t: Symbol, value: "cat"}, items[3])
  assert.Equal(t, Item{t: Symbol, value: "patch"}, items[4])
  assert.Equal(t, Item{t: EOF, value: ""}, items[5])
}


func TestPipe(t *testing.T) {
  l := New("echo 123 | cat patch\n")
  items := l.Tokenize()
  assert.Equal(t, Item{t: Symbol, value: "echo"}, items[0])
  assert.Equal(t, Item{t: Symbol, value: "123"}, items[1])
  assert.Equal(t, Item{t: Pipe, value: "|"}, items[2])
  assert.Equal(t, Item{t: Symbol, value: "cat"}, items[3])
  assert.Equal(t, Item{t: Symbol, value: "patch"}, items[4])
  assert.Equal(t, Item{t: EOF, value: ""}, items[5])
}
