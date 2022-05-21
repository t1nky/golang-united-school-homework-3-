package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	temp := int64(0);
	length := len(input)
	for i := 0; i < int(length / 2); i++ {
		temp = input[i]
		input[i] = input[length - i]
		input[length - i] = temp
	}

	return input
}
