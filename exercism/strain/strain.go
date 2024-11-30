package exstrain

func Keep[T any](arr []T, predicate func(item T) bool) []T {
	res := []T{}
	for _, item := range arr {
		if predicate(item) {
			res = append(res, item)
		}
	}

	return res
}
