# 🏃🏻‍♀️ Race conditions

When working with concurrency, it's possible modify shared resources that different GoRoutines can access.
A ["race condition"](https://en.wikipedia.org/wiki/Race_condition#In_software) issue may result from this, and it happens when two or more threads have simultaneous access to common data and attempt to modify it. 
It`s not possible to predict the sequence in which the threads will attempt to access the shared data since the thread scheduling method might switch between threads at any time.

This code illustrate a shared resource, `Account` and a set of GoRoutine who access to `Account`.
You can check possible race conditions build up the source code with the race flag, like this:
```bash
go build --race main.go
```
Then, it will generate some warnings.

## 👩🏻‍🔬 Fixing race conditions

The package sync offers the [Mutex type](https://go.dev/tour/concurrency/9) which acts as a lock to prohibit concurrent access to a piece of code by several GoRoutines at the same time. Between the `Lock()` and `Unlock()` methods the protection is located.

### Read and Write locks

It's possible to create locks with reading and writing exclusion.
The lock can be held by an arbitrary number of readers or a single writers.

[Read more](https://pkg.go.dev/sync#RWMutex)
