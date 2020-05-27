#### 题目
<p>给你两个有序整数数组&nbsp;<em>nums1 </em>和 <em>nums2</em>，请你将 <em>nums2 </em>合并到&nbsp;<em>nums1&nbsp;</em>中<em>，</em>使 <em>num1 </em>成为一个有序数组。</p>

<p>&nbsp;</p>

<p><strong>说明:</strong></p>

<ul>
	<li>初始化&nbsp;<em>nums1</em> 和 <em>nums2</em> 的元素数量分别为&nbsp;<em>m</em> 和 <em>n </em>。</li>
	<li>你可以假设&nbsp;<em>nums1&nbsp;</em>有足够的空间（空间大小大于或等于&nbsp;<em>m + n</em>）来保存 <em>nums2</em> 中的元素。</li>
</ul>

<p>&nbsp;</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong>
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

<strong>输出:</strong>&nbsp;[1,2,2,3,5,6]</pre>


 #### 题解
 1、合并后排序
 ```go
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-27/008801.png)
 时间复杂度O((m+n)log(m+n)),空间复杂度O(1)
 
 2、双向指针
 ```go
 func merge(nums1 []int, m int, nums2 []int, n int) {
 	var i,j int
 	var newNum = make([]int,m+n)
 	for i < m && j < n {
 		if nums1[i] < nums2[j] {
 			newNum[i+j] = nums1[i]
 			i++
 		} else {
 			newNum[i+j] = nums2[j]
 			j++
 		}
 	}
 	if i != m {
 		for k := i; k < m; k++ {
 			newNum[j+k] = nums1[k]
 		}
 	}
 	if j != n {
 		for k := j; k < n; k++ {
 			newNum[i+k] = nums2[k]
 		}
 	}
 	copy(nums1,newNum)
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-27/008802.png)
 时间复杂度O(m+n),空间复杂度O(m)
 
 3、后置指针
 ```go
 func merge(nums1 []int, m int, nums2 []int, n int) {
 	var i,j = m-1,n-1
 	for i >= 0 && j >= 0 {
 		if nums1[i] > nums2[j] {
 			nums1[i+j+1] = nums1[i]
 			i--
 		} else {
 			nums1[i+j+1] = nums2[j]
 			j--
 		}
 	}
 	if j >= 0 {
 		for k := 0; k <= j; k++ {
 			nums1[k] = nums2[k]
 		}
 	}
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-27/008803.png)
 时间复杂度O(m+n),空间复杂度O(1)