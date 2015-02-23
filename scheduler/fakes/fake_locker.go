// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/scheduler"
)

type FakeLocker struct {
	AcquireWriteLockStub        func([]db.NamedLock) (db.Lock, error)
	acquireWriteLockMutex       sync.RWMutex
	acquireWriteLockArgsForCall []struct {
		arg1 []db.NamedLock
	}
	acquireWriteLockReturns struct {
		result1 db.Lock
		result2 error
	}
	AcquireWriteLockImmediatelyStub        func([]db.NamedLock) (db.Lock, error)
	acquireWriteLockImmediatelyMutex       sync.RWMutex
	acquireWriteLockImmediatelyArgsForCall []struct {
		arg1 []db.NamedLock
	}
	acquireWriteLockImmediatelyReturns struct {
		result1 db.Lock
		result2 error
	}
	AcquireReadLockStub        func([]db.NamedLock) (db.Lock, error)
	acquireReadLockMutex       sync.RWMutex
	acquireReadLockArgsForCall []struct {
		arg1 []db.NamedLock
	}
	acquireReadLockReturns struct {
		result1 db.Lock
		result2 error
	}
}

func (fake *FakeLocker) AcquireWriteLock(arg1 []db.NamedLock) (db.Lock, error) {
	fake.acquireWriteLockMutex.Lock()
	fake.acquireWriteLockArgsForCall = append(fake.acquireWriteLockArgsForCall, struct {
		arg1 []db.NamedLock
	}{arg1})
	fake.acquireWriteLockMutex.Unlock()
	if fake.AcquireWriteLockStub != nil {
		return fake.AcquireWriteLockStub(arg1)
	} else {
		return fake.acquireWriteLockReturns.result1, fake.acquireWriteLockReturns.result2
	}
}

func (fake *FakeLocker) AcquireWriteLockCallCount() int {
	fake.acquireWriteLockMutex.RLock()
	defer fake.acquireWriteLockMutex.RUnlock()
	return len(fake.acquireWriteLockArgsForCall)
}

func (fake *FakeLocker) AcquireWriteLockArgsForCall(i int) []db.NamedLock {
	fake.acquireWriteLockMutex.RLock()
	defer fake.acquireWriteLockMutex.RUnlock()
	return fake.acquireWriteLockArgsForCall[i].arg1
}

func (fake *FakeLocker) AcquireWriteLockReturns(result1 db.Lock, result2 error) {
	fake.AcquireWriteLockStub = nil
	fake.acquireWriteLockReturns = struct {
		result1 db.Lock
		result2 error
	}{result1, result2}
}

func (fake *FakeLocker) AcquireWriteLockImmediately(arg1 []db.NamedLock) (db.Lock, error) {
	fake.acquireWriteLockImmediatelyMutex.Lock()
	fake.acquireWriteLockImmediatelyArgsForCall = append(fake.acquireWriteLockImmediatelyArgsForCall, struct {
		arg1 []db.NamedLock
	}{arg1})
	fake.acquireWriteLockImmediatelyMutex.Unlock()
	if fake.AcquireWriteLockImmediatelyStub != nil {
		return fake.AcquireWriteLockImmediatelyStub(arg1)
	} else {
		return fake.acquireWriteLockImmediatelyReturns.result1, fake.acquireWriteLockImmediatelyReturns.result2
	}
}

func (fake *FakeLocker) AcquireWriteLockImmediatelyCallCount() int {
	fake.acquireWriteLockImmediatelyMutex.RLock()
	defer fake.acquireWriteLockImmediatelyMutex.RUnlock()
	return len(fake.acquireWriteLockImmediatelyArgsForCall)
}

func (fake *FakeLocker) AcquireWriteLockImmediatelyArgsForCall(i int) []db.NamedLock {
	fake.acquireWriteLockImmediatelyMutex.RLock()
	defer fake.acquireWriteLockImmediatelyMutex.RUnlock()
	return fake.acquireWriteLockImmediatelyArgsForCall[i].arg1
}

func (fake *FakeLocker) AcquireWriteLockImmediatelyReturns(result1 db.Lock, result2 error) {
	fake.AcquireWriteLockImmediatelyStub = nil
	fake.acquireWriteLockImmediatelyReturns = struct {
		result1 db.Lock
		result2 error
	}{result1, result2}
}

func (fake *FakeLocker) AcquireReadLock(arg1 []db.NamedLock) (db.Lock, error) {
	fake.acquireReadLockMutex.Lock()
	fake.acquireReadLockArgsForCall = append(fake.acquireReadLockArgsForCall, struct {
		arg1 []db.NamedLock
	}{arg1})
	fake.acquireReadLockMutex.Unlock()
	if fake.AcquireReadLockStub != nil {
		return fake.AcquireReadLockStub(arg1)
	} else {
		return fake.acquireReadLockReturns.result1, fake.acquireReadLockReturns.result2
	}
}

func (fake *FakeLocker) AcquireReadLockCallCount() int {
	fake.acquireReadLockMutex.RLock()
	defer fake.acquireReadLockMutex.RUnlock()
	return len(fake.acquireReadLockArgsForCall)
}

func (fake *FakeLocker) AcquireReadLockArgsForCall(i int) []db.NamedLock {
	fake.acquireReadLockMutex.RLock()
	defer fake.acquireReadLockMutex.RUnlock()
	return fake.acquireReadLockArgsForCall[i].arg1
}

func (fake *FakeLocker) AcquireReadLockReturns(result1 db.Lock, result2 error) {
	fake.AcquireReadLockStub = nil
	fake.acquireReadLockReturns = struct {
		result1 db.Lock
		result2 error
	}{result1, result2}
}

var _ scheduler.Locker = new(FakeLocker)
