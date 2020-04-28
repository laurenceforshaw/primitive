package primitive

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)


//This pares a user shapes file
//TODO: return the shapes as a set of factories
func ParseShapesFile(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	cont := true
	for cont{
		success := scanner.Scan()
		if success{
			sp := strings.Split(scanner.Text(),",")
			switch sp[0]{
			default:
				fmt.Printf("Error reading user shape file: Illegal shape name :%v",sp[0])
				os.Exit(1)
			case "FlexPoly":
				ParseFlexPoly(sp)
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

}


func ParseFlexPoly(sp []string) {

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

}
