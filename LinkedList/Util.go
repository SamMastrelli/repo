package Linkedlist

import "strings"

func (l List) String() string {
  var b strings.Builder
  b.WriteString("[")
  for p:=l.first; p!=nil; p=p.next {
    if p!=l.first {
      b.WriteString(", ")
    }
    b.WriteString(p.payload.String())
  }
  b.WriteString("]")
  return b.String()
}

func (l List) Length() int {
  c:=0
  for p:=l.first; p!=nil; p=p.next {
    c++
  }
  return c
}
