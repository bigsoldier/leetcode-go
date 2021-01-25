#### 题目
<p>给定一位研究者论文被引用次数的数组（被引用次数是非负整数）。编写一个方法，计算出研究者的 <em>h&nbsp;</em>指数。</p>

<p><a href="https://baike.baidu.com/item/h-index/3991452?fr=aladdin" target="_blank">h 指数的定义</a>: &ldquo;h 代表&ldquo;高引用次数&rdquo;（high citations），一名科研人员的 h 指数是指他（她）的 （N 篇论文中）<strong>至多</strong>有 h 篇论文分别被引用了<strong>至少</strong> h 次。（其余的&nbsp;<em>N - h&nbsp;</em>篇论文每篇被引用次数<strong>不多于 </strong><em>h </em>次。）&rdquo;</p>

<p>&nbsp;</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> <code>citations = [3,0,6,1,5]</code>
<strong>输出:</strong> 3 
<strong>解释: </strong>给定数组表示研究者总共有 <code>5</code> 篇论文，每篇论文相应的被引用了 <code>3, 0, 6, 1, 5</code> 次。
&nbsp;    由于研究者有 <code>3 </code>篇论文每篇<strong>至少</strong>被引用了 <code>3</code> 次，其余两篇论文每篇被引用<strong>不多于</strong> <code>3</code> 次，所以她的 <em>h </em>指数是 <code>3</code>。</pre>

<p>&nbsp;</p>

<p><strong>说明:&nbsp;</strong>如果 <em>h </em>有多种可能的值，<em>h</em> 指数是其中最大的那个。</p>


 #### 题解
 1、排序
 ```go
func hIndex(citations []int) int {
	quickSort(citations,0,len(citations)-1)
	var i int
	for ; i < len(citations) && citations[len(citations)-i-1] > i; i++ {}
	return i
}

func quickSort(nums []int, left, right int) {
	if left < right {
		mid := partition(nums,left,right)
		quickSort(nums,left,mid-1)
		quickSort(nums,mid+1,right)
	}
}

func partition(nums []int, left, right int) int {
	var i,j = left,left+1
	for idx := j; idx <= right;idx++ {
		if nums[idx] < nums[i] {
			nums[j],nums[idx] = nums[idx],nums[j]
			j++
		}
	}
	nums[i],nums[j-1] = nums[j-1],nums[i]
	return j-1
}
```
 时间复杂度O(nlogn),空间复杂度O(logn),如果使用堆排序则O(1)
 
 2、计数排序
 ```go
func hIndex(citations []int) int {
	n := len(citations)
	papers := make([]int,n+1)
	// 计数
	for _,c := range citations {
		papers[min(n,c)]++
	}
	// 找出最大的k
	k := n
	for s := papers[n]; s < k; s+=papers[k] {
		k--
	}
	return k
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
 时间复杂度O(n),空间复杂度O(n)