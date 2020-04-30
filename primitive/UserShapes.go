package primitive

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

//This pares a user shapes file
func ParseShapesFile(reader io.Reader) []ShapeFactory{
	scanner := bufio.NewScanner(reader)
	var res []ShapeFactory
	cont := true
	for cont{
		success := scanner.Scan()
		if success{
			sp := strings.Split(scanner.Text(),",")
			switch sp[0] {
			default:
				fmt.Printf("Error reading user shape file: Illegal shape name :%v", sp[0])
				os.Exit(1)
			case "FlexPoly":
				res = append(res, ParseFlexPoly(sp))
			case "RigidPoly":
				res = append(res, ParseRigidPoly(sp))
			}
		} else {
			err := scanner.Err()
			if(err == nil){
				cont = false
			}else{
				fmt.Printf("Error reading user shape file: %s",err.Error())
				os.Exit(1)
			}
		}
	}
	if(len(res) == 0){
		fmt.Printf("Error reading user shape file: at least one shape must be specified")
		os.Exit(1)
	}
	return res
}

//parse a flexible polygon creator from a seperated line
func ParseFlexPoly(sp []string) ShapeFactory{

	if(len(sp) != 2){
		fmt.Printf("Error reading user shape file: FlexPoly requires 2 arguments")
		os.Exit(1)
	}
	s,err := strconv.Atoi(sp[1])
	if(err != nil){
		fmt.Printf("Error reading user shape file: %s",err.Error())
		os.Exit(1)
	}
	if(s < 3){
		fmt.Printf("Error reading user shape file: A polygon must have at least 3 sides.")
		os.Exit(1)
	}
	ret := FlexPolyFactory{s}
	return &ret

}
