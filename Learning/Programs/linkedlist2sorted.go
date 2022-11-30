package main

func linksort(l1 *Listnode, l2 *Listnode) *Listnode {
	var dummy = new(Listnode)
	var p = dummy
	for l1 != nil && l2 != nil {
		if l1.val < l2.val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next

		}
		p = p.Next

	}
	if l1 != nil {
		p.Next = l1

	} else {
		p.Next = l2
	}
	return dummy.Next
}

func main() {

}
