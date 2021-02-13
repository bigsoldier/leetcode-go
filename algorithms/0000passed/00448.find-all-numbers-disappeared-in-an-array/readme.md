#### 题目
<p>给定一个范围在&nbsp; 1 &le; a[i] &le; <em>n</em> (&nbsp;<em>n</em> = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。</p>

<p>找到所有在 [1, <em>n</em>] 范围之间没有出现在数组中的数字。</p>

<p>您能在不使用额外空间且时间复杂度为<em>O(n)</em>的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。</p>

<p><strong>示例:</strong></p>

<pre>
<strong>输入:</strong>
[4,3,2,7,8,2,3,1]

<strong>输出:</strong>
[5,6]
</pre>


 #### 题解
 1、排序
 
 2、哈希表
 
 3、原地修改
 每遇到一个数字，就将数字加n，然后遍历数字时是按照序号来遍历的，找出没有被加n的数字。
 ```go
func findDisappearedNumbers(nums []int) (ans []int) {
    n := len(nums)
    for _,v := range nums {
        v = (v-1)%n
        nums[v] += n
    }
    for i,v := range nums {
        if v <= n {
            ans = append(ans,i+1)
        }
    }
    return 
}
```
 时间复杂度O(n),空间复杂度O(1)