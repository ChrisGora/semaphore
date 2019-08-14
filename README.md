# :vertical_traffic_light: Semaphore

Semaphore provides a simple POSIX-style implementation of semaphores. Internally it uses buffered channels, but the exposed methods should be familiar to all C programmers.

## Install

You can install this library using `go get -u github.com/ChrisGora/semaphore`.

## Example usage

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
