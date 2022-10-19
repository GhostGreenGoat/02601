import "math/rand"

func Quicksort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	left, right := 0, len(list)-1
	pivot := rand.Int() % len(list)
	list[pivot], list[right] = list[right], list[pivot]
	for i, _ := range list {
		if list[i] < list[right] {
			list[left], list[i] = list[i], list[left]
			left++
		}
	}
	list[left], list[right] = list[right], list[left]
	QuickSort(list[:left])
	QuickSort(list[left+1:])
	return list
}