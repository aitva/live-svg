package main

import "strconv"

type Parser struct {
	lexer               *Lexer
	currentTok, nextTok Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{lexer: l}
	p.readToken()
	p.readToken()
	return p
}

func (p *Parser) readToken() {
	p.currentTok = p.nextTok
	p.nextTok = p.lexer.NextToken()
}

func (p *Parser) ParseSVG() *SVG {
	svg := &SVG{}

	p.expect(TokenOpen)
	tag := p.parseIdent()
	if tag != "svg" {
		panic("expected svg")
	}
	attributes := p.parseAttributes()
	svg.Width, _ = strconv.Atoi(attributes["width"])
	svg.Height, _ = strconv.Atoi(attributes["height"])
	p.expect(TokenClose)

	for p.currentTok.Type != TokenOpen || p.nextTok.Type != TokenSlash {
		s := p.parseShape()
		svg.Shapes = append(svg.Shapes, s)
	}

	return svg
}

func (p *Parser) expect(t TokenType) {
	if p.currentTok.Type != t {
		panic("expected " + t + " got " + p.currentTok.Type)
	}
	p.readToken()
}

func (p *Parser) parseIdent() string {
	literal := p.currentTok.Litteral
	p.expect(TokenIdentifier)
	return literal
}

func (p *Parser) parseAttributes() map[string]string {
	attributes := map[string]string{}

	for p.currentTok.Type != TokenClose && p.currentTok.Type != TokenSlash {
		k, v := p.parseAttribute()
		attributes[k] = v
	}

	return attributes
}

func (p *Parser) parseAttribute() (k, v string) {
	k = p.parseIdent()
	p.expect(TokenEqual)
	p.expect(TokenQuote)
	v = p.parseIdent()
	p.expect(TokenQuote)
	return
}

func (p *Parser) parseShape() Drawer {
	p.expect(TokenOpen)
	tag := p.parseIdent()
	attributes := p.parseAttributes()
	p.expect(TokenSlash)
	p.expect(TokenClose)

	switch tag {
	case "circle":
		x, _ := strconv.ParseFloat(attributes["cx"], 64)
		y, _ := strconv.ParseFloat(attributes["cy"], 64)
		sw, _ := strconv.ParseFloat(attributes["stroke-width"], 64)
		r, _ := strconv.ParseFloat(attributes["r"], 64)
		return &Circle{
			Shape: Shape{
				X:           x,
				Y:           y,
				StrokeWidth: sw,
				Fill:        attributes["fill"],
				Stroke:      attributes["stroke"],
			},
			Radius: r,
		}
	default:
		panic("unknown shape")
	}
}
