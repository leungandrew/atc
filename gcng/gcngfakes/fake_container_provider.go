// This file was generated by counterfeiter
package gcngfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/gcng"
)

type FakeContainerProvider struct {
	MarkBuildContainersForDeletionStub        func() error
	markBuildContainersForDeletionMutex       sync.RWMutex
	markBuildContainersForDeletionArgsForCall []struct{}
	markBuildContainersForDeletionReturns     struct {
		result1 error
	}
	FindContainersMarkedForDeletionStub        func() ([]*dbng.DestroyingContainer, error)
	findContainersMarkedForDeletionMutex       sync.RWMutex
	findContainersMarkedForDeletionArgsForCall []struct{}
	findContainersMarkedForDeletionReturns     struct {
		result1 []*dbng.DestroyingContainer
		result2 error
	}
	ContainerDestroyStub        func(*dbng.DestroyingContainer) (bool, error)
	containerDestroyMutex       sync.RWMutex
	containerDestroyArgsForCall []struct {
		arg1 *dbng.DestroyingContainer
	}
	containerDestroyReturns struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerProvider) MarkBuildContainersForDeletion() error {
	fake.markBuildContainersForDeletionMutex.Lock()
	fake.markBuildContainersForDeletionArgsForCall = append(fake.markBuildContainersForDeletionArgsForCall, struct{}{})
	fake.recordInvocation("MarkBuildContainersForDeletion", []interface{}{})
	fake.markBuildContainersForDeletionMutex.Unlock()
	if fake.MarkBuildContainersForDeletionStub != nil {
		return fake.MarkBuildContainersForDeletionStub()
	} else {
		return fake.markBuildContainersForDeletionReturns.result1
	}
}

func (fake *FakeContainerProvider) MarkBuildContainersForDeletionCallCount() int {
	fake.markBuildContainersForDeletionMutex.RLock()
	defer fake.markBuildContainersForDeletionMutex.RUnlock()
	return len(fake.markBuildContainersForDeletionArgsForCall)
}

func (fake *FakeContainerProvider) MarkBuildContainersForDeletionReturns(result1 error) {
	fake.MarkBuildContainersForDeletionStub = nil
	fake.markBuildContainersForDeletionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerProvider) FindContainersMarkedForDeletion() ([]*dbng.DestroyingContainer, error) {
	fake.findContainersMarkedForDeletionMutex.Lock()
	fake.findContainersMarkedForDeletionArgsForCall = append(fake.findContainersMarkedForDeletionArgsForCall, struct{}{})
	fake.recordInvocation("FindContainersMarkedForDeletion", []interface{}{})
	fake.findContainersMarkedForDeletionMutex.Unlock()
	if fake.FindContainersMarkedForDeletionStub != nil {
		return fake.FindContainersMarkedForDeletionStub()
	} else {
		return fake.findContainersMarkedForDeletionReturns.result1, fake.findContainersMarkedForDeletionReturns.result2
	}
}

func (fake *FakeContainerProvider) FindContainersMarkedForDeletionCallCount() int {
	fake.findContainersMarkedForDeletionMutex.RLock()
	defer fake.findContainersMarkedForDeletionMutex.RUnlock()
	return len(fake.findContainersMarkedForDeletionArgsForCall)
}

func (fake *FakeContainerProvider) FindContainersMarkedForDeletionReturns(result1 []*dbng.DestroyingContainer, result2 error) {
	fake.FindContainersMarkedForDeletionStub = nil
	fake.findContainersMarkedForDeletionReturns = struct {
		result1 []*dbng.DestroyingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerProvider) ContainerDestroy(arg1 *dbng.DestroyingContainer) (bool, error) {
	fake.containerDestroyMutex.Lock()
	fake.containerDestroyArgsForCall = append(fake.containerDestroyArgsForCall, struct {
		arg1 *dbng.DestroyingContainer
	}{arg1})
	fake.recordInvocation("ContainerDestroy", []interface{}{arg1})
	fake.containerDestroyMutex.Unlock()
	if fake.ContainerDestroyStub != nil {
		return fake.ContainerDestroyStub(arg1)
	} else {
		return fake.containerDestroyReturns.result1, fake.containerDestroyReturns.result2
	}
}

func (fake *FakeContainerProvider) ContainerDestroyCallCount() int {
	fake.containerDestroyMutex.RLock()
	defer fake.containerDestroyMutex.RUnlock()
	return len(fake.containerDestroyArgsForCall)
}

func (fake *FakeContainerProvider) ContainerDestroyArgsForCall(i int) *dbng.DestroyingContainer {
	fake.containerDestroyMutex.RLock()
	defer fake.containerDestroyMutex.RUnlock()
	return fake.containerDestroyArgsForCall[i].arg1
}

func (fake *FakeContainerProvider) ContainerDestroyReturns(result1 bool, result2 error) {
	fake.ContainerDestroyStub = nil
	fake.containerDestroyReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.markBuildContainersForDeletionMutex.RLock()
	defer fake.markBuildContainersForDeletionMutex.RUnlock()
	fake.findContainersMarkedForDeletionMutex.RLock()
	defer fake.findContainersMarkedForDeletionMutex.RUnlock()
	fake.containerDestroyMutex.RLock()
	defer fake.containerDestroyMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeContainerProvider) recordInvocation(key string, args []interface{}) {
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

var _ gcng.ContainerProvider = new(FakeContainerProvider)