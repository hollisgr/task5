package common

import (
	"errors"
	"testing"
	"testing/quick"
	"time"

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

func TestBasicSumQuick(t *testing.T) {
	f := func(a int, b int) bool {
		res := BasicSum(a, b)
		if res != (a + b) {
			return false
		}
		return true
	}

	err := quick.Check(f, &quick.Config{MaxCount: 10})
	assert.NoError(t, err)
}

func TestSquare(t *testing.T) {
	t.Run("default_test", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		exp := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
		res := Square(input)
		assert.Equal(t, exp, res)
	})
}

func TestTimeToString(t *testing.T) {
	test := time.Date(2025, 12, 1, 0, 0, 0, 0, time.UTC)
	exp := "01-12-2025"
	res := TimeToString(test)
	assert.Equal(t, exp, res)
}
