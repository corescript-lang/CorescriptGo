package main

import (
	"fmt"
	"strings"
	"strconv"
//	"encoding/json"
	"regexp"
)

func main() {
	code := [3]string{"print ¡Hola de Go!", "newline", "print Hello from Go! (test)"}
	
	varVals := map[string]string{}
	varList := map[string]string{}
	err := ""
	
	for i := 0; i < len(code); i++ {
		currentLine := code[i]
		split := strings.Fields(currentLine)
		
		if split[0] == "print" {
			toPrint := currentLine[6:]
			regex := regexp.MustCompile(`\(([a-zA-Z-_]+)\)`)
			match := regex.ReplaceAllString(toPrint, parse("$1", variables))
			fmt.Print(match)
		} else if split[0] == "goto" {
			newI, _ := strconv.Atoi(currentLine[5:])
			newI -= 2 // For loop adds 1?
			
			i = newI
		} else if split[0] == "var" {
			if split[2] != "=" {
				err = "Variable error"
			} else {
				varVals[split[1]] = split[3]
				append(varList, split[3])
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

func parse(string) {
	for i := 0; i < len(varList); i++ {
		if string == varList[i] {
			return varVals[string]
		}
	}
}