package main

import (
    "fmt"
	"log"
	"strconv"
	"math"
)

func main() {
	/*
	1. Написать программу для конвертации рублей в доллары. 
	Программа запрашивает сумму в рублях и выводит сумму 
	в долларах. Курс доллара задайте константой.
	*/
	fmt.Println("======================================")
	fmt.Println("Задание 1")
	fmt.Println("Программа конвертации рублей в доллары")
	conversion()
	
	/*
	2. Даны катеты прямоугольного треугольника. Найти его площадь, периметр и гипотенузу.
	Используйте тип данных float64 и функции из пакета math .
	*/
	fmt.Println("======================================")
	fmt.Println("Задание 2")
	fmt.Println("Найти площадь, периметр и гипотинузу прямоугольного треугольника")

	triangle()

	/*
	3. * Пользователь вводит сумму вклада в банк и годовой процент. 
	Найти сумму вклада через 5 лет.
	*/
	fmt.Println("======================================")
	fmt.Println("Задание 3")
	fmt.Println("Определить сумму вклада через 5 лет")
	vklad()

}

func vklad() {
	var sumRurString string
	fmt.Println("Введите сумму в рублях:")
	fmt.Scan(&sumRurString)

	sumRur, err := strconv.ParseFloat(sumRurString, 64)
	if err != nil {
		log.Fatalln(err)
	}

	var persentYearString string
	fmt.Println("Введите годовой процент:")
	fmt.Scan(&persentYearString)

	persentYear, err := strconv.ParseFloat(persentYearString, 64)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 1; i <= 5; i++ {
		// for j := 1; j <= 12; j++ {
		// 	fmt.Println(i*j)
		// }
		sumRur += sumRur*persentYear/100
		fmt.Println("Сумма вклада через ", i, " год (лет): ", strconv.FormatFloat(sumRur, 'f', 2, 32), " руб.")
	}
}


func triangle() {
	var a float64 = 3.0
	var b float64 = 4.0

	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	p := a + b + c
	s := (a * b)/2.0


	fmt.Println("Площадь треугольника равна: ", s)
	fmt.Println("Гипотинуза треугольника равна: ", c)
	fmt.Println("Периметр треугольника равна: ", p)

}

func conversion() {

	const rateDollar float64 = 62.35
	fmt.Println("Текущий курс доллара составляет: ", rateDollar)

	var sumRurString string
	fmt.Println("Введите сумму в рублях:")
	fmt.Scan(&sumRurString)

	sumRur, err := strconv.ParseFloat(sumRurString, 64)
	if err != nil {
		log.Fatalln(err)
	}

	sumDollar := sumRur/rateDollar


	fmt.Println("Сумма в долларах составляет: ", strconv.FormatFloat(sumDollar, 'f', 2, 32))

}