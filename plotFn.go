package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot"
)

type xy struct{ x, y float64 }

func readData(path string) ([]xy, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var xys []xy

	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f, %f", &x, &y)
		if err != nil {
			log.Printf("Discd &q, %v", s.Text(), err)
		}

		xys = append(xys, xy{x, y})
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

	err = plotData("out.png")
	if err != nil {
		log.Fatalf("Could not plot the graph %v", err)
	}

}

func plotData(path string, xys []xy) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create png")
	}

	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("Cannot create plot funtion")
	}

	plotter
	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		return fmt.Errorf("Could not create a writer %v", err)
	}
	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("Could not write to out.png: %v", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("Could not close another png: %v", err)
	}

}
