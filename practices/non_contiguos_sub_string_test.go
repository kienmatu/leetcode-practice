package practices

import (
	"fmt"
	"testing"
	"time"
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
			x:   "ABCD",
			y:   "C",
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
		{
			x:   "EABCEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEED",
			y:   "AD",
			res: true,
		},
	}

	for _, tc := range testCases {
		now := time.Now()
		actual := findSubString(tc.x, tc.y)
		fmt.Println("take: ", time.Since(now))
		fmt.Println("x: ", tc.x, "\ny: ", tc.y, "\nexpected: ", tc.res, "\ngot:", actual)
		if actual != tc.res {
			t.Error()
		}
	}
}

func BenchmarkFindSubString(t *testing.B) {
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
		{
			x:   "EABCEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEED",
			y:   "AD",
			res: true,
		},
	}

	for _, tc := range testCases {
		findSubString(tc.x, tc.y)
	}
}

/* no map
take:  500ns
x:  ABCD
y:  AC
expected:  true
got: true
take:  458ns
x:  ABCADEF
y:  AEF
expected:  true
got: true
take:  291ns
x:  ABCD
y:  AE
expected:  false
got: false
take:  292ns
x:  EABCD
y:  AE
expected:  false
got: false
take:  9.708µs
x:  EABCEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEED
y:  AD
expected:  true
got: true
PASS
*/
