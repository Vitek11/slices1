// Прибрати всі дублікати з слайсу int
package main

import "fmt"

func main() {
	arr := []int32{4, 1, 4, -4, 6, 3, 6, 8, 8, 1, 1, 1, 5}
	result := make([]int32, len(arr))
	copy(result, arr)

	for j := 0; j < len(result)-1; j++ {
		for i := j + 1; i < len(result); i++ {
			if result[j] == result[i] {
				result = append(result[:i], result[i+1:]...)
				i--
			}
		}
	}
	fmt.Println(arr)
	fmt.Println(result)
 }
