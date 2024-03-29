package practices

func stockSimilarities(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	occ := make(map[int][]int)

	for i, v := range a {
		v2 := b[i]
		_, ok := occ[v]
		if ok {
			occ[v][0]++
		} else {
			occ[v] = []int{1, 0}
		}
		_, ok = occ[v2]
		if ok {
			occ[v2][1]++
		} else {
			occ[v2] = []int{0, 1}
		}
	}

	exceed := true
	for _, v := range occ {
		if v[0]-v[1] > 3 || v[1]-v[0] > 3 {
			exceed = false
			break
		}
	}

	return exceed
}
