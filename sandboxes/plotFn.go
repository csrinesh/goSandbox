package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

type xy struct{ x, y float64 }

func readData(path string) (plotter.XYs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var xys plotter.XYs

	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f, %f", &x, &y)
		if err != nil {
			log.Printf("Discd &q, %v", s.Text(), err)
		}

		xys = append(xys, struct{ X, Y float64 }{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("cannot read: %v", err)
	}

	return xys, nil
}

func main() {
	xys, err := readData("data.txt")
	if err != nil {
		log.Fatalf("cannot read the txt: %v", err)
	}

	_ = xys

	err = plotData("out.png", xys)
	if err != nil {
		log.Fatalf("Could not plot the graph %v", err)
	}
}

func plotData(path string, xys plotter.XYs) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create png")
	}

	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("Cannot create plot funtion")
	}

	s, err := plotter.NewScatter(xys)

	if err != nil {
		return fmt.Errorf("Could not create scatter: %s: %v", path, err)
	}
	s.GlyphStyle.Shape = draw.SquareGlyph{}
	s.Color = color.RGBA{B: 255, A: 255}
	s.Radius = 2
	p.Add(s)

	wt, err := p.WriterTo(256, 256, "png")
	if err != nil {
		return fmt.Errorf("Could not create a writer: %s: %v", path, err)
	}
	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("Could not write to %s: %v", path, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("Could not close another png: %v", err)
	}
	return nil
}
