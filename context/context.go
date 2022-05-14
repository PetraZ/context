package contextimp

import (
	"errors"
	"sync"
	"time"
)

/*
The main purpose of context package is

*/

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

type emptyCtx int

func (ctx emptyCtx) Deadline() (deadline time.Time, ok bool) { return }
func (ctx emptyCtx) Done() <-chan struct{}                   { return nil }
func (ctx emptyCtx) Err() error                              { return nil }
func (ctx emptyCtx) Value(key interface{}) interface{}       { return nil }

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context { return background }

func TODO() Context { return todo }

type cancelCtx struct {
	parent Context
	done   chan struct{}
	err    error
	mu     sync.Mutex
}

func (ctx *cancelCtx) Deadline() (deadline time.Time, ok bool) { return ctx.parent.Deadline() }
func (ctx *cancelCtx) Done() <-chan struct{}                   { return ctx.done }
func (ctx *cancelCtx) Err() error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return ctx.err
}
func (ctx *cancelCtx) Value(key interface{}) interface{} { return ctx.parent.Value(key) }

var Canceled = errors.New("context canceld")

type CancelFunc func()

func WithCancel(parent Context) (Context, CancelFunc) {

	// take parent context and create a new context
	ctx := &cancelCtx{
		parent: parent,
		done:   make(chan struct{}),
	}

	// If current context canceled, we need to do modify it
	cancel := func() {
		if ctx.Err() == nil {
			ctx.cancel(Canceled)
		}
	}

	// If parent context canceled, wee need to modify current context as well
	go func() {
		// something that is blocked util context is done
		select {
		case <-parent.Done():
			ctx.cancel(parent.Err())
		case <-ctx.Done():
		}
	}()

	return ctx, cancel
}

func (ctx *cancelCtx) cancel(err error) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	if ctx.err != nil {
		return
	}
	ctx.err = err
	close(ctx.done)
}
