package practices

import (
	"fmt"
	"testing"
)

func TestFindSubString(t *testing.T) {

	testCases := []struct {
		x   string
		y   string
		res bool
	}{
		{
			x:   "ABCD",
			y:   "AC",
			res: true,
		},
		{
			x:   "ABCADEF",
			y:   "AEF",
			res: true,
		},
		{
			x:   "ABCD",
			y:   "AE",
			res: false,
		},
		{
			x:   "EABCD",
			y:   "AE",
			res: false,
		},
	}

	for _, tc := range testCases {
		actual := findSubString(tc.x, tc.y)
		if actual != tc.res {
			fmt.Println("x: ", tc.x, "\ny: ", tc.y, "\nexpected: ", tc.res)
			t.Error()
		}
	}
}
