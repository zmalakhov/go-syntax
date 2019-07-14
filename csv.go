package main

/**
* Syntax Go. Homework 5
* Zaur Malakhov, dated May 11, 2019
 */

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	//СОЗДАНИЕ ФАЙЛА
	//записи
	records := [][]string{
		{"Наименование", "Код артикула", "Валюта", "Цена", "Доступен для заказа"},
		{"Рюкзак Asgard Р-2401 Бирюза", "Р-2401", "RUB", "2520", "1"},
		{"Рюкзак Asgard Р-455 Синий", "Р-455", "RUB", "2280", "1"},
		{"Рюкзак Asgard Р-5131 Звезды синие-серые", "Р-5131", "RUB", "1200", "1"},
	}

	//создание файла
	file, err := os.Create("file.csv")
	if err != nil {
		log.Fatal("Unable to open output")
	}
	defer file.Close()

	//запись в файл
	w := csv.NewWriter(file)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

	// ЧТЕНИЕ ИЗ ФАЙЛА
	records, errRead := ReadCSV("file.csv")
	if errRead != nil {
		log.Fatal(errRead)
	}

	for _, bag := range records {
		fmt.Println(bag)
	}

}

func ReadCSV(filepath string) ([][]string, error) {
	csvfile, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	records, err := reader.ReadAll()

	return records, nil
}
