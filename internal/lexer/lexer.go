package lexer

import (
  //"fmt"
  "strings"
  "unicode/utf8"
)
// ItemType enum
type ItemType = int
const (
  // EOF Symbol to signal about end of the string
  Symbol ItemType = iota 
  DoubleAmpersand
  DoubleVerticalBar
  Semicolon
  Pipe
  EOF
)

const (
  eof = -1
  //doubleQuote = '"'
  whitespace = ' '
  semicolon = ';'
  ampersand = '&'
  verticalBar = '|'
)

// Item struct
type Item struct { 
  t ItemType
  value string
}

func NewItem(t ItemType, value string) Item { 
  return Item{
    t: t,
    value: value,
  }
}

func (i Item) Type() ItemType {
  return i.t
}

func (i Item) Value() string {
  return i.value
}

type stateFn func(*Lexer) stateFn

type Lexer struct {
  input string
  start int
  pos int
  width int
  items []Item
  
}

// New Lexer for passed string
func New(input string) (*Lexer) { 
  lexer := &Lexer{
    input: input,
    pos:0,
    start:0,
    width:0,
    items: []Item{},
  }

  return lexer
}

func (l *Lexer) emit(t ItemType) {
  if l.start == l.pos {
    //nothing to emit
    return
  }
  l.items = append(l.items, Item{t, l.input[l.start:l.pos]})
  l.start = l.pos
}

func (l *Lexer) emitEOF() {
  l.items = append(l.items, Item{t: EOF, value: ""})
}

func (l *Lexer) Tokenize() []Item {
  for state := symbolState; state != nil; {
    state = state(l)
  }

  l.emitEOF()

  return l.items
}

func (l *Lexer) next() rune {
  if l.pos >= len(l.input) {
    return eof
  }
  var r rune
  r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
  l.pos += l.width
  return r
}

func (l *Lexer) ignore() {
  l.start = l.pos;
}

func (l *Lexer) backup() {
  l.pos -= l.width;
}

func (l *Lexer) peek() rune {
  r := l.next()
  l.backup()
  return r
}

func (l *Lexer) accept(valid string) bool {
  if strings.IndexRune(valid, l.next()) >= 0 {
    return true
  }
  l.backup()
  return false
}

func (l *Lexer) acceptRun(valid string) {
  for strings.IndexRune(valid, l.next()) >= 0 {
  }
  l.backup()
}

func (l *Lexer) runUntil(stop string) {
  for n := l.next(); n!=eof && strings.IndexRune(stop, n) == -1; n=l.next() {
  }
  l.backup()
}

func symbolState(l *Lexer) stateFn {
  l.runUntil("|&; ")
  l.emit(Symbol)

  switch l.peek() {
  case whitespace:
    return skipWhitespaceState
  case semicolon:
    return semicolonState
  case ampersand:
    return amersandState
  case verticalBar:
    return verticalBarState
  case eof:
    return nil
  default:
    //imposible state
    return nil
  }
}

func skipWhitespaceState(l *Lexer) stateFn {
  l.acceptRun(" ")
  l.ignore()
  return symbolState
}

func semicolonState(l *Lexer) stateFn {
  l.accept(";")
  l.emit(Semicolon)
  switch l.peek() {
    case eof:
      return nil
    case whitespace:
      return skipWhitespaceState
    default:
      return symbolState
  }
}

func amersandState(l *Lexer) stateFn {
  l.accept("&")
  switch l.peek() {
    // it is background job case but we don't handle it yet
  case eof:
    return nil
    // it is backgrund job case too but we don't handle it yet too
  case whitespace:
    return skipWhitespaceState
  // double ampersand case
  case ampersand:
    return doubleAmersand
  default:
    return nil
  }
}

func doubleAmersand(l *Lexer) stateFn {
  l.accept("&")
  l.emit(DoubleAmpersand)
  switch l.peek() {
  case eof:
    return nil
  case whitespace:
    return skipWhitespaceState
  default:
    return nil
  }
}

func verticalBarState(l *Lexer) stateFn {
  l.accept("|")

  switch l.peek() {
  case eof:
    return nil
  case verticalBar:
    return doubleVerticalBarState
  default:
    l.emit(Pipe)
    return symbolState
  }
}

func doubleVerticalBarState(l *Lexer) stateFn {
  l.accept("|")
  l.emit(DoubleVerticalBar)

  switch l.peek() {
  case eof:
    return nil
  case whitespace:
    return skipWhitespaceState
  default:
    return nil
  }
}
