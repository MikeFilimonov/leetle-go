type Foo struct {
	starter  chan struct{}
	follower chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		starter:  make(chan struct{}),
		follower: make(chan struct{}),
	}

}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	f.starter <- struct{}{}

}

func (f *Foo) Second(printSecond func()) {

	<-f.starter
	/// Do not change this line
	printSecond()
	f.follower <- struct{}{}

}

func (f *Foo) Third(printThird func()) {

	<-f.follower
	// Do not change this line
	printThird()
}
