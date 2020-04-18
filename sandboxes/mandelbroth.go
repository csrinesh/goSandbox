package mandelbroth

import (
	"fmt"
	"math/cmplx"
	"image/color"
	"log"
	"os"
	"time"

	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot"
)

var (
	n  int = 100
	xvalues []float64
	yvalues []float64
	maxiter int = 100
)

func makeRange(min , max int) []int {
	a := make([]int, n+1)

    for i:= range a {
        a[i] = min + i
    }
    return a
}

func makeArange(min float64, max float64, step int) []float64 {
	n := (max-min)/float64(step)
	a := make([]float64, step+1)

    for i:= range a {
        a[i] = min + float64(i)*n
    }
    return a
}

func loopArray(xvalues, yvalues []float64, nrange []int) (plotter.XYs){
	
	var xys plotter.XYs
	for _, x:= range xvalues {
		for _, y:= range yvalues{
			var z complex128 
			var c complex128 = complex(x, y)
			for _= range(nrange) {
				z = z*z + c
				if cmplx.Abs(z) > 2.0 {
					xys = append(xys, struct{ X, Y float64 }{real(c), imag(c)})
					break
				}
			}
		}					
	}
	return xys
}
func mandelbroth() {
	start := time.Now()

	xvalues := makeArange(-2,2, n)
	yvalues := makeArange(-2,2, n)
	nrange := makeRange(0,100)
	
	xys := loopArray(xvalues, yvalues, nrange)

	fmt.Println("The main starts here ....")
	
	err := plotData("mandelbroth.png", xys)
	if err != nil {
		log.Fatalf("Could not plot the graph %v", err)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("The main ends here .... and the total time elapsed is:", elapsed)
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
	s.Color = color.RGBA{B: 255, A: 255}
	s.Radius = 0.5
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
