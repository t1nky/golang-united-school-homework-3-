package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var a []string
	//Place your code here
	for _, v := range input {
		a = append(a, v)
	}
	sort.Sort(sort.StringSlice(a))
	return
}
