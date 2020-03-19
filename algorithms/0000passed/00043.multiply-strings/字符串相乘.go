package code

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var result = make([]uint8, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := num1[i] - '0'
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := num2[j] - '0'
			sum := result[i+j+1] + n1*n2
			result[i+j+1] = sum % 10
			result[i+j] += sum / 10
		}
	}

	var ret string
	for i := 0; i < len(result); i++ {
		if i == 0 && result[i] == 0 {
			continue
		}
		ret += string(result[i] + '0')
	}
	return ret
}
