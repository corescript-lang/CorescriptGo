package main

import (
	"fmt"
	"strings"
	"strconv"
//	"encoding/json"
//	"regexp"
)

func main() {
	code := [3]string{"print ¡Hola de Go!", "newline", "print Hello from Go!"}
	
	variables := map[string]string{}
	err := ""

	for i := 0; i < len(code); i++ {
		currentLine := code[i]
		split := strings.Fields(currentLine)
		
		if split[0] == "print" {
			toPrint := currentLine[6:]
			
			//match, _ := regexp.ReplaceAllLiteral(toPrint, "\(([a-zA-Z-_]+)\)", "$get1")
			fmt.Print(toPrint)
		} else if split[0] == "goto" {
			newI, _ := strconv.Atoi(currentLine[5:])
			newI -= 2 // For loop adds 1?
			
			i = newI
		} else if split[0] == "var" {
			if split[2] != "=" {
				err = "Variable error"
			} else {
				variables[split[1]] = split[3]
			}
                } else if split[0] == "newline" {
			fmt.Println("")
		}
	}

	if err != "" {
		fmt.Println(err)
	}

	fmt.Println("")
}

