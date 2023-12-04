// Implementation of a binary tree that takes advantage of `nil` values for the receiver.
package main

import "fmt"

// IntTree defines a node of the int tree.
type IntTree struct {
	val         int
	left, right *IntTree
}

// Insert a value in the tree.
func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		// Create new instance in case of nil:
		return &IntTree{val: val}
	}

	// Calling Insert from nil instance creates new left or right instance:
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

// Contains returns true if `val` is in the tree, and false otherwise.
//
// The Contains method doesn’t modify the *IntTree, but it is declared with a pointer receiver.
// This demonstrates the rule about supporting a nil receiver. A method with a value receiver
// can’t check for nil and panics if invoked with a nil receiver.
func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func main() {
	var it *IntTree

	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)

	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))
}
