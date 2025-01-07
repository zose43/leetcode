package easy

// 26. Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums *[]int) int {
	buf := make([]int, 0, len(*nums))
	for i, num := range *nums {
		if len(buf) == 0 || num != (*nums)[i-1] {
			buf = append(buf, num)
		}
	}
	copy(*nums, buf)
	*nums = (*nums)[:len(buf)]
	return len(buf)
}
