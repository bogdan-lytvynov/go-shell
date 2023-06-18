package common

/*
 separators = (; || &&)
 pipe = (| |&)
 maybeList = maybePipe <separator> maybePipe
 maybePipe = maybeCmd <pipe> maybeCmd 
 maybeCmd = parse until any stop token ( <separator> and <pipe> )
 */

type CellType = int
type ListSeparatorType = int


const (
  CmdCell CellType = iota
  PipelineCell
  ListCell

  Next ListSeparatorType = iota
  None

//  Output PipeType = iota // |
//  OutputError            // |&
//
//  And OperatorType = iota // &&
//  Or // ||
//  Next // ;
//  Background // &
//  None
//
//  SameShell GroupType = iota // ()
//  SubShell // {}
  //Separators = []lexer.ItemTypes{
  //  lexer.DoubleAmpersand,
  //  lexer.DoubleVerticalBar,
  //  Semicolon,
  //}
)


type Cell interface {
  Type() CellType 
}

type Cmd struct {
  args []string
}

func NewCmd(args []string) Cmd {
  return Cmd{
    args,
  }
}

func (c Cmd) Type() CellType {
  return CmdCell
}

func (c Cmd) Cmd() string {
  return c.args[0]
}

func (c Cmd) Args() []string {
  return c.args[1:len(c.args)]
}

type Pipeline []Cmd

func NewPipeline(cmdList []Cmd) Pipeline {
  return Pipeline(cmdList)
}

func (p Pipeline) Type() CellType {
  return PipelineCell
}

type ListItem struct {
  cell Cell
  separator ListSeparatorType
}

func NewListItem(cell Cell, separator ListSeparatorType) ListItem {
  return ListItem{
    cell,
    separator,
  }
}

type List []ListItem

func NewList(list []ListItem) List {
  return List(list)
}

func (l List) Type() CellType {
  return ListCell
}


