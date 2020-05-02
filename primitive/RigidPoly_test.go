package primitive

import (
	"math"
	"math/rand"
	"strings"
	"testing"
)

func RigidPolyErrTest(t *testing.T,in string, TargErr string){
	_, err := ParseRigidPoly(strings.Split(in,","))
	if(err == nil){
		t.Errorf("No error for bad input.")
	} else if (err.Error() != TargErr){
		t.Errorf("Error text was \"%s\", expected \"%s\"",err.Error(),TargErr)
	}
}

func TestRigidPolyEmpty(t *testing.T){
	RigidPolyErrTest(t,"RigidPoly","RigidPoly requires at least 2 arguments")
}

func TestRigidPolyBadOrder(t *testing.T){
	badOrder := "bad order"
	RigidPolyErrTest(t,"RigidPoly," + badOrder,"strconv.Atoi: parsing \"" + badOrder + "\": invalid syntax")
}

func TestRigidPolyLowOrder(t *testing.T){
	RigidPolyErrTest(t,"RigidPoly,2,0,0,1,1" ,"A polygon must have at least 3 sides.")
}

func TestRigidPolyBadCoord(t *testing.T){
	badCoord := "bad coord"
	RigidPolyErrTest(t,"RigidPoly,3,0,0,1,1,1," + badCoord,"strconv.ParseFloat: parsing \"" + badCoord + "\": invalid syntax")
}

func TestRigidPolyWrongCoord(t *testing.T){
	RigidPolyErrTest(t,"RigidPoly,3,0,0,1,1,1", "RigidPoly requires 2 arguments after the first two per vertex")
}

func TestGoodRigidPoly(t *testing.T){
	res,err := ParseRigidPoly(strings.Split("RigidPoly,3,0,0,0,1,1,0",","))
	if(err != nil){
		t.Errorf(err.Error())
	}else if(res.(*RigidPolygonFactory).Order != 3){
		t.Errorf("Wrong number of sides.")
	}
	if (!FloatEqual(res.(*RigidPolygonFactory).X,[]float64{0.0,0.0,1.0}) || !FloatEqual(res.(*RigidPolygonFactory).Y,[]float64{0.0,1.0,0.0})){
		t.Errorf("Bad rigid poly point locations")
	}
}

func Dist(X1,Y1,X2,Y2  float64) float64{
	Xdist := X2 - X1
	Ydist := Y2 - Y1
	return math.Sqrt(Xdist*Xdist + Ydist*Ydist)
}

func AngleDot(X1,Y1,X2,Y2,X3,Y3 float64) float64{
	return (X2- X1)*(X3 - X2) + (Y2- Y1)*(Y3 - Y2)
}

//this runs a number of random trial to make sure polygons do, in fact stay rigid.
func TestMakeRigidPoly(t *testing.T){
	nTests := 20
	minOrder,maxOrder := 3,7
	minMutate,maxMutate := 0,1
	marginRoundingError := 0.01
	sRand := int64(1)
	rnd := rand.New(rand.NewSource(sRand))
	Worker := MakeDummyWorker(rnd)
	for i := 0;i < nTests; i++{
		Order := rnd.Intn(maxOrder - minOrder) + minOrder
		X := make([]float64,Order)
		for j := 0;j< Order;j++{
			X[j] = rnd.Float64()
		}
		Y := make([]float64,Order)
		for j := 0;j< Order;j++{
			Y[j] = rnd.Float64()
		}
		fact := RigidPolygonFactory{Order,X,Y,math.Sqrt(2)}
		res := fact.NewShape(Worker).(*RigidPolygon)
		mutate := rnd.Intn(maxMutate - minMutate) + minMutate
		for j := 0; j < mutate;j++{
			res.Mutate()
		}
		if(res.Order != Order){
			t.Errorf("Wrong Polygon Order")
		}
		for j := 0; j < Order; j++{
			j1 := (j +1)% Order
			j2 := (j + 2)% Order
			if (math.Abs(Dist(res.X[j],res.Y[j],res.X[j1],res.Y[j1])/Dist(X[j],Y[j],X[j1],Y[j1])/math.Abs(res.Scale) - 1) > marginRoundingError){
				t.Errorf("shape deformed ")
			}
			if (math.Abs(AngleDot(res.X[j],res.Y[j],res.X[j1],res.Y[j1],res.X[j2],res.Y[j2])/AngleDot(X[j],Y[j],X[j1],Y[j1],X[j2],Y[j2])/res.Scale/res.Scale - 1) >marginRoundingError){
				t.Errorf("shape angle deformed")
			}
		}
	}
	//fact := nil RigidPolygonFactory{3}
	//if (fact.NewShape(MakeDummyWorker(rand.New(rand.NewSource(1)))).(*Polygon).Order != 4){
	//	t.Errorf("Bad order number")
	//}
}
