package main

import (
	"sort"
	"time"
)

type Twitter struct {
	tweets   map[int]tweets // 存储的tweet
	followee map[int][]int  // 存储关注人的id
}

type tweet struct {
	id   int
	time int64
}

type tweets []tweet

func (t tweets) Len() int {
	return len(t)
}

func (t tweets) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t tweets) Less(i, j int) bool {
	return t[i].time < t[j].time
}

/** Initialize your data structure here. */
func Constructor4() Twitter {
	t := make(map[int]tweets)
	f := make(map[int][]int)
	return Twitter{tweets: t, followee: f}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets[userId] = append(this.tweets[userId], tweet{
		id:   tweetId,
		time: time.Now().UnixNano(),
	})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	// 获取本人的tweet
	allTweet := this.tweets[userId]
	// 获取follow的tweet
	for _, id := range this.followee[userId] {
		allTweet = append(allTweet, this.tweets[id]...)
	}
	sort.Sort(allTweet)
	var res = make([]int, 0, 10)
	for i := 0; i < len(allTweet) && i < 10; i++ {
		res = append(res, allTweet[i].id)
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	// 自己不能关注自己
	if followerId == followeeId {
		return
	}
	// 不能重复关注
	for _, id := range this.followee[followerId] {
		if id == followeeId {
			return
		}
	}
	this.followee[followerId] = append(this.followee[followerId], followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	for i, id := range this.followee[followerId] {
		if id == followeeId {
			this.followee[followerId] = append(this.followee[followerId][:i], this.followee[followerId][i+1:]...)
		}
	}
}
