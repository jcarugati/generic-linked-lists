package linked_list

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	a := assert.New(t)

	secondNode := &node[int]{
		item: 2,
	}

	firstNode := &node[int]{
		item:    1,
		follows: secondNode,
	}

	a.Equal(secondNode, firstNode.next())
}

func TestLink_Add(t *testing.T) {
	a := assert.New(t)

	link := NewSimpleLink[int]()

	for i := 0; i < 10; i++ {
		link.Add(i)
	}

	a.Equal(10, link.length)
}

func TestLink_Next(t *testing.T) {
	a := assert.New(t)

	var numbers []int

	link := NewSimpleLink[int]()

	for i := 0; i < 10; i++ {
		num := rand.Int()
		numbers = append(numbers, num)
		link.Add(num)
	}

	index := 0

	for item := link.Next(); item != nil; item = link.Next() {
		expNum := numbers[index]
		value := item.GetValue()
		a.Equal(expNum, value)
		index++
	}
}
