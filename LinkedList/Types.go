package Linkedlist

import "fmt"

type Lengthable interface {
  fmt.Stringer()
  Lenght() int
}

type Node struct {
  next *Node
  payload Lengthable
}
type List struct {
  first *Node
}
