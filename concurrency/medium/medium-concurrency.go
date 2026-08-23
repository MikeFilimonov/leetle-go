package mediumconcurrency

type H2O struct {
	h     chan struct{}
	o     chan struct{}
	water chan struct{}
	done  chan struct{}
}

func NewH2O() *H2O {
	result := &H2O{
		h:     make(chan struct{}, 2),
		o:     make(chan struct{}, 1),
		water: make(chan struct{}, 2), // h releases water
		done:  make(chan struct{}, 2),
	}
	return result
}

func (w *H2O) Hydrogen(releaseHydrogen func()) {

	w.h <- struct{}{}
	<-w.water
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()
	w.done <- struct{}{}

}

func (w *H2O) Oxygen(releaseOxygen func()) {

	w.o <- struct{}{}
	<-w.h
	<-w.h
	// releaseOxygen() outputs "H". Do not change or remove this line.
	releaseOxygen()

	w.water <- struct{}{}
	w.water <- struct{}{}

	<-w.done
	<-w.done
	<-w.o
}
