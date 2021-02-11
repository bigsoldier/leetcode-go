#### 题目
<p>设计一个找到数据流中第 <code>k</code> 大元素的类（class）。注意是排序后的第 <code>k</code> 大元素，不是第 <code>k</code> 个不同的元素。</p>

<p>请实现 <code>KthLargest</code> 类：</p>

<ul>
	<li><code>KthLargest(int k, int[] nums)</code> 使用整数 <code>k</code> 和整数流 <code>nums</code> 初始化对象。</li>
	<li><code>int add(int val)</code> 将 <code>val</code> 插入数据流 <code>nums</code> 后，返回当前数据流中第 <code>k</code> 大的元素。</li>
</ul>

<p> </p>

<p><strong>示例：</strong></p>

<pre>
<strong>输入：</strong>
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
<strong>输出：</strong>
[null, 4, 5, 5, 8, 8]

<strong>解释：</strong>
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8
</pre>

<p> </p>
<strong>提示：</strong>

<ul>
	<li><code>1 <= k <= 10<sup>4</sup></code></li>
	<li><code>0 <= nums.length <= 10<sup>4</sup></code></li>
	<li><code>-10<sup>4</sup> <= nums[i] <= 10<sup>4</sup></code></li>
	<li><code>-10<sup>4</sup> <= val <= 10<sup>4</sup></code></li>
	<li>最多调用 <code>add</code> 方法 <code>10<sup>4</sup></code> 次</li>
	<li>题目数据保证，在查找第 <code>k</code> 大元素时，数组中至少有 <code>k</code> 个元素</li>
</ul>


 #### 题解
 1、优先队列
 使用一个大小为k的优先队列来存储前k大的元素，其中优先队列的队列的队头为队列中最小的元素。
 
 在单次插入操作中，首先将元素插入到优先队列中，如果优先队列的元素大于k，我们将优先队列的队头元素弹出，保证优先队列的大小为k
 
 ```go
type KthLargest struct {
    sort.IntSlice
    k int
}


func Constructor(k int, nums []int) KthLargest {
    kl := KthLargest{k:k}
    for _,num := range nums {
        kl.Add(num)
    }
    return kl
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this,val)
    if this.Len() > this.k {
        heap.Pop(this)
    }
    return this.IntSlice[0]
}

// 实现堆
func (this *KthLargest) Push(x interface{}) {
    this.IntSlice = append(this.IntSlice,x.(int))
}

func (this *KthLargest) Pop() interface{} {
    a := this.IntSlice[len(this.IntSlice)-1]
    this.IntSlice = this.IntSlice[:len(this.IntSlice)-1]
    return a
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
```
 时间复杂度O(nlogk),空间复杂度O(k)