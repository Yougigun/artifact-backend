// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	pb "github.com/instill-ai/protogen-go/artifact/artifact/v1alpha"
)

// RepositoryMock implements service.Repository
type RepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGetRepositoryTag          func(ctx context.Context, name string) (rp1 *pb.RepositoryTag, err error)
	inspectFuncGetRepositoryTag   func(ctx context.Context, name string)
	afterGetRepositoryTagCounter  uint64
	beforeGetRepositoryTagCounter uint64
	GetRepositoryTagMock          mRepositoryMockGetRepositoryTag
}

// NewRepositoryMock returns a mock for service.Repository
func NewRepositoryMock(t minimock.Tester) *RepositoryMock {
	m := &RepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetRepositoryTagMock = mRepositoryMockGetRepositoryTag{mock: m}
	m.GetRepositoryTagMock.callArgs = []*RepositoryMockGetRepositoryTagParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mRepositoryMockGetRepositoryTag struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockGetRepositoryTagExpectation
	expectations       []*RepositoryMockGetRepositoryTagExpectation

	callArgs []*RepositoryMockGetRepositoryTagParams
	mutex    sync.RWMutex
}

// RepositoryMockGetRepositoryTagExpectation specifies expectation struct of the Repository.GetRepositoryTag
type RepositoryMockGetRepositoryTagExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockGetRepositoryTagParams
	results *RepositoryMockGetRepositoryTagResults
	Counter uint64
}

// RepositoryMockGetRepositoryTagParams contains parameters of the Repository.GetRepositoryTag
type RepositoryMockGetRepositoryTagParams struct {
	ctx  context.Context
	name string
}

// RepositoryMockGetRepositoryTagResults contains results of the Repository.GetRepositoryTag
type RepositoryMockGetRepositoryTagResults struct {
	rp1 *pb.RepositoryTag
	err error
}

// Expect sets up expected params for Repository.GetRepositoryTag
func (mmGetRepositoryTag *mRepositoryMockGetRepositoryTag) Expect(ctx context.Context, name string) *mRepositoryMockGetRepositoryTag {
	if mmGetRepositoryTag.mock.funcGetRepositoryTag != nil {
		mmGetRepositoryTag.mock.t.Fatalf("RepositoryMock.GetRepositoryTag mock is already set by Set")
	}

	if mmGetRepositoryTag.defaultExpectation == nil {
		mmGetRepositoryTag.defaultExpectation = &RepositoryMockGetRepositoryTagExpectation{}
	}

	mmGetRepositoryTag.defaultExpectation.params = &RepositoryMockGetRepositoryTagParams{ctx, name}
	for _, e := range mmGetRepositoryTag.expectations {
		if minimock.Equal(e.params, mmGetRepositoryTag.defaultExpectation.params) {
			mmGetRepositoryTag.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetRepositoryTag.defaultExpectation.params)
		}
	}

	return mmGetRepositoryTag
}

// Inspect accepts an inspector function that has same arguments as the Repository.GetRepositoryTag
func (mmGetRepositoryTag *mRepositoryMockGetRepositoryTag) Inspect(f func(ctx context.Context, name string)) *mRepositoryMockGetRepositoryTag {
	if mmGetRepositoryTag.mock.inspectFuncGetRepositoryTag != nil {
		mmGetRepositoryTag.mock.t.Fatalf("Inspect function is already set for RepositoryMock.GetRepositoryTag")
	}

	mmGetRepositoryTag.mock.inspectFuncGetRepositoryTag = f

	return mmGetRepositoryTag
}

// Return sets up results that will be returned by Repository.GetRepositoryTag
func (mmGetRepositoryTag *mRepositoryMockGetRepositoryTag) Return(rp1 *pb.RepositoryTag, err error) *RepositoryMock {
	if mmGetRepositoryTag.mock.funcGetRepositoryTag != nil {
		mmGetRepositoryTag.mock.t.Fatalf("RepositoryMock.GetRepositoryTag mock is already set by Set")
	}

	if mmGetRepositoryTag.defaultExpectation == nil {
		mmGetRepositoryTag.defaultExpectation = &RepositoryMockGetRepositoryTagExpectation{mock: mmGetRepositoryTag.mock}
	}
	mmGetRepositoryTag.defaultExpectation.results = &RepositoryMockGetRepositoryTagResults{rp1, err}
	return mmGetRepositoryTag.mock
}

// Set uses given function f to mock the Repository.GetRepositoryTag method
func (mmGetRepositoryTag *mRepositoryMockGetRepositoryTag) Set(f func(ctx context.Context, name string) (rp1 *pb.RepositoryTag, err error)) *RepositoryMock {
	if mmGetRepositoryTag.defaultExpectation != nil {
		mmGetRepositoryTag.mock.t.Fatalf("Default expectation is already set for the Repository.GetRepositoryTag method")
	}

	if len(mmGetRepositoryTag.expectations) > 0 {
		mmGetRepositoryTag.mock.t.Fatalf("Some expectations are already set for the Repository.GetRepositoryTag method")
	}

	mmGetRepositoryTag.mock.funcGetRepositoryTag = f
	return mmGetRepositoryTag.mock
}

