package main

import (
	"fmt"
	"math"
	"slices"
)

// ********* Advent of Code 2024 *********
// --- Day 1: Historian Hysteria ---
// https://adventofcode.com/2024/day/1

func main() {

	//  ******* Day 1 *******
	var slice1 = []int{3, 4, 2, 1, 3, 3}
	var slice2 = []int{4, 3, 5, 3, 9, 3}

	answer := day1_ver1(slice1, slice2)
	fmt.Println(answer)

	answer = day1_ver2(slice1, slice2)
	fmt.Println(answer)

	//  ******* Day 2 *******
	var reports = []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	answer2 := day2(reports)
	fmt.Println(answer2)
}

// Вариант 1 - без сортировки исходных списков.
// Для выполнения условия задачи достаточно найти разницу между суммами первого и второго списков.
func day1_ver1(list1 []int, list2 []int) float64 {

	sum1 := 0
	sum2 := 0

	for _, v := range list1 {
		sum1 += v
	}

	for _, v := range list2 {
		sum2 += v
	}

	result := math.Abs(float64(sum1) - float64(sum2))

	return result
}

// Вариант 2 - с предварительной сортировкой исходных списков.
// Такой вариант реализует описание алгоритма решения, указанного в постановке задачи.
func day1_ver2(list1 []int, list2 []int) float64 {

	slices.Sort(list1)
	slices.Sort(list2)

	var result float64 = 0

	for i := range list1 {
		result += math.Abs(float64(list1[i]) - float64(list2[i]))
	}

	return result
}

func day2(list []string) int {

	var safeCount int = 0

	return safeCount
}
