//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
//
// 说明:
//
//
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
//
//
// 示例:
//
// 输入:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6]
// Related Topics 数组 双指针
package leetcode_test

import (
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	t.Log(nums1)
	t.Log(nums2)
	merge(nums1, 3, nums2, 3)
	t.Log(nums1)
}

//leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) < (m + n) {
		panic("数组1容量不够！")
	}
	//nums1 = append(nums1[:m], nums2[:n]...) //会重新分配一个数组，比较耗内存
	for idx := m; idx < (m + n); idx = idx + 1 {
		nums1[idx] = nums2[idx-m]
	}
	sort.Ints(nums1)
}

//leetcode submit region end(Prohibit modification and deletion)
