func minMaxDifference(num int) int {
	var maxDigit string
	var minDigit string
	var maxString string
	var minString string

	for i, v := range strconv.Itoa(num) {
		if i == 0 {
			minDigit = string(v)
		}
		if v != '9' {
			maxDigit = string(v)
			break
		}
	}

	for _, v := range strconv.Itoa(num) {
		if string(v) == maxDigit {
			maxString += "9"
		} else {
			maxString += string(v)
		}

		if string(v) == minDigit {
			minString += "0"
		} else {
			minString += string(v)
		}
	}
	maxNum, _ := strconv.Atoi(maxString)
	minNum, _ := strconv.Atoi(minString)

	return maxNum - minNum
}