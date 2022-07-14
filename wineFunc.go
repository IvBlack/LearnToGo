package main

import (
	"errors"
	"fmt"
)
	

const winePrice = 100

func main() {
	change, err:= buyWine(20, 110)
	if err != nil {
		fmt.Println("Failed dial:", err.Error())
	} else {
		fmt.Printf("Your change - %d. have nice day!\n", change)
	}

	change, err= buyWine(17, 110)
	if err != nil {
		fmt.Println("Failed dial:", err.Error())
	} else {
		fmt.Printf("Your change - %d. have nice day", change)
	}

	change, err= buyWine(33, 90)
	if err != nil {
		fmt.Println("Failed dial:", err.Error())
	} else {
		fmt.Printf("Your change - %d. have nice day", change)
	}
}

func buyWine(age, moneyAmount int) (int, error) {
	if age >=18 && moneyAmount >= winePrice {
		return moneyAmount - winePrice, nil
	} else if age < 18 {
		return moneyAmount, errors.New("Drink morse, bro!")
	} else {
		return moneyAmount, errors.New("Mommy's money, bro?")
	}
}

/*
	Basic operations with functions:
	if/else, using formatting %d, work with error lib.

	You can run this by: go run wineFunc.go
*/