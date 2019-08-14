# :vertical_traffic_light: Semaphore

Semaphore provides a simple POSIX-style implementation of semaphores. Internally it uses buffered channels, but the exposed methods should be familiar to all C programmers.

## Install

You can install this library using `go get -u github.com/ChrisGora/semaphores`.

## Example usage

The producer-consumer problem can be solved with the semaphores and Go's mutexes in the following way:

```go
func producer(buffer *buffer, spaceAvailable, workAvailable semaphore.Semaphore, mutex *sync.Mutex, start, delta int) {
	x := start
	for {
		spaceAvailable.Wait()
		mutex.Lock()
		buffer.put(x)
		mutex.Unlock()
		workAvailable.Post()
		x = x + delta
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func consumer(buffer *buffer, spaceAvailable, workAvailable semaphore.Semaphore, mutex *sync.Mutex) {
	for {
		workAvailable.Wait()
		mutex.Lock()
		_ = buffer.get()
		mutex.Unlock()
		spaceAvailable.Post()
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
	}
}
```
