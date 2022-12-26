package statistics

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func averageFunc(nums []int) int {
	var sum int
	for _, number := range nums {
		if number > 0 {
			sum += number
		}
	}
	fmt.Printf("El promedio es: %v\n", sum/len(nums))
	return sum / len(nums)
}

func findMaximum(nums []int) int {
	maximum := nums[0]
	for _, number := range nums {
		if number > maximum {
			maximum = number
		}
	}
	fmt.Printf("El número mayor entre %v es %d\n", nums, maximum)
	return maximum
}

func findMinimum(nums []int) int {
	minimum := nums[0]
	for _, number := range nums {
		if number < minimum {
			minimum = number
		}
	}
	fmt.Printf("El número menor entre %v es %d\n", nums, minimum)
	return minimum
}

func IdentifyOperation(operation string, nums ...int) {
	switch operation {
	case minimum:
		findMinimum(nums)
	case average:
		averageFunc(nums)
	case maximum:
		findMaximum(nums)
	}
}
