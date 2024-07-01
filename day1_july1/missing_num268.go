package main

func missingNumber(nums []int) int {
	var observed_sum int
	for _, v := range nums {
		observed_sum += v
	}
	n := len(nums)
	expected_sum := n * (n + 1) / 2
	return expected_sum - observed_sum
}

func main() {

}
