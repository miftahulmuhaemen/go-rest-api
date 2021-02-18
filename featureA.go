package main

func main() {

}

// func maximumSum(arr []int) int {
// 	result := 0
// 	for index := 0; index < len(arr); index++ {
// 		for jindex := index; jindex < len(arr)-index; jindex++ {

// 		}
// 	}
// }

// func maximumSum(arr []int, k int) (value int) {
// 	for index := 0; index < len(arr)-(k-1); index++ {
// 		result := 0
// 		for indexes := k - 1; indexes >= 0; indexes-- {
// 			result += arr[index+indexes]
// 		}
// 		if result > value {
// 			value = result
// 		}
// 	}

// 	return value
// }

// type Coded struct {
// 	name string
// }

// type chiper interface {
// 	encode() string
// 	decode() string
// }

// func main() {
// 	var coded chiper
// 	coded = Coded{"rizky"}

// 	encoded := coded.encode()
// 	fmt.Println(encoded)
// 	coded = Coded{encoded}

// 	decoded := coded.decode()
// 	fmt.Println(decoded)

// }

// func (coded Coded) encode() (encoded string) {

// 	text := coded.name
// 	shift := 3
// 	for i := 0; i < len(text); i++ {
// 		index := (int(text[i])+shift-97)%26 + 97
// 		if index < 97 {
// 			index += 26
// 		}
// 		encoded += string(index)
// 	}

// 	return encoded
// }

// func (coded Coded) decode() (decoded string) {

// 	text := coded.name
// 	shift := -3
// 	for i := 0; i < len(text); i++ {
// 		index := (int(text[i])+shift-97)%26 + 97
// 		if index < 97 {
// 			index += 26
// 		}
// 		decoded += string(index)
// 	}

// 	return decoded
// }

// func (student Student) average() (average float64) {
// 	for _, value := range student.Score {
// 		average += float64(value)
// 	}
// 	return average / float64(len(student.Score))
// }

// func (student Student) min() (name string, score int) {

// 	list := map[string]int{}

// 	for index, value := range student.Name {
// 		list[value] = student.Score[index]
// 	}

// 	for key, value := range list {
// 		if score == 0 {
// 			score = value
// 		} else if score > value {
// 			score = value
// 			name = key
// 		}
// 	}

// 	return name, score
// }

// func (student Student) max() (name string, score int) {

// 	list := map[string]int{}

// 	for index, value := range student.Name {
// 		list[value] = student.Score[index]
// 	}

// 	for key, value := range list {
// 		if score == 0 {
// 			score = value
// 		} else if score < value {
// 			score = value
// 			name = key
// 		}
// 	}

// 	return name, score
// }

// func (student Student) getMinMaxScore(students []*Student) (maxName string, maxScore int, minName string, minScore int) {

// 	for index, value := range students {
// 		fmt.Println(value)
// 		if index == 0 {
// 			minScore, maxScore = value.Score, value.Score
// 			minName, maxName = value.Name, value.Name
// 		} else {

// 		}
// 	}

// 	return maxName, maxScore, minName, minScore
// }

// func getMinMax(numbers ...*int) (mini, maxi int) {
// 	min := 0
// 	max := 0

// 	for index, num := range numbers {
// 		if index == 0 {
// 			min = *num
// 			max = *num
// 		} else {
// 			if min > *num {
// 				min = *num
// 			}

// 			if max < *num {
// 				max = *num
// 			}
// 		}
// 	}

// 	return min, max
// }

// func swap(a, b *int) {
// 	temp := *a
// 	*a = *b
// 	*b = temp
// }

// -----------

// students := Student{}
// for index := 0; index < 6; index++ {

// 	name := "A" + string(index)
// 	score := 10 + index

// 	fmt.Scan(&name)
// 	fmt.Scan(&score)

// 	students.Name = append(students.Name, name)
// 	students.Score = append(students.Score, score)
// }

// fmt.Print(students.max())
// fmt.Print(*&students[2].Name)
