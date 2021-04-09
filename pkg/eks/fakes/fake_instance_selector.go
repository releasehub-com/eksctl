// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/aws/amazon-ec2-instance-selector/v2/pkg/selector"
	"github.com/weaveworks/eksctl/pkg/eks"
)

type FakeInstanceSelector struct {
	FilterStub        func(selector.Filters) ([]string, error)
	filterMutex       sync.RWMutex
	filterArgsForCall []struct {
		arg1 selector.Filters
	}
	filterReturns struct {
		result1 []string
		result2 error
	}
	filterReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstanceSelector) Filter(arg1 selector.Filters) ([]string, error) {
	fake.filterMutex.Lock()
	ret, specificReturn := fake.filterReturnsOnCall[len(fake.filterArgsForCall)]
	fake.filterArgsForCall = append(fake.filterArgsForCall, struct {
		arg1 selector.Filters
	}{arg1})
	stub := fake.FilterStub
	fakeReturns := fake.filterReturns
	fake.recordInvocation("Filter", []interface{}{arg1})
	fake.filterMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstanceSelector) FilterCallCount() int {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	return len(fake.filterArgsForCall)
}

func (fake *FakeInstanceSelector) FilterCalls(stub func(selector.Filters) ([]string, error)) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = stub
}

func (fake *FakeInstanceSelector) FilterArgsForCall(i int) selector.Filters {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	argsForCall := fake.filterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstanceSelector) FilterReturns(result1 []string, result2 error) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = nil
	fake.filterReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeInstanceSelector) FilterReturnsOnCall(i int, result1 []string, result2 error) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = nil
	if fake.filterReturnsOnCall == nil {
		fake.filterReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.filterReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeInstanceSelector) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInstanceSelector) recordInvocation(key string, args []interface{}) {
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

var _ eks.InstanceSelector = new(FakeInstanceSelector)