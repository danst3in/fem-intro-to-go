package main

/*
import (
	"fmt"
)

 func main() {
	var testSentence = "This is new and weird."
	// create a variable to store a slice of bytes equal 
	// to the characters in the string
	b := []byte(testSentence)
	fmt.Println(b)
	for idx, char := range b {
		if idx%2 > 0 {
			// convert each byte int to a string 
			c := string(char)
		
			// fmt.Println(char)
		
			// fmt.Println(c)
			fmt.Println("The Index", idx, "Character(s)", c)
		}
	}
} */

//refactor below!
import "fmt"

func main() {

	mySentence := "A testy sentence."

	for idx,char := range mySentence{
		if idx%2>0{

			fmt.Println(string(char))
		}
	}
}