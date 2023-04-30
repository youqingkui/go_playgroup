package ch14

import (
	"errors"
	"testing"
)

var ErrLessThanTow = errors.New("n should be not less than 2")
var ErrLargerThenHundred = errors.New("n should be not larger 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, ErrLessThanTow
	}
	if n > 100 {
		return nil, ErrLargerThenHundred

	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestError(t *testing.T) {
	if res, err := GetFibonacci(100); err != nil {
		if err == ErrLessThanTow {
			t.Log("n < 2")
		}
		if err == ErrLargerThenHundred {
			t.Log("n > 100")
		}
	} else {
		t.Log(res)

	}

}
