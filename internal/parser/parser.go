package parser

import (
  "github.com/bogdan-lytvynov/go-shell/internal/lexer"
  "github.com/bogdan-lytvynov/go-shell/internal/common"
)

//var itemsChannelAlreadyClosed = errors.New("Items channel is already closed")

var stopTokens = map[lexer.ItemType]bool{
  lexer.Semicolon:  true,
  lexer.Pipe: true,
  lexer.EOF: true,
}

type Parser struct {
  items []lexer.Item
  start int
  pos int
}

func New(items []lexer.Item) (*Parser) {
  p := &Parser{
    items: items,
  }

  return p
}


func (p *Parser) Parse() common.Cell {
  return p.parseList()
}

func (p *Parser) next() lexer.Item {
  item := p.items[p.pos]
  //first get elemet then increment
  // it is needed to return 0 element on first next() call
  p.pos++
  return item
}

func (p *Parser) peek() lexer.Item {
  return p.items[p.pos]
}

func (p *Parser) acceptUntil(t lexer.ItemType) []lexer.Item  {
  slice := []lexer.Item{} 

  for  p.peek().Type() != lexer.EOF && p.peek().Type() != t {
    slice = append(slice, p.next()) 
  }

  return slice
}

func (p *Parser) ignoreNext() {
  p.pos++
}

func (p *Parser) parseCmd() common.Cmd{
  var args = []string{}

  for !stopTokens[p.peek().Type()] {
    i := p.next()
    args = append(args, i.Value())
  }

  return common.NewCmd(args)
}

func (p *Parser) parsePipeline() common.Cell{
  var cmdList []common.Cmd
  Loop: 
  for {
    cmd := p.parseCmd()
    
    switch p.peek().Type() {
    case lexer.Pipe:
      cmdList = append(cmdList, cmd)
      //ignore pipe token
      p.ignoreNext()
    default:
      if len(cmdList) == 0 {
        return cmd 
      }
      cmdList = append(cmdList, cmd)
      break Loop
    }
  }

  return common.NewPipeline(cmdList)
}

func (p *Parser) parseList() common.Cell {
  var listItems []common.ListItem
  loop: 
  for {
    pipeline := p.parsePipeline()

    switch p.peek().Type() {
    case lexer.Semicolon:
      p.ignoreNext()
      listItems = append(listItems, common.NewListItem(pipeline, common.Next))
    default:
      // it is not a list but maybe pipeline
      if len(listItems) == 0 {
        return pipeline
      }
      // it is a last element of the list so separator is common.None
      listItems = append(listItems, common.NewListItem(pipeline, common.None))
      break loop
    }
  }
  return common.NewList(listItems)
}
