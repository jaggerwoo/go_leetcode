// Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.

// For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

// Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].

package main

import "fmt"

func main() {
    T := []int{73, 74, 75, 71, 69, 72, 76, 73}
    println(dailyTemperatures2(T))
}

func dailyTemperatures(T []int) []int {
	var result []int
	tLen := len(T)

	for i := 0; i < tLen; i++ {
    max := 0
		for j := i+1; j < tLen; j++ {
			if T[j] > T[i] {
				max = j - i
				break
			}
		}
		result = append(result, max)
	}
	return result
}

func dailyTemperatures2(T []int) []int {
    length := len(T);
    result := make([]int, length)

    //从右向左遍历
    for i := length - 2; i >= 0; i-- {
        for j := i + 1; j < length; j += result[j] {
            if T[j] > T[i] {
                result[i] = j - i
                break
            } else if result[j] == 0 {//遇到0表示后面不会有更大的值，那当然当前值就应该也为0
                result[i] = 0
                break
            }
        }
    }

    return result
}