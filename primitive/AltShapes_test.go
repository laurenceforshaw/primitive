package primitive

import (
	"math"
	"math/rand"
	"strings"
	"testing"
)

func TestOldMode0 (t *testing.T){
	nTests := 3000
	MarginError := 0.1
	Buckets := make([]int,7)
	Worker := MakeDummyWorker(rand.New(rand.NewSource(1)))
	for i :=0;i <nTests;i++{
		switch Worker.RandomState(0,1).Shape.(type){
		case *Triangle:
			Buckets[0] = Buckets[0] + 1
		case *Rectangle:
			Buckets[1] = Buckets[1] + 1
		case *Ellipse:
			Buckets[2] = Buckets[2] + 1
		case *RotatedRectangle:
			Buckets[3] = Buckets[3] + 1
		case *Quadratic:
			Buckets[4] = Buckets[5] + 1
		case *RotatedEllipse:
			Buckets[5] = Buckets[5] + 1
		case *Polygon:
			Buckets[6] = Buckets[6] + 1
		default:
			t.Errorf("Unknown shape type")
		}
	}
	for i := 0;i < 7;i++{
		var f int
		if (i == 2){
			f = 4
		}else {
			f = 8
		}
		if(math.Abs(float64(nTests)/float64(Buckets[i]*f) - 1) > MarginError){
			t.Errorf("bad shape ratio")
		}
	}
}

func TestMultiMode (t *testing.T){
	nTests := 3000
	MarginError := 0.1
	Buckets := make([]int,2)
	Worker := MakeDummyWorker(rand.New(rand.NewSource(1)))
	Worker.ModeArr =[]int {1,2}
	for i :=0;i <nTests;i++{
		switch Worker.RandomState(0,1).Shape.(type){
		case *Triangle:
			Buckets[0] = Buckets[0] + 1
		case *Rectangle:
			Buckets[1] = Buckets[1] + 1
		default:
			t.Errorf("Unknown shape type")
		}
	}
	for i := 0;i < 2;i++{
		if(math.Abs(float64(nTests)/float64(Buckets[i]*2) - 1) > MarginError){
			t.Errorf("bad shape ratio")
		}
	}
}

func TestUserShapes (t *testing.T){
	nTests := 3000
	MarginError := 0.1
	Buckets := make([]int,2)
	Worker := MakeDummyWorker(rand.New(rand.NewSource(1)))
	Shapes,err := ParseShapesFile(strings.NewReader("FlexPoly,4\nRigidPoly,3,0,0,0,1,1,0"))
	if(err != nil){
		t.Errorf(err.Error())
	}
	Worker.usSH = Shapes
	for i :=0;i <nTests;i++{
		switch Worker.RandomState(9,1).Shape.(type){
		case *Polygon:
			Buckets[0] = Buckets[0] + 1
		case *RigidPolygon:
			Buckets[1] = Buckets[1] + 1
		default:
			t.Errorf("Unknown shape type" )
		}
	}
	for i := 0;i < 2;i++{
		if(math.Abs(float64(nTests)/float64(Buckets[i]*2) - 1) > MarginError){
			t.Errorf("bad shape ratio")
		}
	}
}


