//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
//
// 示例 1:
//
// 输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
//
//
// 示例 2:
//
// 输入: [4,3,2,1]
//输出: [4,3,2,2]
//解释: 输入数组表示数字 4321。
//
// Related Topics 数组
package leetcode_test

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	digits := []int{4, 3, 2, 1}
	result := plusOne(digits)
	t.Log(result)
	digits = []int{4, 3, 2, 9}
	result = plusOne(digits)
	t.Log(result)
	digits = []int{4, 3, 9, 9}
	result = plusOne(digits)
	t.Log(result)
	digits = []int{4, 9, 9, 9}
	result = plusOne(digits)
	t.Log(result)
	digits = []int{9, 9, 9, 9}
	result = plusOne(digits)
	t.Log(result)
}

//leetcode submit region begin(Prohibit modification and deletion)
func plusOne(digits []int) []int {
	//是否有进位
	var isCarry bool = false
	if digits[len(digits)-1] == 9 {
		isCarry = true
		digits[len(digits)-1] = 0
	} else {
		digits[len(digits)-1] = digits[len(digits)-1] + 1
		return digits
	}

	if isCarry && len(digits) > 1 {
		digits = append(plusOne(digits[:len(digits)-1]), digits[len(digits)-1])
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

//leetcode submit region end(Prohibit modification and deletion)
