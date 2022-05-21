package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var a []int
	var out []string
	//Place your code here
	for k := range input {
		a = append(a, k)
	}
	sort.Ints(a)

	for v := range a {
		out = append(out, input[v])
	}

	return out
}
