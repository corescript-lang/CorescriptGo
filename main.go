package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"io/ioutil"
)

func parse(string string, varListP []string, varValsP []string) string {
	for i := 0; i < len(varListP); i++ {
		if string == varListP[i] {
			return varValsP[i]
		}
	}

	return string
}

func run(code []string) {
	varNames := []string{}
	varVals := []string{}
	
	labelNames := []string{}
	labelLines := []int{}

	for l := 0; l < len(code); l++ {
		firstChar := code[l][0:1]
		
		if firstChar == ":" {
			label := code[l][1:]

			append(labelNames, label)
			append(labelLines, l)
		}
	}
	

	for i := 0; i < len(code); i++ { // -1 is because of naughty files!!
		currentLine := code[i]
		split := strings.Fields(currentLine)

		if split[0] == "print" {
			toPrint := currentLine[6:]
			fmt.Print(parse(toPrint, varNames, varVals))
			
			fmt.Print("\n")

		} else if split[0] == "goto" {
			newI, err := strconv.Atoi(currentLine[5:])
			
			if err == nil {
				newI -= 2 // For loop adds 1?
				i = newI
			}

		} else if split[0] == "var" {
			if split[2] != "=" {
				fmt.Print("asd")
			} else {
				varNames = append(varNames, split[1])
				varVals = append(varVals, split[3])
			}

        } else {
			fmt.Println("error")
		}
	}
}

func main() {
	parameter := os.Args[1]

	if parameter == "help" {
		fmt.Println("Yay")
	} else {
		fileBytes, err := ioutil.ReadFile(parameter)

		if err == nil {
			fileText := string(fileBytes)
			fileArray := strings.Split(fileText, "\n")

			run(fileArray)
		} else {
			fmt.Println("Could not read file, or invalid command.")
			fmt.Println(err)

			os.Exit(1)
		}
	}
}
