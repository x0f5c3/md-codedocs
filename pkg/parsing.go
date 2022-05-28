package pkg

import (
	"errors"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
)

var CodeNodes []*blackfriday.Node
var CodeBlockNodes []*blackfriday.Node

func ParseFile(path string) (*blackfriday.Node, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	parser := blackfriday.New()
	res := parser.Parse(b)
	if res == nil {
		return nil, errors.New("failed to parse")
	}
	return res, nil
}

func TestWalker(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if node.Type == blackfriday.CodeBlock || node.Type == blackfriday.Code {
		//pterm.Info.Printfln("Node type: %s\nEntering: %t\n", node.Type, entering)
		if node.Type == blackfriday.Code {
			//pterm.Info.Println("Found a code node, appending")
			CodeNodes = append(CodeNodes, node)
		}
		if node.Type == blackfriday.CodeBlock && entering {
			//pterm.Info.Println("Found a code block node, appending")
			CodeBlockNodes = append(CodeBlockNodes, node)
		}
	}
	return blackfriday.GoToNext
}
