package main

import (
	"fmt"

	"github.com/kienmatu/go-practice/practices"
)

func main() {
	res := practices.LetterCombinations("23")
	fmt.Printf("%s \n", res)

	res2 := practices.LengthOfLIS([]int{
		10, 9, 2, 5, 3, 7, 101, 18,
	})
	fmt.Println(res2)
}