// When sets expectation for the Repository.GetRepositoryTag which will trigger the result defined by the following
// Then helper
func (mmGetRepositoryTag *mRepositoryMockGetRepositoryTag) When(ctx context.Context, name string) *RepositoryMockGetRepositoryTagExpectation {
	if mmGetRepositoryTag.mock.funcGetRepositoryTag != nil {
		mmGetRepositoryTag.mock.t.Fatalf("RepositoryMock.GetRepositoryTag mock is already set by Set")
	}

	expectation := &RepositoryMockGetRepositoryTagExpectation{
		mock:   mmGetRepositoryTag.mock,
		params: &RepositoryMockGetRepositoryTagParams{ctx, name},
	}
	mmGetRepositoryTag.expectations = append(mmGetRepositoryTag.expectations, expectation)
	return expectation
}

// Then sets up Repository.GetRepositoryTag return parameters for the expectation previously defined by the When method
func (e *RepositoryMockGetRepositoryTagExpectation) Then(rp1 *pb.RepositoryTag, err error) *RepositoryMock {
	e.results = &RepositoryMockGetRepositoryTagResults{rp1, err}
	return e.mock
}

// GetRepositoryTag implements service.Repository
func (mmGetRepositoryTag *RepositoryMock) GetRepositoryTag(ctx context.Context, name string) (rp1 *pb.RepositoryTag, err error) {
	mm_atomic.AddUint64(&mmGetRepositoryTag.beforeGetRepositoryTagCounter, 1)
	defer mm_atomic.AddUint64(&mmGetRepositoryTag.afterGetRepositoryTagCounter, 1)

	if mmGetRepositoryTag.inspectFuncGetRepositoryTag != nil {
		mmGetRepositoryTag.inspectFuncGetRepositoryTag(ctx, name)
	}

	mm_params := RepositoryMockGetRepositoryTagParams{ctx, name}

	// Record call args
	mmGetRepositoryTag.GetRepositoryTagMock.mutex.Lock()
	mmGetRepositoryTag.GetRepositoryTagMock.callArgs = append(mmGetRepositoryTag.GetRepositoryTagMock.callArgs, &mm_params)
	mmGetRepositoryTag.GetRepositoryTagMock.mutex.Unlock()

	for _, e := range mmGetRepositoryTag.GetRepositoryTagMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.rp1, e.results.err
		}
	}

	if mmGetRepositoryTag.GetRepositoryTagMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetRepositoryTag.GetRepositoryTagMock.defaultExpectation.Counter, 1)
		mm_want := mmGetRepositoryTag.GetRepositoryTagMock.defaultExpectation.params
		mm_got := RepositoryMockGetRepositoryTagParams{ctx, name}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetRepositoryTag.t.Errorf("RepositoryMock.GetRepositoryTag got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetRepositoryTag.GetRepositoryTagMock.defaultExpectation.results
		if mm_results == nil {
			mmGetRepositoryTag.t.Fatal("No results are set for the RepositoryMock.GetRepositoryTag")
		}
		return (*mm_results).rp1, (*mm_results).err
	}
	if mmGetRepositoryTag.funcGetRepositoryTag != nil {
		return mmGetRepositoryTag.funcGetRepositoryTag(ctx, name)
	}
	mmGetRepositoryTag.t.Fatalf("Unexpected call to RepositoryMock.GetRepositoryTag. %v %v", ctx, name)
	return
}

// GetRepositoryTagAfterCounter returns a count of finished RepositoryMock.GetRepositoryTag invocations
func (mmGetRepositoryTag *RepositoryMock) GetRepositoryTagAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetRepositoryTag.afterGetRepositoryTagCounter)
}

// GetRepositoryTagBeforeCounter returns a count of RepositoryMock.GetRepositoryTag invocations
func (mmGetRepositoryTag *RepositoryMock) GetRepositoryTagBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetRepositoryTag.beforeGetRepositoryTagCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.GetRepositoryTag.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetRepositoryTag *mRepositoryMockGetRepositoryTag) Calls() []*RepositoryMockGetRepositoryTagParams {
	mmGetRepositoryTag.mutex.RLock()

	argCopy := make([]*RepositoryMockGetRepositoryTagParams, len(mmGetRepositoryTag.callArgs))
	copy(argCopy, mmGetRepositoryTag.callArgs)

	mmGetRepositoryTag.mutex.RUnlock()

	return argCopy
}

// MinimockGetRepositoryTagDone returns true if the count of the GetRepositoryTag invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockGetRepositoryTagDone() bool {
	for _, e := range m.GetRepositoryTagMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetRepositoryTagMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetRepositoryTagCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetRepositoryTag != nil && mm_atomic.LoadUint64(&m.afterGetRepositoryTagCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetRepositoryTagInspect logs each unmet expectation
func (m *RepositoryMock) MinimockGetRepositoryTagInspect() {
	for _, e := range m.GetRepositoryTagMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.GetRepositoryTag with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetRepositoryTagMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetRepositoryTagCounter) < 1 {
		if m.GetRepositoryTagMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.GetRepositoryTag")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.GetRepositoryTag with params: %#v", *m.GetRepositoryTagMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetRepositoryTag != nil && mm_atomic.LoadUint64(&m.afterGetRepositoryTagCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.GetRepositoryTag")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetRepositoryTagInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetRepositoryTagDone()
}
