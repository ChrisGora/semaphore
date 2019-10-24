# :vertical_traffic_light: Semaphore

Semaphore provides a simple POSIX-style implementation of semaphores. Internally it uses buffered channels, but the exposed methods should be familiar to all C programmers.

## Install

You can install this library using `go get -u github.com/ChrisGora/semaphore`.

## Basic Usage
### Creating a new semaphore
`semaphore.Init(max, value)` takes a maximum value and initial value for the new semaphore.

```go
maxValue := 3
initValue := 0

sem := semaphore.Init(maxValue, initValue)
```

### Wait / Pend / P
To decrease the value of the semaphore (also known as waiting, pending, etc.):
```go
sem.Wait()
```

### Post / Signal / V
To increase the value of the semaphore (also known as posting, signalling, etc.):
```go
sem.Post()
```

## Example Usage
### Producer-Consumer

The producer-consumer problem can be solved with the semaphores and Go's mutexes in the following way:

```go
func producer(buffer *buffer, spaceAvailable, workAvailable semaphore.Semaphore, mutex *sync.Mutex) {
	for {
		spaceAvailable.Wait()
		mutex.Lock()
		buffer.put(1)
		mutex.Unlock()
		workAvailable.Post()
	}
}

func consumer(buffer *buffer, spaceAvailable, workAvailable semaphore.Semaphore, mutex *sync.Mutex) {
	for {
		workAvailable.Wait()
		mutex.Lock()
		_ = buffer.get()
		mutex.Unlock()
		spaceAvailable.Post()
	}
}
```
