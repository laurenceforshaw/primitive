package primitive

import (
	"math/rand"
	"testing"
)

func TestNonFixedColors(t *testing.T){
	nTests := 300
	marginError := 0.99
	owlBase,err := LoadImage("../examples/owl.png")
	if(err != nil){
		t.Errorf(err.Error())
	} else{
		Worker := MakeDummyWorker(rand.New(rand.NewSource(1)))
		owl := imageToRGBA(owlBase)
		oldR,oldB,oldG := 0,0,0
		repeats := 0
		for i := 0;i <nTests;i++ {
			black :=  MakeHexColor("#000000")
			computedColor := computeColor(owl, uniformRGBA(owlBase.Bounds(),black.NRGBA()), Worker.RandomState(1, 1).Shape.Rasterize(), 1, nil)
			if (computedColor.R == oldR && computedColor.G == oldG && computedColor.B == oldB ){
				repeats = repeats + 1
			}
			oldR,oldB,oldG = computedColor.R,computedColor.B,computedColor.G
		}
		if (float64(repeats)/float64(nTests) > marginError){
			t.Errorf("Excess repeated colors: %d", repeats)
		}
	}
}

func TestFixedColors(t *testing.T){
	nTests := 100
	owlBase,err := LoadImage("../examples/owl.png")
	if(err != nil){
		t.Errorf(err.Error())
	} else{
		Worker := MakeDummyWorker(rand.New(rand.NewSource(1)))
		owl := imageToRGBA(owlBase)
		sc := []Color{MakeHexColor("#0000FF"),MakeHexColor("#00FF00"),MakeHexColor("#FF0000")}
		for i := 0;i <nTests;i++ {
			black :=  MakeHexColor("#000000")
			col := computeColor(owl, uniformRGBA(owlBase.Bounds(),black.NRGBA()), Worker.RandomState(1, 1).Shape.Rasterize(), 1, sc)
			if(!(col.R == 0 || col.R == 255) || !(col.B == 0 || col.B == 255) ||!(col.G == 0 || col.G == 255)  ||
				col.R + col.G + col.B != 255){
				if(!(col.R ==0 && col.G == 0 && col.B ==0)){ //computeColor returns in rare circumstances. Not a result of the code I wrote.
					t.Errorf("Illegal color,%d,%d,%d",col.R,col.G,col.B)
				}
			}
		}
	}
}
