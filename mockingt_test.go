package be_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

type mockingT struct {
	testing.T
	m         sync.Mutex
	hasFailed bool
	cleanups  []func()
}

func (m *mockingT) setFailed(b bool) {
	m.m.Lock()
	defer m.m.Unlock()
	m.hasFailed = b
}

func (m *mockingT) failed() bool {
	m.m.Lock()
	defer m.m.Unlock()
	return m.hasFailed
}

func (m *mockingT) Run(name string, f func(t *testing.T)) {
	m.setFailed(false)
	ch := make(chan struct{})
	// Use a goroutine so Fatalf can call Goexit
	go func() {
		defer func() {
			for _, f := range m.cleanups {
				defer f()
			}
			close(ch)
		}()
		f(&m.T)
	}()
	<-ch
}

func (m *mockingT) Cleanup(f func()) {
	m.cleanups = append(m.cleanups, f)
}

func (*mockingT) Log(args ...any) {
	fmt.Println(args...)
}

func (*mockingT) Helper() {}

func (m *mockingT) Fatalf(format string, args ...any) {
	m.Errorf(format, args...)
	runtime.Goexit()
}

func (m *mockingT) Errorf(format string, args ...any) {
	m.setFailed(true)
	fmt.Printf(format+"\n", args...)
}

func (m *mockingT) Failed() bool { return m.failed() }
