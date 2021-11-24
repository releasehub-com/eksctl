// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"bytes"
	"sync"

	"github.com/weaveworks/eksctl/pkg/karpenter/providers"
	"helm.sh/helm/v3/pkg/getter"
)

type FakeGetter struct {
	GetStub        func(string, ...getter.Option) (*bytes.Buffer, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 []getter.Option
	}
	getReturns struct {
		result1 *bytes.Buffer
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *bytes.Buffer
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGetter) Get(arg1 string, arg2 ...getter.Option) (*bytes.Buffer, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 []getter.Option
	}{arg1, arg2})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGetter) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeGetter) GetCalls(stub func(string, ...getter.Option) (*bytes.Buffer, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeGetter) GetArgsForCall(i int) (string, []getter.Option) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGetter) GetReturns(result1 *bytes.Buffer, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *bytes.Buffer
		result2 error
	}{result1, result2}
}

func (fake *FakeGetter) GetReturnsOnCall(i int, result1 *bytes.Buffer, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *bytes.Buffer
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *bytes.Buffer
		result2 error
	}{result1, result2}
}

func (fake *FakeGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGetter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ providers.Getter = new(FakeGetter)
