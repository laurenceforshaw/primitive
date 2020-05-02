package primitive

import (
	"math/rand"
	"strings"
	"testing"
)

func FlexPolyErrTest(t *testing.T,in string, TargErr string){
	_, err := ParseFlexPoly(strings.Split(in,","))
	if(err == nil){
		t.Errorf("No error for bad input.")
	} else if (err.Error() != TargErr){
		t.Errorf("Error text was \"%s\", expected \"%s\"",err.Error(),TargErr)
	}
}

func TestFlexPolyEmpty(t *testing.T){
	FlexPolyErrTest(t,"FlexPoly","FlexPoly requires 2 arguments")
	FlexPolyErrTest(t,"FlexPoly,2,3","FlexPoly requires 2 arguments")
}

func TestFlexPolyBadOrder(t *testing.T){
	badOrder := "bad order"
	FlexPolyErrTest(t,"FlexPoly," + badOrder,"strconv.Atoi: parsing \"" + badOrder + "\": invalid syntax")
}

func TestFlexPolyLowOrder(t *testing.T){
	FlexPolyErrTest(t,"FlexPoly,2" ,"A polygon must have at least 3 sides.")
}

func TestGoodFlexPoly(t *testing.T){
	res,err := ParseFlexPoly(strings.Split("FlexPoly,3",","))
	if(err != nil){
		t.Errorf(err.Error())
	}else if(res.(*FlexPolyFactory).n != 3){
		t.Errorf("Wrong number of sides.")
	}
}

func MakeDummyWorker(rnd *rand.Rand) *Worker{
	ret := Worker{100, 100, nil,nil,nil,nil,nil,nil,rnd,0,0,nil,nil,nil}
	return &ret
}

func TestMakeFlexPoly(t *testing.T){
	fact := FlexPolyFactory{4}
	if (fact.NewShape(MakeDummyWorker(rand.New(rand.NewSource(1)))).(*Polygon).Order != 4){
		t.Errorf("Bad order number")
	}
}
