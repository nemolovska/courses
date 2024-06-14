package main

import (
	"botes/botes"
	"bufio"
	"fmt"
	"os"
)

func main() {

	bots := botes.CreateBotes(10)
	info := "info.txt"

	var stProg string
	fmt.Printf("Оберіть програму: \n 1. Список ID\n 2. Пошук по тексту\n")

	fmt.Scanln(&stProg)

	for {
		switch stProg {

		case "1":
			ProgBotList(bots)

		case "2":
			ProgTextRead(info)

		default:
			fmt.Printf("Спробуте ще.\n")
			fmt.Scan(&stProg)
			continue

		}
		break
	}

}

func ProgBotList(bots []botes.Bote) {

	for _, b := range bots {
		b.Print()
	}
}

func ProgTextRead(s string) {

	lines, err := readTextFromFile(s)

	if err != nil {
		fmt.Println("Помилка при читанні файлу:", err)
		return
	}

	var searchTerm string
	fmt.Println("Введіть рядок для пошуку:")
	fmt.Scanln(&searchTerm)

	results := searchInText(lines, searchTerm)

	if len(results) > 0 {
		fmt.Println("Знайдені рядки:")
		for _, line := range results {
			fmt.Println(line)
		}
	} else {
		fmt.Println("Рядки не знайдено.")
	}

}

func readTextFromFile(info string) ([]string, error) {
	file, err := os.Open(info)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			if s[i+j] != substr[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func searchInText(lines []string, searchTerm string) []string {
	var results []string
	for _, line := range lines {
		if contains(line, searchTerm) {
			results = append(results, line)
		}
	}
	return results
}
