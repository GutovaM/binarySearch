package main

import (
	"fmt"
	"sync"
)

// binarySearch выполняет бинарный поиск
func binarySearch(arr []int, s int) int {
	var leftPointer = 0             //крайний левый индекс массива
	var rightPointer = len(arr) - 1 //крайний правый индекс массива

	for leftPointer <= rightPointer {
		var midPointer = (leftPointer + rightPointer) / 2 // Находим средний индекс
		var midValue = arr[midPointer]                    // Получаем значение по среднему индексу

		if midValue == s {
			return midPointer // Возвращаем индекс найденного элемента
		} else if midValue < s {
			leftPointer = midPointer + 1 // Сужаем поиск к правой половине
		} else {
			rightPointer = midPointer - 1 // Сужаем поиск к левой половине
		}
	}

	return -1 // Если элемент не найден, возвращаем -1
}

// quickSort выполняет быструю сортировку
func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	pivotIndex := partition(arr, start, end)
	quickSort(arr, start, pivotIndex-1)
	quickSort(arr, pivotIndex+1, end)
}

// partition разбивает массив и возвращает индекс пивота
func partition(arr []int, start, end int) int {
	pivot := arr[end] // Выбираем последний элемент как пивот
	iPivot := start   // Индекс для перемещения элементов

	for i := start; i < end; i++ {
		if arr[i] <= pivot {
			arr[i], arr[iPivot] = arr[iPivot], arr[i] // Меняем местами
			iPivot++                                  // Увеличиваем индекс пивота
		}
	}

	arr[iPivot], arr[end] = arr[end], arr[iPivot] // Ставим пивот на его правильное место
	return iPivot                                 // Возвращаем индекс пивота
}

func main() {
	n := []int{32, 11, 76, 5, 20, 3, 68, 57, 4, 34, 2} // Массив
	searchItems := []int{11, 70, 32, 11, 76, 5}        // Элементы для поиска

	// Выполняем быструю сортировку массива n
	quickSort(n, 0, len(n)-1)

	var wg sync.WaitGroup
	results := make(chan struct {
		value int
		index int
	},
		len(searchItems)) // Канал для хранения результатов

	for _, item := range searchItems {
		wg.Add(1) // Увеличиваем счетчик горутин
		go func(s int) {
			defer wg.Done()             // Уменьшаем счетчик горутин по завершении
			index := binarySearch(n, s) // Выполняем бинарный поиск
			results <- struct {
				value int
				index int
			}{value: s, index: index} // Отправляем результат в канал
		}(item)
	}

	wg.Wait()      // Ожидаем завершения всех горутин
	close(results) // Закрываем канал после завершения всех горутин

	// Выводим результаты поиска
	for result := range results {
		if result.index != -1 {
			fmt.Printf("Элемент - %d: его индекс - %d\n", result.value, result.index)
		} else {
			fmt.Printf("Элемент %d не найден\n", result.value)
		}
	}
}
