package main

import (
	"fmt"
)

func isVowelSwitch(character rune) {  
	switch character {  
	case 'a', 'e', 'i', 'o', 'u':  
	 fmt.Printf(" %c is vowel\n", character)  
	default:  
	 fmt.Printf(" %c is consonant\n", character)  
	}  
   }  

func isVowel(character rune)  {
	if character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u'{
		fmt.Println("Character is vowel")
	} else {
		fmt.Println("Character is not a vowel")
	}
}

func main()  {
	isVowel('a')
	isVowel('b')
	isVowelSwitch('a')
	isVowelSwitch('b')
}