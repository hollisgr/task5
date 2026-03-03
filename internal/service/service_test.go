package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicSum(t *testing.T) {
	a := 1
	b := 2
	exp := 3
	res := BasicSum(a, b)
	assert.Equal(t, exp, res)
}

func TestBasicErr(t *testing.T) {
	t.Run("no_error", func(t *testing.T) {
		err := BasicErr(false)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		expErr := errors.New("basic err")
		err := BasicErr(true)
		assert.Error(t, err, expErr)
	})
}

func TestBasicSlice(t *testing.T) {
	input := []int{
		1, 2, 3,
	}
	exp := []int{
		1, 4, 9,
	}

	res := BasicSlice(input)
	assert.Equal(t, exp, res)
}

func TestBasicMap(t *testing.T) {
	input := map[string]int{
		"cat":  3,
		"dog":  4,
		"fish": -1,
		"bird": 0,
	}
	exp := map[string]int{
		"cat": 3,
		"dog": 4,
	}

	res := BasicMap(input)
	assert.Equal(t, exp, res)
}
