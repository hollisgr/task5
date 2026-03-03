package service

import "errors"

type Service struct{}

func New() *Service {
	return &Service{}
}

// BasicSum returns the arithmetic sum of two integers.
func BasicSum(a int, b int) int {
	return a + b
}

// BasicErr returns a "basic err" error if b is true; otherwise, it returns nil.
func BasicErr(b bool) error {
	if b {
		return errors.New("basic err")
	}
	return nil
}

// BasicSlice returns a new slice containing the squares of the elements in the input slice.
func BasicSlice(input []int) []int {
	res := []int{}
	for _, num := range input {
		temp := num * num
		res = append(res, temp)
	}
	return res
}

// BasicMap filters the input map and returns a new map containing only
// the entries where the value is greater than zero.
func BasicMap(input map[string]int) map[string]int {
	res := map[string]int{}
	for k, v := range input {
		if v > 0 {
			res[k] = v
		}
	}
	return res
}
