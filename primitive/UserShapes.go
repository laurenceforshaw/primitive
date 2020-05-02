package primitive

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type ShapeFactory interface {
	NewShape(worker *Worker)  Shape
}

type FlexPolyFactory struct{
	n int
}


func(fact *FlexPolyFactory) NewShape(worker *Worker) Shape{
	return NewRandomPolygon(worker,fact.n,false)
}

//This parses a user shapes file
func ParseShapesFile(reader io.Reader) ([]ShapeFactory, error){
	scanner := bufio.NewScanner(reader)
	var res []ShapeFactory
	cont := true
	for cont{
		success := scanner.Scan()
		if success{
			sp := strings.Split(scanner.Text(),",")
			switch sp[0] {
			default:
				return nil,fmt.Errorf("Illegal shape name: %v", sp[0])
			case "FlexPoly":
				next,err := ParseFlexPoly(sp)
				if (err != nil){
					return nil,err
				}
				res = append(res, next)
			case "RigidPoly":
				next,err := ParseRigidPoly(sp)
				if (err != nil){
					return nil,err
				}
				res = append(res, next)
			}
		} else {
			err := scanner.Err()
			if(err == nil){
				cont = false
			}else{
				return nil,err
			}
		}
	}
	if(len(res) == 0){
		return nil, fmt.Errorf("at least one shape must be specified")
	}
	return res,nil
}

//parse a flexible polygon creator from a seperated line
func ParseFlexPoly(sp []string) (ShapeFactory,error){

	if(len(sp) != 2){
		return nil, fmt.Errorf("FlexPoly requires 2 arguments")
	}
	s,err := strconv.Atoi(sp[1])
	if(err != nil){
		return nil,err
	}
	if(s < 3){
		return nil, fmt.Errorf("A polygon must have at least 3 sides.")
	}
	ret := FlexPolyFactory{s}
	return &ret,nil

}
