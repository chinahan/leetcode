//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例：
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// Related Topics 链表
package leetcode_test

import (
	"testing"
	//"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := initListNode([]int{1, 2, 4})
	for item := *l1; ; {
		t.Log(item.Val)
		if item.Next == nil {
			break
		} else {
			item = *item.Next
		}
	}
	l2 := initListNode([]int{1, 3, 4})
	for item := *l2; ; {
		t.Log(item.Val)
		if item.Next == nil {
			break
		} else {
			item = *item.Next
		}
	}
	l := mergeTwoLists(l1, l2)
	for item := *l; ; {
		t.Log(item.Val)
		if item.Next == nil {
			break
		} else {
			item = *item.Next
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func initListNode(args []int) *ListNode {
	if args == nil || len(args) <= 0 {
		return nil
	}
	result := ListNode{args[0], nil}
	result.Next = initListNode(args[1:])

	return &result
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

//leetcode submit region end(Prohibit modification and deletion)
