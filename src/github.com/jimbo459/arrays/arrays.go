package main

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numsToSum ...[]int) (sums []int) {
	for _, slice := range numsToSum {
		sums = append(sums, Sum(slice))
	}
	return
}
