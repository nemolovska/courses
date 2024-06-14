package main

import (
	"fmt"
)

func main() {

	transletion()

}

func transletion() {

	var st map[string]string

	var word string

	st = make(map[string]string)

	st["кіт"] = "cat"
	st["сонце"] = "sun"
	st["мама"] = "mather"
	st["машина"] = "car"
	st["собака"] = "dog"
	st["яблоко"] = "apple"

	for {
		fmt.Println("Введіть слово українською:")

		fmt.Scan(&word)

		translation, exists := st[word]

		if exists {
			fmt.Printf("Переклад: %s\n", translation)
		} else {
			fmt.Println("Такого слова в нашому словнику немає")
		}
	}
}
