package graph

func reverseInPlace(a []*vertex) []*vertex {
	n := len(a)

	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}

	return a
}
