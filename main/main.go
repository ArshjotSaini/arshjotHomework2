package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed Fortunes.txt
var fileData string

func main() {
	ch := make(chan string)
	go fortune(ch)

	for {
		fmt.Println("You want another fortune? (y/n)") //asking the user
		var input string                               //creating variable to store response
		fmt.Scanln(&input)                             //putting value inside variable with &

		if strings.ToUpper(input) == "Y" {
			ch <- "Y" //if response is yes weather in samll or capital letter it will send a message down the channel
		} else if strings.ToUpper(input) == "N" {
			break //if response is no loop breaks
		} else {
			fmt.Println("Can't understand, provide a valid response")
		}
	}
	close(ch)
} //main function ends here

/*
This function picks up random text from text file based upon user input
*/
func fortune(c chan string) {
	lines := strings.Split(fileData, "%%")
	for {
		if strings.ToUpper(<-c) == "Y" {
			fmt.Println(lines[rand.Intn(len(lines))])
		}
	}
}
