package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	cfg := &config{
		next: "",
		previous : "",

	} 

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0{
			continue
		}

		if cmd, exists := commandList[words[0]]; exists {
			if err := cmd.callback(cfg); err != nil{
				fmt.Println("Error: ", err)
			}
		} else{
				fmt.Println("Unknown command")
		}
	}
}