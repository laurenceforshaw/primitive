package primitive

import (
	"fmt"
	"io"
	"os"
)

//This pares a user shapes file
//TODO: return the shapes as a set of factories
func ParseShapesFile(reader io.Reader){
	n := 1
	for n == 1 {
		var t int
		n,err := fmt.Fscanf(reader,"%d:", &t)
		if(err != nil){
			fmt.Printf("Error reading user shape file: %s",err.Error())
			os.Exit(1)
		}
		if n == 1 {
			switch t {
			default:
				fmt.Printf("Error reading user shape file: Illegal shape number :%d",t)
				os.Exit(1)
			}
		}
	}

}
