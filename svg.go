package main

import "github.com/fogleman/gg"

type Drawer interface {
	Draw(ctx *gg.Context)
}

type SVG struct {
	Width, Height int
	Shapes        []Drawer
}

type Shape struct {
	X, Y         float64
	Stroke, Fill string
	StrokeWidth  float64
}

type Circle struct {
	Shape
	Radius float64
}

func (c *Circle) Draw(ctx *gg.Context) {
	ctx.DrawCircle(c.X, c.Y, c.Radius)
	ctx.SetHexColor(c.Fill)
	ctx.Fill()

	ctx.DrawCircle(c.X, c.Y, c.Radius)
	ctx.SetHexColor(c.Stroke)
	ctx.SetLineWidth(c.StrokeWidth)
	ctx.Stroke()
}
