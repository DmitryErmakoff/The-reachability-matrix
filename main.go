package main

import (
	"fmt"
	"os"
)

func main() {
	var size int
	var input int
	fmt.Println("Введите размер матрицы: ")
	fmt.Scanln(&size)
	err := CheckValid(size)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Выберите вид ввода\n1: Рандом\n2: Вручную")
	fmt.Scanln(&input)
	m := FillMatrix(input, size)
	fmt.Println("Наша матрица")
	matrix := m
	PrintMatrix(m)
	m = AddMatrix(EdinichniyMatrix(size), m)
	for i := 0; i < size; i++ {
		m = AddMatrix(m, MulMatrix(m, matrix))
	}
	fmt.Println("Итоговая матрица")
	PrintMatrix(Final(m))
	fmt.Println("---------------")
	PrintFinal(Final(m))

}
