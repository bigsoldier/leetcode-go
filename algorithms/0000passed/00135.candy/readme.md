#### 题目
<p>老师想给孩子们分发糖果，有 <em>N</em>&nbsp;个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。</p>

<p>你需要按照以下要求，帮助老师给这些孩子分发糖果：</p>

<ul>
	<li>每个孩子至少分配到 1 个糖果。</li>
	<li>相邻的孩子中，评分高的孩子必须获得更多的糖果。</li>
</ul>

<p>那么这样下来，老师至少需要准备多少颗糖果呢？</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> [1,0,2]
<strong>输出:</strong> 5
<strong>解释:</strong> 你可以分别给这三个孩子分发 2、1、2 颗糖果。
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [1,2,2]
<strong>输出:</strong> 4
<strong>解释:</strong> 你可以分别给这三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。</pre>


 #### 题解
 两次遍历
 
 左遍历:当rating[i-1]<rating[i],i号孩子得到的糖果比i-1得到的多
 右遍历:当rating[i]>rating[i+1],i号孩子得到的糖果比i+1得到的多
 ```go
func candy(ratings []int) (ans int) {
	var left = make([]int,len(ratings))
	for i := 0; i < len(ratings); i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1]+1
		} else {
			left[i] = 1
		}
	}
	var right int
	for i := len(ratings)-1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return ans
}

```
 时间复杂度O(n),空间复杂度O(n)
 
 一次遍历
 
 inc是递增序列长度,dec是递减序列长度.
 - [1,2,3,4,5]是递增序列,糖果从1发到5.
 - [1,2,3,4,3,2]递增序列4>递减序列2,从1发到4，然后再发1,2，最后两位互换一下
 - [1,2,4,3,2,1]递增序列3=递减序列3,从1发到3，再从1发到3，但为了区分评分4和3，所以要给评分4的多给1个
 - [1,2,4,3,4,5]两个递增序列，评分3后出现4，表示又是递增，所以dec=0，而评分3只给1个
 ```go
func candy(ratings []int) int {
	// prev是上一个同学分配到的，ans是总数
	prev,ans,inc,dec := 1,1,1,0
	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				prev = 1
			} else {
				prev ++
			}
			ans += prev
			inc = prev
		}  else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			prev = 1
		}
	}
	return ans
}
```
 时间复杂度O(n),空间复杂度O(1)