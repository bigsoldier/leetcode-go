#### 题目
<p>N 对情侣坐在连续排列的 2N 个座位上，想要牵到对方的手。 计算最少交换座位的次数，以便每对情侣可以并肩坐在一起。 <em>一</em>次交换可选择任意两人，让他们站起来交换座位。</p>

<p>人和座位用&nbsp;<code>0</code>&nbsp;到&nbsp;<code>2N-1</code>&nbsp;的整数表示，情侣们按顺序编号，第一对是&nbsp;<code>(0, 1)</code>，第二对是&nbsp;<code>(2, 3)</code>，以此类推，最后一对是&nbsp;<code>(2N-2, 2N-1)</code>。</p>

<p>这些情侣的初始座位&nbsp;&nbsp;<code>row[i]</code>&nbsp;是由最初始坐在第 i 个座位上的人决定的。</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> row = [0, 2, 1, 3]
<strong>输出:</strong> 1
<strong>解释:</strong> 我们只需要交换row[1]和row[2]的位置即可。
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> row = [3, 2, 0, 1]
<strong>输出:</strong> 0
<strong>解释:</strong> 无需交换座位，所有的情侣都已经可以手牵手了。
</pre>

<p><strong>说明:</strong></p>

<ol>
	<li><code>len(row)</code> 是偶数且数值在&nbsp;<code>[4, 60]</code>范围内。</li>
	<li>可以保证<code>row</code> 是序列&nbsp;<code>0...len(row)-1</code>&nbsp;的一个全排列。</li>
</ol>


 #### 题解
 1、并查集
 将每队情侣看成一条边，最终将所有的边连接在一起。
 ```go
func minSwapsCouples(row []int) int {
    n := len(row)
    uf := newUnionFind(n/2)
    for i:=0;i<n;i+=2 {
        uf.union(row[i]/2,row[i+1]/2)
    }
    return n/2-uf.size
}

type unionFind struct {
    fa,rank []int
    size int
}

func newUnionFind(n int) *unionFind {
    fa,rank := make([]int,n),make([]int,n)
    for i := range fa {
        fa[i] = i
        rank[i] = 1
    }
    return &unionFind{fa:fa,rank:rank,size:n}
}

func (u *unionFind) find(x int) int {
    if x != u.fa[x] {
        u.fa[x] = u.find(u.fa[x])
    }
    return u.fa[x]
}

func (u *unionFind) union(x,y int) {
    fx,fy := u.find(x),u.find(y)
    if fx == fy {
        return 
    }
    if u.rank[fx] < u.rank[fy] {
        fx,fy = fy,fx
    }
    u.rank[fx] += u.rank[fy]
    u.fa[fy] = fx
    u.size--
}
```
 时间复杂度O(nlogn),空间复杂度O(n)
 
 2、广度优先
 ```go
func minSwapsCouples(row []int) (ans int) {
    n := len(row)
    graph := make([][]int,n/2)
    for i:=0;i<n;i+=2 {
        l,r := row[i]/2,row[i+1]/2
        if l!=r {
            graph[l] = append(graph[l],r)
            graph[r] = append(graph[r],l)
        }
    }
    visited := make([]bool,n/2)
    for i,v := range visited {
        if !v {
            visited[i] = true
            cnt := 0
            q := []int{i}
            for len(q) > 0 {
                cnt++
                a := q[0]
                q = q[1:]
                for _,w := range graph[a] {
                    if !visited[w] {
                        visited[w] = true
                        q = append(q,w)
                    }
                }
            }
            ans += cnt-1
        }
    }
    return
}
```
 时间复杂度O(n),空间复杂度O(n)