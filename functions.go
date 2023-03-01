package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PopulateRandomValues(size int) [][]int {
	m := make([][]int, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var random int
			go rand.Seed(time.Now().UnixNano())
			random = 0 - rand.Intn(1-0+1)
			m[i] = append(m[i], random)
		}
	}
	return ClearUnits(m)
}

func CheckValid(size int) error {
	if size <= 0 {
		return errors.New("input error")
	} else {
		return nil
	}
}
func PopulateValues(size int) [][]int {
	m := make([][]int, size)
	var symbol = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("Введите [%s %s] элемент матрицы ", symbol[i], symbol[j])
			var input int
			fmt.Scanln(&input)
			m[i] = append(m[i], input)
		}
	}
	return ClearUnits(m)
}

func FillMatrix(input int, size int) [][]int {
	if input == 1 {
		m := PopulateRandomValues(size)
		return m
	} else if input == 2 {
		m := PopulateValues(size)
		return m
	} else {
		fmt.Println("input error")
		os.Exit(1)
		return nil
	}
}

func PrintMatrix(matrix [][]int) {
	var symbol = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}
	size := len(matrix)
	for i := 0; i < size; i++ {
		fmt.Print(" " + symbol[i] + " ")
	}
	fmt.Println("")
	for i := 0; i < size; i++ {
		fmt.Print(symbol[i])
		fmt.Println(matrix[i])
	}
}

func EdinichniyMatrix(size int) [][]int {
	m := make([][]int, size)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if i == j {
				m[i] = append(m[i], 1)
			} else {
				m[i] = append(m[i], 0)
			}
		}
	}
	return ClearUnits(m)
}

func MulMatrix(matrix1 [][]int, matrix2 [][]int) [][]int {
	result := make([][]int, len(matrix1))
	for i := 0; i < len(matrix1); i++ {
		result[i] = make([]int, len(matrix1))
		for j := 0; j < len(matrix2); j++ {
			for k := 0; k < len(matrix2); k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
				if result[i][j] != 0 {
					result[i][j] = 1
				}
			}
		}
	}
	return result
}

func AddMatrix(matrix1 [][]int, matrix2 [][]int) [][]int {
	result := make([][]int, len(matrix1))
	for i, a := range matrix1 {
		for j, _ := range a {
			result[i] = append(result[i], matrix1[i][j]+matrix2[i][j])
		}
	}
	return ClearUnits(result)
}

func ClearUnits(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] != 0 {
				matrix[i][j] = 1
			}
		}
	}
	return matrix
}

func Final(matrix [][]int) [][]int {
	m := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			number := matrix[i][j] * matrix[j][i]
			m[i] = append(m[i], number)
		}
	}
	return ClearUnits(m)
}

func PrintFinal(matrix [][]int) {
	var symbol = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}
	for i := 0; i < len(matrix); i++ {
		fmt.Printf("K%d-{", i+1)
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] == 1 {
				fmt.Printf("%s", symbol[j])
			}

		}
		fmt.Println("}")
	}
}
