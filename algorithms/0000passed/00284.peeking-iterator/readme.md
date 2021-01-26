#### 题目
<p>给定一个迭代器类的接口，接口包含两个方法：&nbsp;<code>next()</code>&nbsp;和&nbsp;<code>hasNext()</code>。设计并实现一个支持&nbsp;<code>peek()</code>&nbsp;操作的顶端迭代器 -- 其本质就是把原本应由&nbsp;<code>next()</code>&nbsp;方法返回的元素&nbsp;<code>peek()</code>&nbsp;出来。</p>

<p><strong>示例:</strong></p>

<pre>假设迭代器被初始化为列表&nbsp;<strong><code>[1,2,3]</code></strong>。

调用&nbsp;<strong><code>next() </code></strong>返回 <strong>1</strong>，得到列表中的第一个元素。
现在调用&nbsp;<strong><code>peek()</code></strong>&nbsp;返回 <strong>2</strong>，下一个元素。在此之后调用&nbsp;<strong><code>next() </code></strong>仍然返回 <strong>2</strong>。
最后一次调用&nbsp;<strong><code>next()</code></strong>&nbsp;返回 <strong>3</strong>，末尾元素。在此之后调用&nbsp;<strong><code>hasNext()</code></strong>&nbsp;应该返回 <strong>false</strong>。
</pre>

<p><strong>进阶：</strong>你将如何拓展你的设计？使之变得通用化，从而适应所有的类型，而不只是整数型？</p>


 #### 题解
 调用peek时从iterator中获取，然后保存，下一次调用next()时，看是否peek过
 ```go
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    iterator *Iterator
    isPeeked bool
    item int
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{iter,false,0}
}

func (this *PeekingIterator) hasNext() bool {
    return this.isPeeked || this.iterator.hasNext()
}

func (this *PeekingIterator) next() int {
    if !this.isPeeked {
        return this.iterator.next()
    }
    result := this.item
    this.isPeeked = false
    return result
}

func (this *PeekingIterator) peek() int {
    if !this.isPeeked {
        this.item = this.iterator.next()
        this.isPeeked = true
    }
    return this.item
}
```