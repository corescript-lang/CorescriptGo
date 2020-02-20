package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	code := [3]string{"print ¡Hola de Go!", "print Hello from Go!", "goto 1"}
	
	for i := 0; i < len(code); i++ {
		currentLine := code[i]
		split := strings.Fields(currentLine)
		
		if split[0] == "print" {
			fmt.Println(currentLine[6:])
		} else if split[0] == "goto" {
			newI, _ := strconv.Atoi(currentLine[5:])
			newI -= 2 // For loop already adds 1?
			
			i = newI
		}
	}
}

