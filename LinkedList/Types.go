package Linkedlist

import "fmt"

type Lengthable interface {
  fmt.Stringer
  
}

type Node struct {
  next *Node
  payload Lengthable
}
type List struct {
  first *Node
}
