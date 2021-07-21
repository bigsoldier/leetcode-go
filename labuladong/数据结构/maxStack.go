package main

// 最大频率栈
type FreqStack struct {
	valueToFreq  map[int]int
	freqToValues map[int][]int
	maxFreq      int
}

func Constructor2() FreqStack {
	return FreqStack{
		valueToFreq:  make(map[int]int),
		freqToValues: make(map[int][]int),
	}
}

func (this *FreqStack) Push(val int) {
	freq := this.valueToFreq[val]
	this.valueToFreq[val]++
	this.freqToValues[freq+1] = append(this.freqToValues[freq+1], val)
	this.maxFreq = max(this.maxFreq, freq+1)
}

func (this *FreqStack) Pop() int {
	vals := this.freqToValues[this.maxFreq]
	v := vals[0]
	vals = vals[1:]
	this.valueToFreq[v]--
	if len(vals) == 0 {
		this.maxFreq--
	}
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
