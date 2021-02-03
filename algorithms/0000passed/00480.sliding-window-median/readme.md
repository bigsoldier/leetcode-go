#### 题目
<p>中位数是有序序列最中间的那个数。如果序列的大小是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。</p>

<p>例如：</p>

<ul>
	<li><code>[2,3,4]</code>，中位数是&nbsp;<code>3</code></li>
	<li><code>[2,3]</code>，中位数是 <code>(2 + 3) / 2 = 2.5</code></li>
</ul>

<p>给你一个数组 <em>nums</em>，有一个大小为 <em>k</em> 的窗口从最左端滑动到最右端。窗口中有 <em>k</em> 个数，每次窗口向右移动 <em>1</em> 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<p>给出&nbsp;<em>nums</em> = <code>[1,3,-1,-3,5,3,6,7]</code>，以及&nbsp;<em>k</em> = 3。</p>

<pre>窗口位置                      中位数
---------------               -----
[1  3  -1] -3  5  3  6  7       1
 1 [3  -1  -3] 5  3  6  7      -1
 1  3 [-1  -3  5] 3  6  7      -1
 1  3  -1 [-3  5  3] 6  7       3
 1  3  -1  -3 [5  3  6] 7       5
 1  3  -1  -3  5 [3  6  7]      6
</pre>

<p>&nbsp;因此，返回该滑动窗口的中位数数组&nbsp;<code>[1,-1,-1,3,5,6]</code>。</p>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>你可以假设&nbsp;<code>k</code>&nbsp;始终有效，即：<code>k</code> 始终小于输入的非空数组的元素个数。</li>
	<li>与真实值误差在 <code>10 ^ -5</code> 以内的答案将被视作正确答案。</li>
</ul>


 #### 题解
 1、暴力法
 按照滑动窗口来查询数据
 ```go
func medianSlidingWindow(nums []int, k int) (ans []float64) {
	getMidian := func(a []int) float64 {
		slice := make([]int,len(a))
		copy(slice,a)
		sort.Ints(slice)
		if len(slice)%2 == 0 {
			return float64(slice[len(slice)/2-1]+slice[len(slice)/2])/2
		} else {
			return float64(slice[len(slice)/2])
		}
	}
	for i := 0; i < len(nums)-k+1; i++ {
		ans = append(ans, getMidian(nums[i:i+k]))
	}
	return
}
```
 时间复杂度O(nklogk),空间复杂度O(nlogk)
 
 2、双队列
 使用大根堆和小根堆存储中位数的左边和中位数的右边，
 当窗口左移时，移除元素。这时我们不能
 ```go
var small,large *hp
var delayed map[int]int

func medianSlidingWindow(nums []int, k int) (ans []float64) {
	delayed = map[int]int{} // 哈希表，记录延迟删除的元素
	small,large = &hp{},&hp{} // 大根堆，小根堆
	makeBalance := func() {
		// 调整small和large中的元素个数
		if small.size > large.size+1 {
			large.push(-small.pop())
			small.prune()
		} else if small.size < large.size {
			small.push(-large.pop())
			large.prune()
		}
	}

	insert := func(x int) {
		if small.Len() == 0 || x <= -small.IntSlice[0] {
			small.push(-x)
		} else {
			large.push(x)
		}
		makeBalance()
	}
	erase := func(x int) {
		delayed[x]++
		if x <= -small.IntSlice[0] {
			small.size--
			if x == -small.IntSlice[0] {
				small.prune()
			}
		} else {
			large.size--
			if x == large.IntSlice[0] {
				large.prune()
			}
		}
		makeBalance()
	}
	getMedian := func() float64 {
		if k&1 > 0 { // 奇数
			return float64(-small.IntSlice[0])
		}
		return float64(-small.IntSlice[0]+large.IntSlice[0])/2
	}
	for _, num := range nums[:k] {
		insert(num)
	}
	n := len(nums)
	ans = make([]float64,1,n-k+1)
	ans[0] = getMedian()
	for i := k; i < n; i++ {
		insert(nums[i])
		erase(nums[i-k])
		ans = append(ans, getMedian())
	}
	return
}

type hp struct {
	sort.IntSlice
	size int
}

func (h *hp) Push(v interface{}) {h.IntSlice = append(h.IntSlice, v.(int))}
func (h *hp) Pop() interface{} {
	p := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return p
}
func (h *hp) push(x int) {
	h.size++
	heap.Push(h,x)
}
func (h *hp) pop() int {
	h.size--
	return heap.Pop(h).(int)
}
func (h *hp) prune() {
	for h.Len() > 0 {
		top := h.IntSlice[0]
		if h == small {
			top = -top
		}
		if d, has := delayed[top]; has {
			if d > 1 {
				delayed[top]--
			} else {
				delete(delayed,top)
			}
			heap.Pop(h)
		} else {
			break
		}
	}
}
```