package main

import (
	"fmt"
	"strings"
)

func main() {
	//Переменные
	var age int
	var name string
	//Вввод имени
	fmt.Println("*Введите имя:")
	fmt.Scan(&name)
	name = strings.TrimRight(name)
	//Вввод возраста
	fmt.Println("*Введите возраст:")
	fmt.Scan(&age)
	//Вывод результата
	fmt.Println("Ваше имя:", name)
	//Исключения для возраста
	switch age {
	case 0:
		fmt.Println("*Некоректный возраст*")
	case 1, 21, 31, 41, 51, 61, 71, 81, 91:
		fmt.Println("*Ваше возраст:", age, "Год*")
	case 2, 3, 4, 22, 23, 24, 32, 33, 34, 42, 43, 44, 52, 53, 54, 62, 63, 64, 72, 73, 74, 82, 83, 84, 92, 93, 94:
		fmt.Println("*Ваше возраст:", age, "Года*")
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 25, 26, 27, 28, 29, 30, 35, 36, 37, 38, 39, 40, 45, 46, 47, 48, 49, 50, 55, 56, 57, 58, 59, 60, 65, 66, 67, 68, 69, 70, 75, 76, 77, 78, 79, 80, 85, 86, 87, 88, 89, 90, 95, 96, 97, 98, 99, 100:
		fmt.Println("*Ваше возраст:", age, "Лет*")
	default:
		fmt.Println("*Вам не", age, "попробуй еще раз:*")
		fmt.Scan(&age)
		if age < 101 {
			fmt.Println("Ваш возраст", age)
		} else if age > 100 {
			fmt.Println("*Врешь*")
		}
	}
