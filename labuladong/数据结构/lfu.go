package main

type LFUCache struct {
	// key到val的映射
	keyToVal map[int]int
	// key到freq的映射
	keyToFreq map[int]int
	// 最小频次
	minFreq int
	// freq到key列表的映射
	freqToKeys map[int][]int
	// 容量
	cap int
}

func Constructor1(capacity int) LFUCache {
	return LFUCache{
		keyToVal:   make(map[int]int),
		keyToFreq:  make(map[int]int),
		minFreq:    0,
		freqToKeys: make(map[int][]int),
		cap:        capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if val, ok := this.keyToVal[key]; !ok {
		return -1
	} else {
		this.increaseFreq(key)
		return val
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}
	if _, ok := this.keyToVal[key]; ok {
		this.keyToVal[key]++
		this.increaseFreq(key)
		return
	}
	// 容量已满，去除最小频次的key
	if this.cap <= len(this.keyToVal) {
		this.removeMinFreqKey()
	}
	this.keyToVal[key] = value
	this.keyToFreq[key]++
	this.freqToKeys[1] = append(this.freqToKeys[1], key)
	this.minFreq = 1
}

func (this *LFUCache) increaseFreq(key int) {
	freq := this.keyToFreq[key]
	this.keyToFreq[key]++
	keys := this.freqToKeys[freq]
	// 删除freq列表中的key
	for i, v := range keys {
		if v == key {
			keys = append(keys[:i], keys[i+1:]...)
		}
	}
	// 将key加入到freq+1的列表中
	this.freqToKeys[freq+1] = append(this.freqToKeys[freq+1], key)
	// 如果freq对应的列表空了，移除freq列表
	if len(keys) == 0 {
		delete(this.freqToKeys, freq)
		if this.minFreq == freq {
			this.minFreq++
		}
	}
}

func (this *LFUCache) removeMinFreqKey() {
	keyList := this.freqToKeys[this.minFreq]
	key := keyList[0]
	keyList = keyList[1:]
	if len(keyList) == 0 {
		delete(this.freqToKeys, this.minFreq)
	} else {
		this.freqToKeys[this.minFreq] = keyList
	}
	delete(this.keyToVal, key)
	delete(this.keyToFreq, key)
}
