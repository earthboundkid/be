package be_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

type mockingT struct {
	testing.T
	m        sync.Mutex
	didFail  bool
	cleanups []func()
}

func (m *mockingT) setFail(b bool) {
	m.m.Lock()
	defer m.m.Unlock()
	m.didFail = b
}

func (m *mockingT) fail() bool {
	m.m.Lock()
	defer m.m.Unlock()
	return m.didFail
}

func (m *mockingT) Run(name string, f func(t *testing.T)) {
	ch := make(chan struct{})
	// Use a goroutine so Fatalf can call Goexit
	go func() {
		defer func() { m.setFail(false) }()
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
	m.setFail(true)
	fmt.Printf(format+"\n", args...)
}

func (m *mockingT) Failed() bool { return m.fail() }
