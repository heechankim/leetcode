func duplicateZeros(arr []int)  {
    length := len(arr)
	idx := 0
	for {
		if idx >= length {
			break
		}

		if arr[idx] == 0 {
			for i := length - 1; i > idx; i-- {
				arr[i] = arr[i-1]
			}
			idx++
		}
		idx++
	}
}