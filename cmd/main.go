package main

import (
	calculator "devops/domain"
	"fmt"
	"log"
)

func main() {

	fmt.Println("Привет пользователь!")
	fmt.Print("Введите два числа: ")
	var x, y float64
	fmt.Scanf("%g %g", &x, &y)

	fmt.Println("Введите номер операции")
	fmt.Print("(1 - Add, 2 - Sub, 3 - Mul, 4 - Div, 5 - Pow): ")

	var op int
	fmt.Scanf("%d", &op)

	var res float64
	switch op {
	case 1:
		res = calculator.Add(x, y)
	case 2:
		res = calculator.Sub(x, y)
	case 3:
		res = calculator.Mul(x, y)
	case 4:
		res = calculator.Div(x, y)
	case 5:
		res = calculator.Pow(x, y)
	default:
		log.Fatal("Нет операции с таким номером")
	}

	fmt.Println(res)
}
