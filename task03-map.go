package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	for _, value := range input {
		result = append(result, value)
	}
	sort.Strings(result)
	return
}
