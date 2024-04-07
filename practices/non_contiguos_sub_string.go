package practices

import (
	"strings"
)

/*
x: ABCD
y: AC
=> true
x: ABCDE
y: CA
=> false
*/

func findSubString(x, y string) bool {
	xArr := strings.Split(x, "")
	yArr := strings.Split(y, "")

	return dp(0, 0, xArr, yArr)
}

func dp(i, j int, xArr, yArr []string) bool {
	if j == len(yArr) {
		return true
	}
	if i == len(xArr) {
		return false
	}

	if xArr[i] == yArr[j] {
		return dp(i+1, j+1, xArr, yArr)
	} else {
		return dp(i+1, j, xArr, yArr)
	}
}

/*
PSEUDOCODE

f(x, y) = f(x[0], y) && f(x[1], y) ... f(x[len x - 1], y)
					&& f(x[1], y)...
f(x[0], y) = f(x[0], y[0]) || f(x[0], y[1]) ... f(x[0], y[len y -1])

*/
