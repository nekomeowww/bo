package bo

import (
	"context"
	"sync"
)

var _ lifeCycler = Hook{}

type lifeCycler interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type Hook struct {
	OnStart func(ctx context.Context) error
	OnStop  func(ctx context.Context) error
}

func (l Hook) Start(ctx context.Context) error {
	if l.OnStart == nil {
		return nil
	}

	return l.OnStart(ctx)
}

func (l Hook) Stop(ctx context.Context) error {
	if l.OnStop == nil {
		return nil
	}

	return l.OnStop(ctx)
}

type LifeCycle interface {
	Append(hook Hook)
}

type lifeCycle struct {
	hooks []lifeCycler

	mutex sync.Mutex
}

func (l *lifeCycle) GetHooks() []lifeCycler {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return append([]lifeCycler{}, l.hooks...) // 返回一个拷贝，避免外部修改
}

func (l *lifeCycle) Append(hook Hook) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.hooks = append(l.hooks, hook)
}

type EmptyLifeCycle struct{}

func (*EmptyLifeCycle) Append(Hook) {}

func newLifeCycle() *lifeCycle {
	return &lifeCycle{
		hooks: make([]lifeCycler, 0),
	}
}

func NewEmptyLifeCycle() LifeCycle {
	return &EmptyLifeCycle{}
}
