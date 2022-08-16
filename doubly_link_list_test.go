package linked_list

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList_Prev(t *testing.T) {
	a := assert.New(t)
	var arr []int

	link := NewDoublyLink[int]()

	for i := 0; i < 10; i++ {
		num := rand.Int()
		arr = append(arr, num)
		link.Add(num)
	}

	index := 0

	for item := link.Next(); item != nil; item = link.Next() {
		expNum := arr[index]
		value := item.GetValue()
		a.Equal(expNum, value)
		index++
	}

	for item := link.Prev(); item != nil; item = link.Prev() {
		expNum := arr[index]
		value := item.GetValue()
		a.Equal(expNum, value)
		index--
	}
}
