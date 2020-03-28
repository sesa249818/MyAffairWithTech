//9. Find middle element in linked linst.

func middleElementInList(head *ListNode)*Node{
	f,s:=head,head
	for f !=nil && f.Next !=nil{
		f=f.Next.Next
		s=s.Next
	}
	return s
}
