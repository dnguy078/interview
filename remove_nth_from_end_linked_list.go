package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	preHead := &ListNode{
		Next: head,
	}
	fast, slow := head, preHead
	for i := 0; fast != nil && i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return preHead.Next
}

/*
   public static Node copyList(Node head)
   {
       Node current = head;    // used to iterate over the original list
       Node newList = null;    // head of the new list
       Node tail = null;       // point to the last node in a new list

       while (current != null)
       {
           // special case for the first new node
           if (newList == null)
           {
               newList = new Node(current.data, null);
               tail = newList;
           }
           else {
               tail.next = new Node();
               tail = tail.next;
               tail.data = current.data;
               tail.next = null;
           }
           current = current.next;
       }

       return newList;
   }


   func copyRandomList(head *Node) *Node {
	cur := head
	dummy := &Node{}
	dCur := dummy
	visited := map[*Node]*Node{}
	// copy list first
	for cur != nil {
		n := &Node{Val: cur.Val}
		visited[cur] = n
		dCur.Next = n
		dCur = dCur.Next
		cur = cur.Next
	}
	// copy random pointers
	cur, dCur = head, dummy
	for cur != nil {
		dCur.Next.Random = visited[cur.Random]
		dCur = dCur.Next
		cur = cur.Next
	}
	return dummy.Next
}

*/
