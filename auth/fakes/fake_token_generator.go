// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/concourse/atc/auth"
)

type FakeTokenGenerator struct {
	GenerateTokenStub        func(expiration time.Time) (auth.TokenType, auth.TokenValue, error)
	generateTokenMutex       sync.RWMutex
	generateTokenArgsForCall []struct {
		expiration time.Time
	}
	generateTokenReturns struct {
		result1 auth.TokenType
		result2 auth.TokenValue
		result3 error
	}
}

func (fake *FakeTokenGenerator) GenerateToken(expiration time.Time) (auth.TokenType, auth.TokenValue, error) {
	fake.generateTokenMutex.Lock()
	fake.generateTokenArgsForCall = append(fake.generateTokenArgsForCall, struct {
		expiration time.Time
	}{expiration})
	fake.generateTokenMutex.Unlock()
	if fake.GenerateTokenStub != nil {
		return fake.GenerateTokenStub(expiration)
	} else {
		return fake.generateTokenReturns.result1, fake.generateTokenReturns.result2, fake.generateTokenReturns.result3
	}
}

func (fake *FakeTokenGenerator) GenerateTokenCallCount() int {
	fake.generateTokenMutex.RLock()
	defer fake.generateTokenMutex.RUnlock()
	return len(fake.generateTokenArgsForCall)
}

func (fake *FakeTokenGenerator) GenerateTokenArgsForCall(i int) time.Time {
	fake.generateTokenMutex.RLock()
	defer fake.generateTokenMutex.RUnlock()
	return fake.generateTokenArgsForCall[i].expiration
}

func (fake *FakeTokenGenerator) GenerateTokenReturns(result1 auth.TokenType, result2 auth.TokenValue, result3 error) {
	fake.GenerateTokenStub = nil
	fake.generateTokenReturns = struct {
		result1 auth.TokenType
		result2 auth.TokenValue
		result3 error
	}{result1, result2, result3}
}

var _ auth.TokenGenerator = new(FakeTokenGenerator)
