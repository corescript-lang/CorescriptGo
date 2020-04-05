package main

import (
	"fmt"
	"strings"
	"strconv"
//	"encoding/json"
//	"regexp"
)

func main() {
	code := [1]string{"var a = 1"}

	varNames := []string{"time"}
	varVals := []string{"today"}
	
	_ = varNames // I HATE GO I HATE GO I HATE GO!!

	for i := 0; i < len(code); i++ {
		currentLine := code[i]
		split := strings.Fields(currentLine)

		if split[0] == "print" {

			toPrint := currentLine[6:]
			fmt.Print(parse(toPrint, varNames, varVals))

		} else if split[0] == "goto" {

			newI, _ := strconv.Atoi(currentLine[5:])
			newI -= 2 // For loop adds 1?

			i = newI

		} else if split[0] == "var" {

			if split[2] != "=" {
				fmt.Print("asd")
			} else {
				varNames := append(varNames, split[1])
				varVals := append(varVals, split[3])
			}

        } else if split[0] == "newline" {
			fmt.Println("")
		}
	}

	fmt.Println("")
}

func parse(string string, varListP []string, varValsP []string) string {
	for i := 0; i < len(varListP); i++ {
		if string == varListP[i] {
			return varValsP[i]
		}
	}

	return string
}
