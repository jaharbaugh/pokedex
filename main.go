package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0{
			continue
		}

		if cmd, exists := commandMap[words[0]]; exists {
			if err := cmd.callback(); err != nil{
				fmt.Println("Error: ", err)
			}
		} else{
				fmt.Println("Unknown command")
		}
	}
}