package arrays_and_slices

func SumAll(numbersToSum ...[]int) (sums []int) {
	numberOfArguments := len(numbersToSum)
	sums = make([]int, numberOfArguments)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
