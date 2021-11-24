// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/weaveworks/eksctl/pkg/karpenter/providers"
)

type FakeHelmInstaller struct {
	AddRepoStub        func(string, string) error
	addRepoMutex       sync.RWMutex
	addRepoArgsForCall []struct {
		arg1 string
		arg2 string
	}
	addRepoReturns struct {
		result1 error
	}
	addRepoReturnsOnCall map[int]struct {
		result1 error
	}
	InstallChartStub        func(context.Context, string, string, string, string, map[string]interface{}) error
	installChartMutex       sync.RWMutex
	installChartArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
		arg6 map[string]interface{}
	}
	installChartReturns struct {
		result1 error
	}
	installChartReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHelmInstaller) AddRepo(arg1 string, arg2 string) error {
	fake.addRepoMutex.Lock()
	ret, specificReturn := fake.addRepoReturnsOnCall[len(fake.addRepoArgsForCall)]
	fake.addRepoArgsForCall = append(fake.addRepoArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.AddRepoStub
	fakeReturns := fake.addRepoReturns
	fake.recordInvocation("AddRepo", []interface{}{arg1, arg2})
	fake.addRepoMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHelmInstaller) AddRepoCallCount() int {
	fake.addRepoMutex.RLock()
	defer fake.addRepoMutex.RUnlock()
	return len(fake.addRepoArgsForCall)
}

func (fake *FakeHelmInstaller) AddRepoCalls(stub func(string, string) error) {
	fake.addRepoMutex.Lock()
	defer fake.addRepoMutex.Unlock()
	fake.AddRepoStub = stub
}

func (fake *FakeHelmInstaller) AddRepoArgsForCall(i int) (string, string) {
	fake.addRepoMutex.RLock()
	defer fake.addRepoMutex.RUnlock()
	argsForCall := fake.addRepoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHelmInstaller) AddRepoReturns(result1 error) {
	fake.addRepoMutex.Lock()
	defer fake.addRepoMutex.Unlock()
	fake.AddRepoStub = nil
	fake.addRepoReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmInstaller) AddRepoReturnsOnCall(i int, result1 error) {
	fake.addRepoMutex.Lock()
	defer fake.addRepoMutex.Unlock()
	fake.AddRepoStub = nil
	if fake.addRepoReturnsOnCall == nil {
		fake.addRepoReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addRepoReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmInstaller) InstallChart(arg1 context.Context, arg2 string, arg3 string, arg4 string, arg5 string, arg6 map[string]interface{}) error {
	fake.installChartMutex.Lock()
	ret, specificReturn := fake.installChartReturnsOnCall[len(fake.installChartArgsForCall)]
	fake.installChartArgsForCall = append(fake.installChartArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
		arg6 map[string]interface{}
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	stub := fake.InstallChartStub
	fakeReturns := fake.installChartReturns
	fake.recordInvocation("InstallChart", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.installChartMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHelmInstaller) InstallChartCallCount() int {
	fake.installChartMutex.RLock()
	defer fake.installChartMutex.RUnlock()
	return len(fake.installChartArgsForCall)
}

func (fake *FakeHelmInstaller) InstallChartCalls(stub func(context.Context, string, string, string, string, map[string]interface{}) error) {
	fake.installChartMutex.Lock()
	defer fake.installChartMutex.Unlock()
	fake.InstallChartStub = stub
}

func (fake *FakeHelmInstaller) InstallChartArgsForCall(i int) (context.Context, string, string, string, string, map[string]interface{}) {
	fake.installChartMutex.RLock()
	defer fake.installChartMutex.RUnlock()
	argsForCall := fake.installChartArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeHelmInstaller) InstallChartReturns(result1 error) {
	fake.installChartMutex.Lock()
	defer fake.installChartMutex.Unlock()
	fake.InstallChartStub = nil
	fake.installChartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmInstaller) InstallChartReturnsOnCall(i int, result1 error) {
	fake.installChartMutex.Lock()
	defer fake.installChartMutex.Unlock()
	fake.InstallChartStub = nil
	if fake.installChartReturnsOnCall == nil {
		fake.installChartReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.installChartReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHelmInstaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addRepoMutex.RLock()
	defer fake.addRepoMutex.RUnlock()
	fake.installChartMutex.RLock()
	defer fake.installChartMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHelmInstaller) recordInvocation(key string, args []interface{}) {
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

var _ providers.HelmInstaller = new(FakeHelmInstaller)
