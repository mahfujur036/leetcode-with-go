// Solution is pretty straight-forward. Used Euclidean Algorithm for finding 
// out GCD of every pair of elements and inserted the result in the middle
// https://www.geeksforgeeks.org/euclidean-algorithms-basic-and-extended/

/*
* type ListNode struct {
*  	Val  int
*	 Next *ListNode
* }
*/

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		gcdResult := gcd(current.Val, current.Next.Val)
		newNode := &ListNode{Val: gcdResult, Next: current.Next}
		current.Next = newNode
		current = newNode.Next
	}
	return head
}
