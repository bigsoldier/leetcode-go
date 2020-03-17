#### 题目
<p>给定两个大小为 m 和 n 的有序数组&nbsp;<code>nums1</code> 和&nbsp;<code>nums2</code>。</p>

<p>请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为&nbsp;O(log(m + n))。</p>

<p>你可以假设&nbsp;<code>nums1</code>&nbsp;和&nbsp;<code>nums2</code>&nbsp;不会同时为空。</p>

<p><strong>示例 1:</strong></p>

<pre>nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
</pre>

<p><strong>示例 2:</strong></p>

<pre>nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
</pre>


 #### 题解
 1、暴力法
 将两个数组合并，两个有序数组合并成有序数组，然后根据奇偶数返回中位数。
 ```go
 func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
 	m := len(nums1)
 	n := len(nums2)
 	var nums = make([]int,m+n)
 	// nums1 为空，返回nums2的中位数
 	if m == 0 {
 		if n % 2 == 0 {
 			return float64(nums2[n/2-1] + nums2[n/2]) / 2.0
 		} else {
 			return float64(nums2[n/2])
 		}
 	}
 
 	// nums2 为空，返回nums1的中位数
 	if n == 0 {
 		if m % 2 == 0 {
 			return float64(nums1[m/2-1] + nums1[m/2]) / 2.0
 		} else {
 			return float64(nums1[m/2])
 		}
 	}
 
 	// 将两数组合并
 	var i,j,count int
 	for count < m+n {
 		if i == m {
 			for j < n {
 				nums[count] = nums2[j]
 				j++
 				count++
 			}
 			break
 		}
 		if j == n {
 			for i < m {
 				nums[count] = nums1[i]
 				i++
 				count++
 			}
 			break
 		}
 
 		if nums1[i] < nums2[j] {
 			nums[count] = nums1[i]
 			i++
 			count++
 
 		} else {
 			nums[count] = nums2[j]
 			j++
 			count++
 		}
 	}
 	if count%2 == 0 {
 		return float64(nums[count/2-1] + nums[count/2])/2.0
 	} else {
 		return float64(nums[count/2])
 	}
 }
 ```
 ![暴力法](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-01-14/000401.png)
 时间复杂度：O(m+n),遍历了整个数组
 空间复杂度：O(m+n)，开辟了一个新数组存放
 虽然成功执行，但时间复杂度并不符合题意。
 
 2、优化暴力法
 我们可以得到总数组的长度为两个数组之和，那么获取中位数就为：如果长度为偶数，获取中间两位；如果长度为奇数，获取中间一位。
 那么就可以简化为 最多遍历 （总长度的一半+1）次，记录最后两次的遍历结果。
 ```go
 func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
 	len1, len2 := len(nums1), len(nums2)
 	sumLen := len1 + len2
 	i, j := 0, 0
 	var left,right int
 	for k := 0; k < sumLen/2+1; k++ {
 		left = right
 		if i < len1 && (j>=len2 || nums1[i] < nums2[j]) {
 			right = nums1[i]
 			i++
 			continue
 		} else {
 			right = nums2[j]
 			j++
 		}
 	}
 
 	if sumLen & 1 == 0 {
 		return float64(right + left) / 2.0
 	} else {
 		return float64(right)
 	}
 }
 ```
 ![优化暴力法](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-01-14/000402.png)
 时间复杂度：O((m+n)/2+1) = O(m+n)
 空间复杂度：O(1)
 
 3、再优化
 上面两种方法时间复杂度都达不到 *O(log(m+n))* ,只有二分法可以达到。
 将两个数组合并，中位数假定为 *k* 。
 |   Left   |    C  |    Right  |
 | ---- | ---- | ---- |
 | a~1~ a~2~ ... a~i~     |   c~i~   |  a~i+1~ ... a~m~    |
 | b~1~ b~2~ ... b~j~     |   c~j~   |  b~j+1~ ... b~n~    |
 显然，L1Left <= L1Right, L2Left <= L2Right。
 什么时候满足两个数组取到中位数呢？左边的元素个数加起来刚好等于k，那么 *max(a~j~,b~j~)就是要找的元素。也就是*L1Left <= L2Right, L2Left <= L1Right。
 但这收到奇偶数的影响，如果切到数字上，L1Left=L1Right，这就不好运算。所以对数组填充 ‘#’，使总数变成 2m+1 + 2n+1，肯定为偶数。
 第一个数组 c1 切割，第二个数组切割 c2=m+n-c1。
 它满足这个式子：c1+c2+1 = m+n+1
 对数组切一刀，使数组一份为二
 **例子**
 L1:2 / 3 5
 L2:1 4 / 7 9 
 此时，L1Left = 1，L1Right=3，L2Left=4，L2Right=5。
 这个例子很显然没有达到，那么我们就要移动刀片，将L1的刀片向右移动，L2向左移动
 此时变成：
 L1:2 3 / 5
 L2:1 / 4 7 9 
 这样就符合了，我们取左侧最大的和右侧最小的，那么第三个元素为3。
 ```go
 func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
 	m,n := len(nums1),len(nums2)
 	if m > n { // 保证 nums1 最短
 		return findMedianSortedArrays(nums2,nums1)
 	}
 
 	var nums1Left,nums1Right,nums2Left,nums2Right int
 	var right = m *2
 	var left = 0
 	for left <= right {
 		c1 := (left + right) / 2
 		c2 := m + n - c1
 
 		if c1 != 0 {
 			nums1Left = nums1[(c1-1)/2]
 		} else {
 			nums1Left = math.MinInt64
 		}
 		if c1 != 2 * m {
 			nums1Right = nums1[c1/2]
 		} else {
 			nums1Right = math.MaxInt64
 		}
 		if c2 != 0 {
 			nums2Left = nums2[(c2-1)/2]
 		} else {
 			nums2Left = math.MinInt64
 		}
 		if c2 != 2 * n {
 			nums2Right = nums2[c2/2]
 		} else {
 			nums2Right = math.MaxInt64
 		}
 
 		if nums1Left > nums2Right {
 			right = c1 - 1
 		} else if nums2Left > nums1Right {
 			left = c1 + 1
 		} else {
 			break
 		}
 	}
 
 	return (math.Max(float64(nums1Left),float64(nums2Left)) + math.Min(float64(nums1Right),float64(nums2Right)))/2.0
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-01-14/000403.png)
 
 时间复杂度：O(log(m+n))
 
 空间复杂度：O(1)