package code

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
	item     int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter, false, 0}
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
