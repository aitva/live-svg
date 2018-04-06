package main

import (
	"fmt"
	"io/ioutil"

	"github.com/fogleman/gg"
)

func main() {
	data, _ := ioutil.ReadFile("circle.svg")
	fmt.Println(string(data))

	lexer := NewLexer(string(data))
	//for tok := lexer.NextToken(); tok.Type != TokenEOF; tok = lexer.NextToken() {
	//	fmt.Println(tok)
	//}
	parser := NewParser(lexer)
	svg := parser.ParseSVG()
	//fmt.Printf("%#v\n", svg)
	ctx := gg.NewContext(svg.Width, svg.Height)
	for _, s := range svg.Shapes {
		s.Draw(ctx)
	}
	ctx.SavePNG("out.png")
}
