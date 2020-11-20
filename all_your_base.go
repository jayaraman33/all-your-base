package allyourbase

import "fmt"

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, fmt.Errorf("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, fmt.Errorf("output base must be >= 2")
	}
	value := 0
	for _, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		value = value*inputBase + digit
	}
	if value == 0 {
		return []int{0}, nil
	}

	result := []int{}
	for value > 0 {
		result = append(result, value%outputBase)
		value /= outputBase
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result, nil
}
