package practices

import "testing"

func Test_StockSimilarities(t *testing.T) {
	type testcase struct {
		a       []int
		b       []int
		similar bool
	}
	testcases := []testcase{
		{
			a:       []int{10, 10, 20, 10, 10, 20},
			b:       []int{20, 20, 10, 20, 20, 30},
			similar: true,
		},
		{
			a:       []int{10, 10, 10, 10, 10, -20, -20, 30},
			b:       []int{10, -20, -20, -20, -20, 10, -20, -20},
			similar: false,
		},
	}

	for _, tc := range testcases {
		res := stockSimilarities(tc.a, tc.b)
		if res != tc.similar {
			t.Error("Expected", tc.similar, "but got", res)
		}
	}
}
