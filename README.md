# Generic Linked Lists Utils
Package that offers generic linked list functionality

## Quick Start


### Simple linked list
```go
package main

import (
	"fmt"
	"math/rand"
)

func Foo() {

	// Instantiate a Simple link of type int
	link := NewSimpleLink[int]()

	// Add Values to the linked list
	for i := 0; i < 10; i++ {
		num := rand.Int()
		link.Add(num)
	}

	// Iterate
	for item := link.Next(); item != nil; item = link.Next() {
		// Get value associated with node
		value := item.GetValue()
		fmt.Println(value)
	}
}
```

### Doubly linked list
```go
package main

import (
	"fmt"
	"math/rand"
)

func Foo() {

	// Instantiate a Simple link of type int
	link := NewDoublyLink[int]()

	// Add Values to the linked list
	for i := 0; i < 10; i++ {
		num := rand.Int()
		link.Add(num)
	}

	// Iterate to tail value
	for item := link.Next(); item != nil; item = link.Next() {
		// Get value associated with node
		value := item.GetValue()
		fmt.Println(value)
	}

	// Iterate back to head value
	for item := link.Prev(); item != nil; item = link.Prev() {
		// Get value associated with node
		value := item.GetValue()
		fmt.Println(value)
	}
}
```
