//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
// Related Topics 数组 哈希表
package leetcode_test

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	t.Log(result)
	result = twoSum(nums, 26)
	t.Log(result)
	result = twoSum(nums, 13)
	t.Log(result)
	result = twoSum(nums, 18)
	t.Log(result)
}

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	for slowIdx := 0; slowIdx < len(nums)-1; slowIdx = slowIdx + 1 {
		for fastIdx := slowIdx + 1; fastIdx < len(nums); fastIdx = fastIdx + 1 {
			if nums[slowIdx]+nums[fastIdx] == target {
				return []int{slowIdx, fastIdx}
			}
			continue
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
