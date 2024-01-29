package rgroup

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

type Group struct {
	err     error
	wg      sync.WaitGroup
	errOnce sync.Once

	workerOnce sync.Once
	ch         chan func(ctx context.Context) error
	chs        []func(ctx context.Context) error

	ctx     context.Context
	cancel  func()
	skipErr bool
}

func WithContext(ctx context.Context) *Group {
	return &Group{ctx: ctx}
}

func WithCancel(ctx context.Context) *Group {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{ctx: ctx, cancel: cancel}
}

func (g *Group) do(f func(ctx context.Context) error) {
	ctx := g.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	var err error
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			err = fmt.Errorf("errgroup: panic recovered: %s\n%s", r, buf)
		}
		if err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
		g.wg.Done()
	}()
	if !g.skipErr {
		select {
		case <-g.ctx.Done():
			fmt.Println("ctx.Done()")
			return
		default:
			err = f(ctx)
		}
	}
	if g.skipErr {
		err = f(ctx)
	}
}

// GOMAXPROCS set max goroutine to work.
func (g *Group) GOMAXPROCS(n int) {
	if n <= 0 {
		panic("errgroup: GOMAXPROCS must great than 0")
	}
	g.workerOnce.Do(func() {
		g.ch = make(chan func(context.Context) error, n)
		for i := 0; i < n; i++ {
			go func() {
				for f := range g.ch {
					g.do(f)
				}
			}()
		}
	})
}

func (g *Group) Go(f func(ctx context.Context) error) {
	g.wg.Add(1)
	if g.ch != nil {
		select {
		case g.ch <- f:
		default:
			g.chs = append(g.chs, f)
		}
		return
	}
	go g.do(f)
}

func (g *Group) Wait() error {
	if g.ch != nil {
		for _, f := range g.chs {
			g.ch <- f
		}
	}
	g.wg.Wait()
	if g.ch != nil {
		close(g.ch) // let all receiver exit
	}
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}

func Finish(ctx context.Context, gnum int, fns ...func(ctx context.Context) error) (err error) {
	g := WithCancel(ctx)
	if gnum > 1 {
		g.GOMAXPROCS(gnum)
	}
	g.skipErr = false

	for _, fn := range fns {
		g.Go(fn)
	}

	if err = g.Wait(); err != nil {
		return
	}
	return
}

func FinishVoidErr(ctx context.Context, gnum int, fns ...func(ctx context.Context) error) {
	g := WithCancel(ctx)
	if gnum > 1 {
		g.GOMAXPROCS(gnum)
	}
	g.skipErr = true
	for _, fn := range fns {
		g.Go(fn)
	}

	if err := g.Wait(); err != nil {
		return
	}
	return
}
