package minmeetingrooms

import "sort"

func FindRoomNums(input [][]int) int {
	lenInput := len(input)

	sort.Slice(input, func(i, j int) bool {
		start1 := input[i][0]
		start2 := input[j][0]
		return start1 < start2
	})
	// slices.Sort(input)
	count := 1
	for i, slot := range input {
		if i == lenInput-1 {
			break
		}
		nextSlot := input[i+1]
		nextStart, _ := nextSlot[0], nextSlot[1]
		// overlapped, need a new room
		if nextStart < slot[1] {
			count += 1
		}
	}
	return count
}
