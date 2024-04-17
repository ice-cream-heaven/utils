package wait

import (
	"github.com/ice-cream-heaven/utils/routine"
	"sync"
)

type Group struct {
	w   sync.WaitGroup
	max int
	c   chan struct{}
}

func (p *Group) Add(logic func()) {
	p.w.Add(1)

	if p.max > 0 {
		p.c <- struct{}{}
	}

	routine.GoWithRecover(func() (err error) {
		defer func() {
			if p.max > 0 {
				<-p.c
			}

			p.w.Done()
		}()

		logic()
		return nil
	})
}

func (p *Group) Wait() {
	p.w.Wait()
}

func (p *Group) WaitAndExit() {
	p.Wait()
}

func NewGroup(max ...int) *Group {
	p := &Group{}

	if len(max) > 0 {
		p.max = max[0]
	}

	if p.max > 0 {
		p.c = make(chan struct{}, p.max)
	}

	return p
}
