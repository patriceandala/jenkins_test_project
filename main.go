package main

func main() {

}

func BubbleSort(array []int) []int {
	isSorted := false
	counter := 0
	for !isSorted {
		isSorted = true
		for i := 0; i < len(array)-1-counter; i++ {
			if array[i] > array[i+1] {
				swap(array, i, i+1)
				isSorted = false
			}
		}
		counter++
	}

	return array
}

func swap(array []int, i int, j int) {
	array[i], array[j] = array[j], array[i]
}