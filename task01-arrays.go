package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	sum := float32(0.0)
	for i := range input {
		sum += input[i]
	}

	return sum / float32(len(input))
}
