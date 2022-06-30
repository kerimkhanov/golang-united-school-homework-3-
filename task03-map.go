package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	intArr := []int{}
	for key, _ := range input {
		intArr = append(intArr, key)
	}
	sort.Ints(intArr)
	for i := 0; i < len(intArr); i++ {
		result = append(result, input[intArr[i]])
	}
	return
}
