package main

import "fmt"

func main() {
	nums := make([]int, 4)
	nums[0] = 2
	nums[1] = 5
	nums[2] = 5
	nums[3] = 11

	fmt.Println(twoSum(nums, 10))

}

func twoSum(nums []int, target int) []int {
	retn := make([]int, 2)
	status := false
	var chanRetn = make(chan int, 2)
	for i, num := range nums {
		go forSearch(i, num, nums, target, chanRetn)
	}

	for i := 0; i < 2; i++ {
		select {
		case p1, _ := <-chanRetn:
			if !status {
				retn[0] = p1
				status = true
			}
			if status {
				retn[1] = p1
			}
		}
	}

	return retn
}

func forSearch(i int, num int, nums []int, target int, chanRetn chan int) {
	for k, num2 := range nums {
		if num+num2 == target && i != k {

			chanRetn <- i

			chanRetn <- k

		}
	}
}
