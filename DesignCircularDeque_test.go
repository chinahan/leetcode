//设计实现双端队列。
//你的实现需要支持以下操作：
//
//
// MyCircularDeque(k)：构造函数,双端队列的大小为k。
// insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
// insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
// deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
// deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
// getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
// getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
// isEmpty()：检查双端队列是否为空。
// isFull()：检查双端队列是否满了。
//
//
// 示例：
//
// MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
//circularDeque.insertLast(1);			        // 返回 true
//circularDeque.insertLast(2);			        // 返回 true
//circularDeque.insertFront(3);			        // 返回 true
//circularDeque.insertFront(4);			        // 已经满了，返回 false
//circularDeque.getRear();  				// 返回 2
//circularDeque.isFull();				        // 返回 true
//circularDeque.deleteLast();			        // 返回 true
//circularDeque.insertFront(4);			        // 返回 true
//circularDeque.getFront();				// 返回 4
//
//
//
//
// 提示：
//
//
// 所有值的范围为 [1, 1000]
// 操作次数的范围为 [1, 1000]
// 请不要使用内置的双端队列库。
//
// Related Topics 设计 队列
package leetcode_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestCircularDeque(t *testing.T) {
	deque := Constructor(5)
	ok := deque.InsertFront(7)
	t.Log(deque.string(), ok)
	ok = deque.InsertLast(0)
	t.Log(deque.string(), ok)
	val := deque.GetFront()
	t.Log(deque.string(), val)
	ok = deque.InsertLast(3)
	t.Log(deque.string(), ok)
	val = deque.GetFront()
	t.Log(deque.string(), val)
	ok = deque.InsertFront(9)
	t.Log(deque.string(), ok)
	val = deque.GetRear()
	t.Log(deque.string(), val)
	val = deque.GetFront()
	t.Log(deque.string(), val)
	val = deque.GetFront()
	t.Log(deque.string(), val)
	ok = deque.DeleteLast()
	t.Log(deque.string(), ok)
	val = deque.GetRear()
	t.Log(deque.string(), val)
}

func (this *MyCircularDeque) string() string {
	if this.head == nil || this.capacity == 0 || this.length <= 0 {
		return ""
	}
	var result strings.Builder
	for node := this.head.Next; node != nil; node = node.Next {
		result.WriteString(strconv.Itoa(node.data))
		result.WriteString(" ")
	}
	return result.String()
}

//leetcode submit region begin(Prohibit modification and deletion)
type Node struct {
	Pre  *Node
	data int
	Next *Node
}

type MyCircularDeque struct {
	//节点数据
	head *Node
	//链容量
	capacity int
	//链长度
	length int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	circularDeque := MyCircularDeque{&Node{}, k, 0}
	return circularDeque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this == nil || this.length >= this.capacity || this.capacity <= 0 {
		return false
	}

	node := Node{data: value}
	if this.length == 0 {
		this.head.Pre = &node
		this.head.Next = &node
	} else {
		this.head.Next.Pre = &node
		node.Next = this.head.Next
		this.head.Next = &node
	}
	this.length = this.length + 1

	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this == nil || this.length >= this.capacity || this.capacity <= 0 {
		return false
	}

	node := Node{data: value}
	if this.length == 0 {
		this.head.Pre = &node
		this.head.Next = &node
	} else {
		node.Pre = this.head.Pre
		this.head.Pre.Next = &node
		this.head.Pre = &node
	}

	this.length = this.length + 1

	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this == nil || this.length <= 0 || this.head.Next == nil {
		return false
	}

	if this.head.Next.Next == nil {
		this.head.Next = nil
		this.head.Pre = nil
		this.length = 0
		return true
	}

	this.head.Next = this.head.Next.Next
	this.head.Next.Pre = nil
	this.length = this.length - 1

	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this == nil || this.length <= 0 || this.head.Pre == nil {
		return false
	}

	if this.head.Pre.Pre == nil {
		this.head.Next = nil
		this.head.Pre = nil
		this.length = 0
		return true
	}

	lastNode := this.head.Pre
	this.head.Pre = lastNode.Pre
	this.head.Pre.Next = nil
	this.length = this.length - 1

	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this == nil || this.head.Next == nil {
		return -1
	}
	return this.head.Next.data
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this == nil || this.head.Pre == nil {
		return -1
	}
	return this.head.Pre.data
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this == nil || this.length == 0 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.capacity > 0 && this.capacity == this.length {
		return true
	}
	return false
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
//leetcode submit region end(Prohibit modification and deletion)
