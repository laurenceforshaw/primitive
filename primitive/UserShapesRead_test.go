package primitive

import (
	"strings"
	"testing"
)

func PSFErrTest(t *testing.T,in string, targErr string){
	_,err := ParseShapesFile(strings.NewReader(in))
	if(err == nil){
		t.Errorf("No error for bad input.")
	} else if (err.Error() != targErr){
		t.Errorf("Error text was \"%s\", expected \"%s\"",err.Error(),targErr)
	}
}

func TestBadName(t *testing.T){
	badName := "Bad name"
	PSFErrTest(t,badName + ",7,3","Illegal shape name: " + badName)
}

func TestShapeFileLineError(t *testing.T){
	PSFErrTest(t,"FlexPoly,2","A polygon must have at least 3 sides.")
}

func TestEmptyShapeFile(t *testing.T){
	PSFErrTest(t,"","at least one shape must be specified")
}

func FloatEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestShapeFileGood(t *testing.T){
	res,err := ParseShapesFile(strings.NewReader("FlexPoly,4\nRigidPoly,3,0,0,0,1,1,0"))
	if (err != nil){
		t.Errorf(err.Error())
	}else{
		flex := res[0].(*FlexPolyFactory)
		if (flex.n != 4){
			t.Errorf("FlexPolygon side number should be 4")
		}
		rigid := res[1].(*RigidPolygonFactory)
		if (rigid.Order != 3){
			t.Errorf("RigidPolygon side number should be 3")
		}
		if (!FloatEqual(rigid.X,[]float64{0.0,0.0,1.0}) || !FloatEqual(rigid.X,[]float64{0.0,0.0,1.0})){
			t.Errorf("Bad rigid poly point locations")
		}
	}
}