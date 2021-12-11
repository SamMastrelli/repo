package Linkedlist

func (l *List) Append(payload Lengthable) {
  if l.first==nil {
    l.first=&Node{payload:payload}
    return
  }
  p:=l.first
  for p.next != nil {
    p=p.next
  }
  p.next=&Node{payload:payload}
}
