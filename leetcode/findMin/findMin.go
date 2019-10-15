
/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

You may assume no duplicate exists in the array.

Example 1:

Input: [3,4,5,1,2] 
Output: 1
*/

func findMin(nums []int) int {
	if len(nums) == 0 || nums == nil { return -1 }
	if len(nums) == 1 { return nums[0] }
	return binarySearchRotated(nums)
}

func binarySearchRotated(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (end + start) / 2
		if start == end {
			return nums[start]
		} else if start < mid && nums[mid - 1] > nums[mid] {
			return nums[mid]
		} else if end > mid && nums[mid + 1] < nums[mid] {
			return nums[mid + 1]
		} else if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return 0
}