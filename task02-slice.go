package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	result = input
	copy(result, input)
	i := 0
	j := len(result) - 1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return
}
