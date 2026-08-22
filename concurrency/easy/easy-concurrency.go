package leetconcurrency

import "fmt"

type MetaScheduler struct {
	starter  chan struct{}
	follower chan struct{}
}

func NewMetaScheduler() MetaScheduler {

	return MetaScheduler{
		starter:  make(chan struct{}),
		follower: make(chan struct{}),
	}

}

func (ms *MetaScheduler) first() {

	fmt.Println("first")
	ms.starter <- struct{}{}

}

func (ms *MetaScheduler) second() {

	<-ms.starter
	fmt.Println("second")
	ms.follower <- struct{}{}

}

func (ms *MetaScheduler) third() {

	<-ms.follower
	fmt.Println("third")

}
