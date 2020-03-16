/**
 *Created by Xie Jian on 2020/2/18 16:35
 */
package _027_Remove_Element

//Given an array nums and a value val, remove all instances of that value in-pla
//ce and return the new length.
//
// Do not allocate extra space for another array, you must do this by modifying
//the input array in-place with O(1) extra memory.
//
// The order of elements can be changed. It doesn't matter what you leave beyond
// the new length.
//
// Example 1:
//
//
//Given nums = [3,2,2,3], val = 3,
//
//Your function should return length = 2, with the first two elements of nums be
//ing 2.
//
//It doesn't matter what you leave beyond the returned length.
//
//
// Example 2:
//
//
//Given nums = [0,1,2,2,3,0,4,2], val = 2,
//
//Your function should return length = 5, with the first five elements of nums c
//ontaining 0, 1, 3, 0, and 4.
//
//Note that the order of those five elements can be arbitrary.
//
//It doesn't matter what values are set beyond the returned length.
//
// Clarification:
//
// Confused why the returned value is an integer but your answer is an array?
//
// Note that the input array is 0000passed in by reference, which means modification
// to the input array will be known to the caller as well.
//
// Internally you can think of this:
//
//
//// nums is 0000passed in by reference. (i.e., without making a copy)
//int len = removeElement(nums, val);
//
//// any modification to nums in your function would be known by the caller.
//// using the length returned by your function, it prints the first len element
//s.
//for (int i = 0; i < len; i++) {
//    print(nums[i]);
//} Related Topics Array Two Pointers

//leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	var i int
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

//leetcode submit region end(Prohibit modification and deletion)
