package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(frog([]int{30, 10, 60, 10, 60, 50}))
}

func frog(arr []int) (minimum int) {

	if len(arr) == 2 {
		return abs(arr[0] - arr[1])
	}

	for index := 0; index <= len(arr)-2; index++ {

		plusOne := abs(arr[index] - arr[index+1])
		plusTwo := abs(arr[index] - arr[index+2])
		step, value := min(plusOne, plusTwo)
		minimum += value
		index += step

		if len(arr)-3 == index && step == 0 {
			minimum += abs(arr[index+1] - arr[index+2])
			break
		}

	}

	return minimum
}

func min(valueOne, valueTwo int) (step, value int) {
	if valueOne < valueTwo {
		return 0, valueOne
	}
	return 1, valueTwo

}

func abs(value int) int {
	if value < 0 {
		return -(value)
	} else {
		return value
	}
}

func simpleEquation(a, b, c int) (result string) {

	if a < 1 {
		return "No Solution"
	}

	if c > 10000 {
		return "No Solution"
	}

	A := a - 3
	B := 1
	C := 1

	temp := a - (A + B + C)

	if temp < 0 {
		A = -(temp)
	} else if temp == 0 {
		A = 1
	}

	for indexI := A; indexI < 100; indexI++ {
		for indexJ := B; indexJ < 100; indexJ++ {
			for indexK := C; indexK < 100; indexK++ {
				if a == indexI+indexJ+indexK {
					if b == indexI*indexJ*indexK {
						if c == (indexI*indexI)+(indexJ*indexJ)+(indexK*indexK) {
							result = strconv.Itoa(indexI) + "/" + strconv.Itoa(indexJ) + "/" + strconv.Itoa(indexK)
						} else if c < indexI*indexJ*indexK {
							break
						}
					} else if b < indexI*indexJ*indexK {
						break
					}
				} else if a < indexI+indexJ+indexK {
					break
				}
			}
		}
	}

	return result
}

func fibonacciBottomUp(value int) int {

	array := []int{0, 1}

	if value == 0 {
		return value
	} else if value == 1 {
		return value
	}

	for index := 2; index <= value; index++ {
		array = append(array, array[index-1]+array[index-2])
	}

	return array[value]
}

func fibonacciTopDown(value int) int {
	if value <= 1 {
		return value
	}
	return fibonacciTopDown(value-1) + fibonacciTopDown(value-2)
}

func binarySearch(arr []int, search int) (value int) {
	temp := arr[:]
	limit := int(math.Sqrt(float64(len(arr))))

	if arr[0] == search {
		return search
	}

	for index := 0; index <= limit; index++ {
		middle := int(len(temp) / 2)
		if temp[middle] > search {
			temp = temp[:middle]
		} else if temp[middle] < search {
			temp = temp[middle:]
		} else if temp[middle] == search {
			value = search
		}
	}

	if value == 0 {
		value = -1
	}

	return value
}

func tatakae(dragon, knight []int) string {

	if len(dragon) > len(knight) {
		return "Y O U  L O S E"
	}

	sort.Ints(dragon)
	sort.Ints(knight)

	startIndex := 0
	minimum := 0

	for _, value := range dragon {
		for index := startIndex; index < len(knight); index++ {
			if value <= knight[index] {
				minimum += knight[index]
				startIndex = index
				break
			} else if index < len(knight) && value > knight[index] {
				return "YOU LOSE"
			}
		}
	}

	return strconv.Itoa(minimum)
}

func moneyCoins(input int) (arr []int) {
	coins := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	returnCoin := []int{}
	for _, value := range coins {
		for 0 <= input-value {
			returnCoin = append(returnCoin, value)
			input -= value
		}
	}

	return returnCoin
}
