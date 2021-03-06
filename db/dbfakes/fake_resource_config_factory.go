// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/creds"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/db/lock"
)

type FakeResourceConfigFactory struct {
	FindOrCreateResourceConfigStub        func(logger lager.Logger, user db.ResourceUser, resourceType string, source atc.Source, resourceTypes creds.VersionedResourceTypes) (*db.UsedResourceConfig, error)
	findOrCreateResourceConfigMutex       sync.RWMutex
	findOrCreateResourceConfigArgsForCall []struct {
		logger        lager.Logger
		user          db.ResourceUser
		resourceType  string
		source        atc.Source
		resourceTypes creds.VersionedResourceTypes
	}
	findOrCreateResourceConfigReturns struct {
		result1 *db.UsedResourceConfig
		result2 error
	}
	findOrCreateResourceConfigReturnsOnCall map[int]struct {
		result1 *db.UsedResourceConfig
		result2 error
	}
	FindResourceConfigStub        func(logger lager.Logger, resourceType string, source atc.Source, resourceTypes creds.VersionedResourceTypes) (*db.UsedResourceConfig, bool, error)
	findResourceConfigMutex       sync.RWMutex
	findResourceConfigArgsForCall []struct {
		logger        lager.Logger
		resourceType  string
		source        atc.Source
		resourceTypes creds.VersionedResourceTypes
	}
	findResourceConfigReturns struct {
		result1 *db.UsedResourceConfig
		result2 bool
		result3 error
	}
	findResourceConfigReturnsOnCall map[int]struct {
		result1 *db.UsedResourceConfig
		result2 bool
		result3 error
	}
	CleanConfigUsesForFinishedBuildsStub        func() error
	cleanConfigUsesForFinishedBuildsMutex       sync.RWMutex
	cleanConfigUsesForFinishedBuildsArgsForCall []struct{}
	cleanConfigUsesForFinishedBuildsReturns     struct {
		result1 error
	}
	cleanConfigUsesForFinishedBuildsReturnsOnCall map[int]struct {
		result1 error
	}
	CleanConfigUsesForInactiveResourceTypesStub        func() error
	cleanConfigUsesForInactiveResourceTypesMutex       sync.RWMutex
	cleanConfigUsesForInactiveResourceTypesArgsForCall []struct{}
	cleanConfigUsesForInactiveResourceTypesReturns     struct {
		result1 error
	}
	cleanConfigUsesForInactiveResourceTypesReturnsOnCall map[int]struct {
		result1 error
	}
	CleanConfigUsesForInactiveResourcesStub        func() error
	cleanConfigUsesForInactiveResourcesMutex       sync.RWMutex
	cleanConfigUsesForInactiveResourcesArgsForCall []struct{}
	cleanConfigUsesForInactiveResourcesReturns     struct {
		result1 error
	}
	cleanConfigUsesForInactiveResourcesReturnsOnCall map[int]struct {
		result1 error
	}
	CleanConfigUsesForPausedPipelinesResourcesStub        func() error
	cleanConfigUsesForPausedPipelinesResourcesMutex       sync.RWMutex
	cleanConfigUsesForPausedPipelinesResourcesArgsForCall []struct{}
	cleanConfigUsesForPausedPipelinesResourcesReturns     struct {
		result1 error
	}
	cleanConfigUsesForPausedPipelinesResourcesReturnsOnCall map[int]struct {
		result1 error
	}
	CleanConfigUsesForOutdatedResourceConfigsStub        func() error
	cleanConfigUsesForOutdatedResourceConfigsMutex       sync.RWMutex
	cleanConfigUsesForOutdatedResourceConfigsArgsForCall []struct{}
	cleanConfigUsesForOutdatedResourceConfigsReturns     struct {
		result1 error
	}
	cleanConfigUsesForOutdatedResourceConfigsReturnsOnCall map[int]struct {
		result1 error
	}
	CleanUselessConfigsStub        func() error
	cleanUselessConfigsMutex       sync.RWMutex
	cleanUselessConfigsArgsForCall []struct{}
	cleanUselessConfigsReturns     struct {
		result1 error
	}
	cleanUselessConfigsReturnsOnCall map[int]struct {
		result1 error
	}
	AcquireResourceCheckingLockStub        func(logger lager.Logger, resourceUser db.ResourceUser, resourceType string, resourceSource atc.Source, resourceTypes creds.VersionedResourceTypes) (lock.Lock, bool, error)
	acquireResourceCheckingLockMutex       sync.RWMutex
	acquireResourceCheckingLockArgsForCall []struct {
		logger         lager.Logger
		resourceUser   db.ResourceUser
		resourceType   string
		resourceSource atc.Source
		resourceTypes  creds.VersionedResourceTypes
	}
	acquireResourceCheckingLockReturns struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	acquireResourceCheckingLockReturnsOnCall map[int]struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfig(logger lager.Logger, user db.ResourceUser, resourceType string, source atc.Source, resourceTypes creds.VersionedResourceTypes) (*db.UsedResourceConfig, error) {
	fake.findOrCreateResourceConfigMutex.Lock()
	ret, specificReturn := fake.findOrCreateResourceConfigReturnsOnCall[len(fake.findOrCreateResourceConfigArgsForCall)]
	fake.findOrCreateResourceConfigArgsForCall = append(fake.findOrCreateResourceConfigArgsForCall, struct {
		logger        lager.Logger
		user          db.ResourceUser
		resourceType  string
		source        atc.Source
		resourceTypes creds.VersionedResourceTypes
	}{logger, user, resourceType, source, resourceTypes})
	fake.recordInvocation("FindOrCreateResourceConfig", []interface{}{logger, user, resourceType, source, resourceTypes})
	fake.findOrCreateResourceConfigMutex.Unlock()
	if fake.FindOrCreateResourceConfigStub != nil {
		return fake.FindOrCreateResourceConfigStub(logger, user, resourceType, source, resourceTypes)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findOrCreateResourceConfigReturns.result1, fake.findOrCreateResourceConfigReturns.result2
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigCallCount() int {
	fake.findOrCreateResourceConfigMutex.RLock()
	defer fake.findOrCreateResourceConfigMutex.RUnlock()
	return len(fake.findOrCreateResourceConfigArgsForCall)
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigArgsForCall(i int) (lager.Logger, db.ResourceUser, string, atc.Source, creds.VersionedResourceTypes) {
	fake.findOrCreateResourceConfigMutex.RLock()
	defer fake.findOrCreateResourceConfigMutex.RUnlock()
	return fake.findOrCreateResourceConfigArgsForCall[i].logger, fake.findOrCreateResourceConfigArgsForCall[i].user, fake.findOrCreateResourceConfigArgsForCall[i].resourceType, fake.findOrCreateResourceConfigArgsForCall[i].source, fake.findOrCreateResourceConfigArgsForCall[i].resourceTypes
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigReturns(result1 *db.UsedResourceConfig, result2 error) {
	fake.FindOrCreateResourceConfigStub = nil
	fake.findOrCreateResourceConfigReturns = struct {
		result1 *db.UsedResourceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigReturnsOnCall(i int, result1 *db.UsedResourceConfig, result2 error) {
	fake.FindOrCreateResourceConfigStub = nil
	if fake.findOrCreateResourceConfigReturnsOnCall == nil {
		fake.findOrCreateResourceConfigReturnsOnCall = make(map[int]struct {
			result1 *db.UsedResourceConfig
			result2 error
		})
	}
	fake.findOrCreateResourceConfigReturnsOnCall[i] = struct {
		result1 *db.UsedResourceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigFactory) FindResourceConfig(logger lager.Logger, resourceType string, source atc.Source, resourceTypes creds.VersionedResourceTypes) (*db.UsedResourceConfig, bool, error) {
	fake.findResourceConfigMutex.Lock()
	ret, specificReturn := fake.findResourceConfigReturnsOnCall[len(fake.findResourceConfigArgsForCall)]
	fake.findResourceConfigArgsForCall = append(fake.findResourceConfigArgsForCall, struct {
		logger        lager.Logger
		resourceType  string
		source        atc.Source
		resourceTypes creds.VersionedResourceTypes
	}{logger, resourceType, source, resourceTypes})
	fake.recordInvocation("FindResourceConfig", []interface{}{logger, resourceType, source, resourceTypes})
	fake.findResourceConfigMutex.Unlock()
	if fake.FindResourceConfigStub != nil {
		return fake.FindResourceConfigStub(logger, resourceType, source, resourceTypes)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findResourceConfigReturns.result1, fake.findResourceConfigReturns.result2, fake.findResourceConfigReturns.result3
}

func (fake *FakeResourceConfigFactory) FindResourceConfigCallCount() int {
	fake.findResourceConfigMutex.RLock()
	defer fake.findResourceConfigMutex.RUnlock()
	return len(fake.findResourceConfigArgsForCall)
}

func (fake *FakeResourceConfigFactory) FindResourceConfigArgsForCall(i int) (lager.Logger, string, atc.Source, creds.VersionedResourceTypes) {
	fake.findResourceConfigMutex.RLock()
	defer fake.findResourceConfigMutex.RUnlock()
	return fake.findResourceConfigArgsForCall[i].logger, fake.findResourceConfigArgsForCall[i].resourceType, fake.findResourceConfigArgsForCall[i].source, fake.findResourceConfigArgsForCall[i].resourceTypes
}

func (fake *FakeResourceConfigFactory) FindResourceConfigReturns(result1 *db.UsedResourceConfig, result2 bool, result3 error) {
	fake.FindResourceConfigStub = nil
	fake.findResourceConfigReturns = struct {
		result1 *db.UsedResourceConfig
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceConfigFactory) FindResourceConfigReturnsOnCall(i int, result1 *db.UsedResourceConfig, result2 bool, result3 error) {
	fake.FindResourceConfigStub = nil
	if fake.findResourceConfigReturnsOnCall == nil {
		fake.findResourceConfigReturnsOnCall = make(map[int]struct {
			result1 *db.UsedResourceConfig
			result2 bool
			result3 error
		})
	}
	fake.findResourceConfigReturnsOnCall[i] = struct {
		result1 *db.UsedResourceConfig
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForFinishedBuilds() error {
	fake.cleanConfigUsesForFinishedBuildsMutex.Lock()
	ret, specificReturn := fake.cleanConfigUsesForFinishedBuildsReturnsOnCall[len(fake.cleanConfigUsesForFinishedBuildsArgsForCall)]
	fake.cleanConfigUsesForFinishedBuildsArgsForCall = append(fake.cleanConfigUsesForFinishedBuildsArgsForCall, struct{}{})
	fake.recordInvocation("CleanConfigUsesForFinishedBuilds", []interface{}{})
	fake.cleanConfigUsesForFinishedBuildsMutex.Unlock()
	if fake.CleanConfigUsesForFinishedBuildsStub != nil {
		return fake.CleanConfigUsesForFinishedBuildsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanConfigUsesForFinishedBuildsReturns.result1
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForFinishedBuildsCallCount() int {
	fake.cleanConfigUsesForFinishedBuildsMutex.RLock()
	defer fake.cleanConfigUsesForFinishedBuildsMutex.RUnlock()
	return len(fake.cleanConfigUsesForFinishedBuildsArgsForCall)
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForFinishedBuildsReturns(result1 error) {
	fake.CleanConfigUsesForFinishedBuildsStub = nil
	fake.cleanConfigUsesForFinishedBuildsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForFinishedBuildsReturnsOnCall(i int, result1 error) {
	fake.CleanConfigUsesForFinishedBuildsStub = nil
	if fake.cleanConfigUsesForFinishedBuildsReturnsOnCall == nil {
		fake.cleanConfigUsesForFinishedBuildsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanConfigUsesForFinishedBuildsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourceTypes() error {
	fake.cleanConfigUsesForInactiveResourceTypesMutex.Lock()
	ret, specificReturn := fake.cleanConfigUsesForInactiveResourceTypesReturnsOnCall[len(fake.cleanConfigUsesForInactiveResourceTypesArgsForCall)]
	fake.cleanConfigUsesForInactiveResourceTypesArgsForCall = append(fake.cleanConfigUsesForInactiveResourceTypesArgsForCall, struct{}{})
	fake.recordInvocation("CleanConfigUsesForInactiveResourceTypes", []interface{}{})
	fake.cleanConfigUsesForInactiveResourceTypesMutex.Unlock()
	if fake.CleanConfigUsesForInactiveResourceTypesStub != nil {
		return fake.CleanConfigUsesForInactiveResourceTypesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanConfigUsesForInactiveResourceTypesReturns.result1
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourceTypesCallCount() int {
	fake.cleanConfigUsesForInactiveResourceTypesMutex.RLock()
	defer fake.cleanConfigUsesForInactiveResourceTypesMutex.RUnlock()
	return len(fake.cleanConfigUsesForInactiveResourceTypesArgsForCall)
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourceTypesReturns(result1 error) {
	fake.CleanConfigUsesForInactiveResourceTypesStub = nil
	fake.cleanConfigUsesForInactiveResourceTypesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourceTypesReturnsOnCall(i int, result1 error) {
	fake.CleanConfigUsesForInactiveResourceTypesStub = nil
	if fake.cleanConfigUsesForInactiveResourceTypesReturnsOnCall == nil {
		fake.cleanConfigUsesForInactiveResourceTypesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanConfigUsesForInactiveResourceTypesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResources() error {
	fake.cleanConfigUsesForInactiveResourcesMutex.Lock()
	ret, specificReturn := fake.cleanConfigUsesForInactiveResourcesReturnsOnCall[len(fake.cleanConfigUsesForInactiveResourcesArgsForCall)]
	fake.cleanConfigUsesForInactiveResourcesArgsForCall = append(fake.cleanConfigUsesForInactiveResourcesArgsForCall, struct{}{})
	fake.recordInvocation("CleanConfigUsesForInactiveResources", []interface{}{})
	fake.cleanConfigUsesForInactiveResourcesMutex.Unlock()
	if fake.CleanConfigUsesForInactiveResourcesStub != nil {
		return fake.CleanConfigUsesForInactiveResourcesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanConfigUsesForInactiveResourcesReturns.result1
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourcesCallCount() int {
	fake.cleanConfigUsesForInactiveResourcesMutex.RLock()
	defer fake.cleanConfigUsesForInactiveResourcesMutex.RUnlock()
	return len(fake.cleanConfigUsesForInactiveResourcesArgsForCall)
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourcesReturns(result1 error) {
	fake.CleanConfigUsesForInactiveResourcesStub = nil
	fake.cleanConfigUsesForInactiveResourcesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourcesReturnsOnCall(i int, result1 error) {
	fake.CleanConfigUsesForInactiveResourcesStub = nil
	if fake.cleanConfigUsesForInactiveResourcesReturnsOnCall == nil {
		fake.cleanConfigUsesForInactiveResourcesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanConfigUsesForInactiveResourcesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForPausedPipelinesResources() error {
	fake.cleanConfigUsesForPausedPipelinesResourcesMutex.Lock()
	ret, specificReturn := fake.cleanConfigUsesForPausedPipelinesResourcesReturnsOnCall[len(fake.cleanConfigUsesForPausedPipelinesResourcesArgsForCall)]
	fake.cleanConfigUsesForPausedPipelinesResourcesArgsForCall = append(fake.cleanConfigUsesForPausedPipelinesResourcesArgsForCall, struct{}{})
	fake.recordInvocation("CleanConfigUsesForPausedPipelinesResources", []interface{}{})
	fake.cleanConfigUsesForPausedPipelinesResourcesMutex.Unlock()
	if fake.CleanConfigUsesForPausedPipelinesResourcesStub != nil {
		return fake.CleanConfigUsesForPausedPipelinesResourcesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanConfigUsesForPausedPipelinesResourcesReturns.result1
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForPausedPipelinesResourcesCallCount() int {
	fake.cleanConfigUsesForPausedPipelinesResourcesMutex.RLock()
	defer fake.cleanConfigUsesForPausedPipelinesResourcesMutex.RUnlock()
	return len(fake.cleanConfigUsesForPausedPipelinesResourcesArgsForCall)
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForPausedPipelinesResourcesReturns(result1 error) {
	fake.CleanConfigUsesForPausedPipelinesResourcesStub = nil
	fake.cleanConfigUsesForPausedPipelinesResourcesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForPausedPipelinesResourcesReturnsOnCall(i int, result1 error) {
	fake.CleanConfigUsesForPausedPipelinesResourcesStub = nil
	if fake.cleanConfigUsesForPausedPipelinesResourcesReturnsOnCall == nil {
		fake.cleanConfigUsesForPausedPipelinesResourcesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanConfigUsesForPausedPipelinesResourcesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForOutdatedResourceConfigs() error {
	fake.cleanConfigUsesForOutdatedResourceConfigsMutex.Lock()
	ret, specificReturn := fake.cleanConfigUsesForOutdatedResourceConfigsReturnsOnCall[len(fake.cleanConfigUsesForOutdatedResourceConfigsArgsForCall)]
	fake.cleanConfigUsesForOutdatedResourceConfigsArgsForCall = append(fake.cleanConfigUsesForOutdatedResourceConfigsArgsForCall, struct{}{})
	fake.recordInvocation("CleanConfigUsesForOutdatedResourceConfigs", []interface{}{})
	fake.cleanConfigUsesForOutdatedResourceConfigsMutex.Unlock()
	if fake.CleanConfigUsesForOutdatedResourceConfigsStub != nil {
		return fake.CleanConfigUsesForOutdatedResourceConfigsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanConfigUsesForOutdatedResourceConfigsReturns.result1
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForOutdatedResourceConfigsCallCount() int {
	fake.cleanConfigUsesForOutdatedResourceConfigsMutex.RLock()
	defer fake.cleanConfigUsesForOutdatedResourceConfigsMutex.RUnlock()
	return len(fake.cleanConfigUsesForOutdatedResourceConfigsArgsForCall)
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForOutdatedResourceConfigsReturns(result1 error) {
	fake.CleanConfigUsesForOutdatedResourceConfigsStub = nil
	fake.cleanConfigUsesForOutdatedResourceConfigsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForOutdatedResourceConfigsReturnsOnCall(i int, result1 error) {
	fake.CleanConfigUsesForOutdatedResourceConfigsStub = nil
	if fake.cleanConfigUsesForOutdatedResourceConfigsReturnsOnCall == nil {
		fake.cleanConfigUsesForOutdatedResourceConfigsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanConfigUsesForOutdatedResourceConfigsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanUselessConfigs() error {
	fake.cleanUselessConfigsMutex.Lock()
	ret, specificReturn := fake.cleanUselessConfigsReturnsOnCall[len(fake.cleanUselessConfigsArgsForCall)]
	fake.cleanUselessConfigsArgsForCall = append(fake.cleanUselessConfigsArgsForCall, struct{}{})
	fake.recordInvocation("CleanUselessConfigs", []interface{}{})
	fake.cleanUselessConfigsMutex.Unlock()
	if fake.CleanUselessConfigsStub != nil {
		return fake.CleanUselessConfigsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanUselessConfigsReturns.result1
}

func (fake *FakeResourceConfigFactory) CleanUselessConfigsCallCount() int {
	fake.cleanUselessConfigsMutex.RLock()
	defer fake.cleanUselessConfigsMutex.RUnlock()
	return len(fake.cleanUselessConfigsArgsForCall)
}

func (fake *FakeResourceConfigFactory) CleanUselessConfigsReturns(result1 error) {
	fake.CleanUselessConfigsStub = nil
	fake.cleanUselessConfigsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanUselessConfigsReturnsOnCall(i int, result1 error) {
	fake.CleanUselessConfigsStub = nil
	if fake.cleanUselessConfigsReturnsOnCall == nil {
		fake.cleanUselessConfigsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanUselessConfigsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) AcquireResourceCheckingLock(logger lager.Logger, resourceUser db.ResourceUser, resourceType string, resourceSource atc.Source, resourceTypes creds.VersionedResourceTypes) (lock.Lock, bool, error) {
	fake.acquireResourceCheckingLockMutex.Lock()
	ret, specificReturn := fake.acquireResourceCheckingLockReturnsOnCall[len(fake.acquireResourceCheckingLockArgsForCall)]
	fake.acquireResourceCheckingLockArgsForCall = append(fake.acquireResourceCheckingLockArgsForCall, struct {
		logger         lager.Logger
		resourceUser   db.ResourceUser
		resourceType   string
		resourceSource atc.Source
		resourceTypes  creds.VersionedResourceTypes
	}{logger, resourceUser, resourceType, resourceSource, resourceTypes})
	fake.recordInvocation("AcquireResourceCheckingLock", []interface{}{logger, resourceUser, resourceType, resourceSource, resourceTypes})
	fake.acquireResourceCheckingLockMutex.Unlock()
	if fake.AcquireResourceCheckingLockStub != nil {
		return fake.AcquireResourceCheckingLockStub(logger, resourceUser, resourceType, resourceSource, resourceTypes)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.acquireResourceCheckingLockReturns.result1, fake.acquireResourceCheckingLockReturns.result2, fake.acquireResourceCheckingLockReturns.result3
}

func (fake *FakeResourceConfigFactory) AcquireResourceCheckingLockCallCount() int {
	fake.acquireResourceCheckingLockMutex.RLock()
	defer fake.acquireResourceCheckingLockMutex.RUnlock()
	return len(fake.acquireResourceCheckingLockArgsForCall)
}

func (fake *FakeResourceConfigFactory) AcquireResourceCheckingLockArgsForCall(i int) (lager.Logger, db.ResourceUser, string, atc.Source, creds.VersionedResourceTypes) {
	fake.acquireResourceCheckingLockMutex.RLock()
	defer fake.acquireResourceCheckingLockMutex.RUnlock()
	return fake.acquireResourceCheckingLockArgsForCall[i].logger, fake.acquireResourceCheckingLockArgsForCall[i].resourceUser, fake.acquireResourceCheckingLockArgsForCall[i].resourceType, fake.acquireResourceCheckingLockArgsForCall[i].resourceSource, fake.acquireResourceCheckingLockArgsForCall[i].resourceTypes
}

func (fake *FakeResourceConfigFactory) AcquireResourceCheckingLockReturns(result1 lock.Lock, result2 bool, result3 error) {
	fake.AcquireResourceCheckingLockStub = nil
	fake.acquireResourceCheckingLockReturns = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceConfigFactory) AcquireResourceCheckingLockReturnsOnCall(i int, result1 lock.Lock, result2 bool, result3 error) {
	fake.AcquireResourceCheckingLockStub = nil
	if fake.acquireResourceCheckingLockReturnsOnCall == nil {
		fake.acquireResourceCheckingLockReturnsOnCall = make(map[int]struct {
			result1 lock.Lock
			result2 bool
			result3 error
		})
	}
	fake.acquireResourceCheckingLockReturnsOnCall[i] = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceConfigFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findOrCreateResourceConfigMutex.RLock()
	defer fake.findOrCreateResourceConfigMutex.RUnlock()
	fake.findResourceConfigMutex.RLock()
	defer fake.findResourceConfigMutex.RUnlock()
	fake.cleanConfigUsesForFinishedBuildsMutex.RLock()
	defer fake.cleanConfigUsesForFinishedBuildsMutex.RUnlock()
	fake.cleanConfigUsesForInactiveResourceTypesMutex.RLock()
	defer fake.cleanConfigUsesForInactiveResourceTypesMutex.RUnlock()
	fake.cleanConfigUsesForInactiveResourcesMutex.RLock()
	defer fake.cleanConfigUsesForInactiveResourcesMutex.RUnlock()
	fake.cleanConfigUsesForPausedPipelinesResourcesMutex.RLock()
	defer fake.cleanConfigUsesForPausedPipelinesResourcesMutex.RUnlock()
	fake.cleanConfigUsesForOutdatedResourceConfigsMutex.RLock()
	defer fake.cleanConfigUsesForOutdatedResourceConfigsMutex.RUnlock()
	fake.cleanUselessConfigsMutex.RLock()
	defer fake.cleanUselessConfigsMutex.RUnlock()
	fake.acquireResourceCheckingLockMutex.RLock()
	defer fake.acquireResourceCheckingLockMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceConfigFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.ResourceConfigFactory = new(FakeResourceConfigFactory)
