package service

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

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

// Square takes a slice of integers and returns a new slice with their squares,
// processing each element concurrently using goroutines.
func Square(input []int) []int {
	len := len(input)
	res := make([]int, len)
	wg := sync.WaitGroup{}
	for i, num := range input {
		wg.Add(1)
		go func(i int, num int) {
			defer wg.Done()
			res[i] = num * num
		}(i, num)
	}
	wg.Wait()
	return res
}

func TimeToString(t time.Time) string {
	day := t.Day()
	mon := t.Month()
	year := t.Year()
	return fmt.Sprintf("%d-%d-%d", day, mon, year)
}
