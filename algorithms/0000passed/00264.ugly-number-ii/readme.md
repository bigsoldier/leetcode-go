#### 题目
<p>编写一个程序，找出第 <code>n</code> 个丑数。</p>

<p>丑数就是只包含质因数&nbsp;<code>2, 3, 5</code> 的<strong>正整数</strong>。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> n = 10
<strong>输出:</strong> 12
<strong>解释: </strong><code>1, 2, 3, 4, 5, 6, 8, 9, 10, 12</code> 是前 10 个丑数。</pre>

<p><strong>说明:&nbsp;</strong>&nbsp;</p>

<ol>
	<li><code>1</code>&nbsp;是丑数。</li>
	<li><code>n</code>&nbsp;<strong>不超过</strong>1690。</li>
</ol>


 #### 题解
 1、堆
 ```go
func nthUglyNumber(n int) int {
	var h nums
	var set = map[int]bool{}
	var primes = []int{2,3,5}
	heap.Init(&h)
	h.push(1)
	for i := 2;i < n+2;i++ {
		x := h.pop()
		for _,j := range primes {
			if !set[x*j] {
				h.push(x*j)
				set[x*j] = true
			}
		}
		if i-1 == n {
			return x
		}
	}
	return -1
}

type nums []int

func (n nums) Len() int {
	return len(n)
}

func (n nums) Swap(i, j int) {
	n[i],n[j] = n[j],n[i]
}
func (n nums) Less(i, j int) bool {
	return n[i]<n[j]
}

func (n *nums) Push(x interface{}) {
	*n = append(*n, x.(int))
}

func (n *nums) Pop() interface{} {
	p := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return p
}

func (n *nums) push(x int) {
	heap.Push(n,x)
}

func (n *nums) pop() int {
	return heap.Pop(n).(int)
}
```
 
 2、手写小根堆
 ```go
func nthUglyNumber(n int) int {
	var h = NewHeap(n)
	var set = map[int]bool{}
	var primes = []int{2,3,5}
	h.push(1)
	for i := 2;i < n+2;i++ {
		x := h.pop()
		for _,j := range primes {
			if !set[x*j] {
				h.push(x*j)
				set[x*j] = true
			}
		}
		if i-1 == n {
			return x
		}
	}
	return -1
}

type Heapt struct {
	index int // 可存放节点的索引
	nums []int
}

func NewHeap(n int) *Heapt {
	return &Heapt{nums: make([]int,n*3)}
}

func (h *Heapt) push(x int) {
	i := h.index
	for i > 0 {
		parent := (i-1)/2
		if x >= h.nums[parent] {
			break
		}
		h.nums[i] = h.nums[parent]
		i = parent
	}

	h.nums[i] = x
	h.index++
}

func (h *Heapt) pop() int {
	ret := h.nums[0] // 根节点
	h.index--
	x := h.nums[h.index]
	//h.nums[h.index] = ret
	i := 0
	for  {
		left,right := 2*i+1,2*i+2
		if left >= h.index { // 越界
			break
		}
		// 获取较小的边
		if right <= h.index && h.nums[left] > h.nums[right] {
			left = right
		}
		if x <= h.nums[left] {
			break
		}
		h.nums[i] = h.nums[left]
		i = left
	}
	h.nums[i] = x
	return ret
}
```

 3、动态规划
 ```go
func nthUglyNumber(n int) int {
	var nums = make([]int,1690)
	nums[0] = 1
	var i,j,k int
	for p := 1; p < 1690; p++ {
		ugly := min(nums[i]*2,nums[j]*3,nums[k]*5)
		nums[p] = ugly
		if ugly == nums[i]*2 {
			i++
		}
		if ugly == nums[j]*3 {
			j++
		}
		if ugly == nums[k]*5 {
			k++
		}
	}
	return nums[n-1]
}

func min(nums ...int) int {
	var m = nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}
```
```go
func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}
	var res = make([]int,n)
	res[0] = 1
	pos := []int{0,0,0}
	factors := []int{2,3,5}
	candidates := []int{2,3,5}
	for i := 1; i < n; i++ {
		res[i] = min(candidates)
		for j := 0; j < 3; j++ {
			if res[i] == candidates[j] {
				pos[j]++
				candidates[j] = res[pos[j]]*factors[j]
			}
		}
	}
	return res[n-1]
}

func min(nums []int) int {
	var m = nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}
```