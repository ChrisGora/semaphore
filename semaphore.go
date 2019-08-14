// Copyright 2019 Chris Gora

// Package semaphore provides a POSIX-style semaphore implementation.
package semaphore

type Semaphore interface {
	Post()
	Wait()
	GetValue() int
}

type semaphore struct {
	max int
	sem chan bool
}

// Init creates a new semaphore with an initial value and some upper bound.
// This is the only difference compared to POSIX semaphores.
func Init(max, value int) Semaphore {
	s := &semaphore{max: max, sem: make(chan bool, max)}
	for i := 0; i < value; i++ {
		s.Post()
	}
	return s
}

// Post increments the value of the semaphore.
func (s *semaphore) Post() {
	s.sem <- true
}

// Wait decrements the value of the semaphore or blocks until that is possible.
func (s *semaphore) Wait() {
	<-s.sem
}

// GetValue returns the value of the semaphore
func (s *semaphore) GetValue() int {
	return len(s.sem)
}
