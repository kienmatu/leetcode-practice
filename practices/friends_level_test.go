package practices

import (
	"fmt"
	"testing"
)

func TestFriendCircle(t *testing.T) {
	testCases := []struct {
		conn  []string
		name1 string
		name2 string
		res   int
	}{
		{
			conn:  []string{"fred:joe", "joe:mary", "mary:fred", "mary:bill"},
			name1: "fred",
			name2: "bill",
			res:   2,
		},
		{
			conn:  []string{"fred:joe", "joe:mary", "kate:sean", "sean:sally"},
			name1: "kate",
			name2: "sally",
			res:   1,
		},
		{
			conn:  []string{"fred:joe"},
			name1: "fred",
			name2: "joe",
			res:   0,
		},
		{
			conn:  []string{"fred:joe", "joe:mary", "kate:sean", "sean:sally"},
			name1: "joe",
			name2: "larry",
			res:   -1,
		},
		{
			conn:  []string{"fred:joe", "joe:mary", "kate:sean", "sean:sally"},
			name1: "dante",
			name2: "joe",
			res:   -1,
		},
	}

	for _, tc := range testCases {
		actual := friendCircle(tc.conn, tc.name1, tc.name2)
		fmt.Println("conn: ", tc.conn, "\nn1: ", tc.name1, "\nn2: ", tc.name2, "\nexpected:", tc.res, "\ngot:", actual)
		if actual != tc.res {
			t.Error()
		}
	}
}
