package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//ввод массива
	fmt.Println("Введите элементы массива через пробел:")
	var numbers []int
	var numStr string

	for {
		_, err := fmt.Scan(&numStr)
		if err != nil {
			break
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			break
		}

		numbers = append(numbers, num)
	}

	//выбор сортировки
	fmt.Println("Выберите тип сортировки:")
	fmt.Println("1. По возрастанию (стандартная сортировка)")
	fmt.Println("2. По убыванию (стандартная сортировка)")
	fmt.Println("3. Сортировка слиянием")
	fmt.Println("4. Сортировка подсчетом")
	fmt.Println("5. Быстрая сортировка")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		sort.Ints(numbers)
	case 2:
		sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	case 3:
		mergeSort(numbers)
	case 4:
		countingSort(numbers)
	case 5:
		quickSort(numbers, 0, len(numbers)-1)
	default:
		sort.Ints(numbers)
	}

	//вывод
	fmt.Println("Отсортированный массив:", numbers)
}

// мердж сорт
func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	mergeSort(left)
	mergeSort(right)

	merge(arr, left, right)
}

func merge(arr, left, right []int) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}

// подсчетом
func countingSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}

	counts := make([]int, maxVal+1)

	for _, val := range arr {
		counts[val]++
	}

	index := 0
	for i, count := range counts {
		for j := 0; j < count; j++ {
			arr[index] = i
			index++
		}
	}
}

// быстрая
func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
