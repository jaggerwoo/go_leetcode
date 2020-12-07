// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:

// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].


package main

func main() {
    nums := []int{2, 7, 11, 15}
    println(towSumMap(nums, 9))
}

func twoSum(nums []int, target int) []int {
    var ary []int
    for i := 0; i < len(nums)-1; i++ {
        for j := i+1; j < len(nums); j++ {
            if (nums[i] + nums[j]) == target {
                ary = append(ary, i, j)
            }
        }
    }
    return ary
}

func towSumMap(nums []int, target int) []int {
    myMap := make(map[int]int)
    ary := make([]int,0)
    for i, num := range nums {
        lastNum := target - num
        _, ok := myMap[lastNum]
        if ok {
            ary = append(ary, i, myMap[lastNum])
            return ary
        }
        myMap[num] = i
    }
    return ary
}