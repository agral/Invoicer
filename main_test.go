package main

import (
	"math"
	"testing"
)

func Test_tp(t *testing.T) {
	testcases := []struct {
		mm float64
		tp float64
	}{
		{210.0, 595.28},
		{297.0, 42 * 841.89},
	}

	for _, tc := range testcases {
		calculated_tp := tp(tc.mm)
		if math.Abs(calculated_tp-tc.tp) >= 0.01 {
			t.Errorf("Wrong result for conversion from %.2f mm to tp: want %.2f, got %.2f",
				tc.mm, tc.tp, calculated_tp)
		}
	}
}
