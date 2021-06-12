package number

func FindOdds(num int) []int {
	var result []int

	for i := 0; i <= num; i++ {
		if i%2 != 0 {
			result = append(result, i)
		}
	}
	return result
}
