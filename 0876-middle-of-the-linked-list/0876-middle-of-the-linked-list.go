/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
  cur := head
  
  var length int
  for cur != nil {
    
    length++
    
    cur = cur.Next
    
  }
  
  fmt.Println(length)
  
  var target float64
  if length % 2 == 0 {
    target = float64(length) / 2
    target += 1.0
  } else {
    target = float64(length) / 2
  }
  fmt.Println("target: ", target)
  for i := 1; i < int(math.Ceil(target)); i++ {
    head = head.Next
  }
  
  return head
  
}