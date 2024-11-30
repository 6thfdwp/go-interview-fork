package detectcycles

import (
	"fmt"

	"github.com/shomali11/go-interview/datastructures/linkedlists/singlylinkedlists"
)

// ContainsCycle checks if the list contains a cycle
// n1 -> n2 -> n3 -> n4 -> n2
//
//		    ^          |
//	 		  |----------|
func ContainsCycle[T comparable](head *singlylinkedlists.SLLNode[T]) bool {
	fastPointer := head
	slowPointer := head
	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next
		fmt.Printf("fast: %v slow: %v \n", fastPointer.Value, slowPointer.Value)
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}
