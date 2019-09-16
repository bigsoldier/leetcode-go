## [Two Sum](https://leetcode-cn.com/problems/two-sum/)

#### **题目**

 Given an array of integers, return indices of the two numbers such that they add up to a specific target. 

 You may assume that each input would have exactly one solution, and you may not use the same element twice. 

 Example: 

``` 
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
``` 
 Related Topics Array Hash Table
 
 #### 题目大意
 在数组中找到2个数字之和等于给定的数字，返回它们在数组的下标
 
 #### 解题思路
 可以暴力破解，两层循环遍历，时间复杂度O（nlogn）
 顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。时间复杂度O（n）