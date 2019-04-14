package armstrong

import (
	"fmt"
	"testing"
)

func TestIsArmstrong(t *testing.T) {
	testCases := []struct {
		number      float64
		isArmstrong bool
	}{
		{1, true},
		{153, true},
		{54748, true},
		{548834, true},
		{1741725, true},
		{24678050, true},
		{146511208, true},
		{4679307774, true},
		{32164049650, true},
		{28116440335967, true},
		{10, false},
		{152, false},
		{154, false},
		{146511207, false},
		{146511209, false},
		{4679307773, false},
		{4679307775, false},
		{28116440335966, false},
		{28116440335968, false},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v = %v", testCase.number, testCase.isArmstrong), func(t *testing.T) {
			if got := IsArmstrong(testCase.number); got != testCase.isArmstrong {
				t.Errorf("IsArmstrong(%v) = %v, want %v", testCase.number, got, testCase.isArmstrong)
			}
		})
	}
}
