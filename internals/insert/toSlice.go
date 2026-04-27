package insert

func SisipAngka(input int, idx int) []int {
	angka := []int{
		50, 75, 66, 20, 32, 90,
	}
	newSlice := make([]int, len(angka))
	copy(newSlice, angka)

	index := idx
	result := make([]int, 0)

	result = append(newSlice[:index], append([]int{input}, newSlice[index:]...)...)
	return result
}
