package be

// Panicked runs the callback and returns the recovered panic, if any.
func Panicked(fn func()) (r any) {
	defer func() {
		r = recover()
	}()
	fn()
	return
}
