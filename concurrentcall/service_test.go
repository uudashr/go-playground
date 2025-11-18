package concurrentcall_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/uudashr/go-playground/concurrentcall"
)

// try to run test multiple times using `-race`, sometimes you get it passed and sometimes failed if not using race condition free
const raceConditionFree = true

func TestService_Call(t *testing.T) {
	tests := []struct {
		name      string
		callers   func() []*CallerMock
		expectRes []string
		expectErr error
	}{
		{
			name: "foo first",
			callers: func() []*CallerMock {
				called := make(chan struct{})

				caller1 := new(CallerMock)
				caller2 := new(CallerMock)

				if raceConditionFree {
					caller1.On("Call").Return(func() (string, error) {
						close(called)
						return "foo", nil
					})

					caller2.On("Call").Return(func() (string, error) {
						<-called
						return "bar", nil
					})
				} else {
					caller1.On("Call").Run(func(args mock.Arguments) {
						close(called)
					}).Return("foo", nil)

					caller2.On("Call").Return(func(args mock.Arguments) {
						<-called
					}).Return("bar", nil)
				}

				return []*CallerMock{caller1, caller2}
			},
			expectRes: []string{"foo", "bar"},
		},
		{
			name: "bar first",
			callers: func() []*CallerMock {
				called := make(chan struct{})

				caller1 := new(CallerMock)
				caller2 := new(CallerMock)

				if raceConditionFree {
					caller1.On("Call").Return(func() (string, error) {
						<-called
						return "foo", nil
					})

					caller2.On("Call").Return(func() (string, error) {
						close(called)
						return "bar", nil
					})
				} else {
					caller1.On("Call").Run(func(args mock.Arguments) {
						<-called
					}).Return("foo", nil)

					caller2.On("Call").Run(func(args mock.Arguments) {
						close(called)
					}).Return("bar", nil)
				}

				return []*CallerMock{caller1, caller2}
			},
			expectRes: []string{"bar", "foo"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			callers := tt.callers()
			callersArg := make([]concurrentcall.Caller, len(callers))
			for i, c := range callers {
				callersArg[i] = c
			}

			svc := concurrentcall.NewService(callersArg)

			// Act
			res, err := svc.Do()

			// Assert
			require.Equal(t, tt.expectErr, err)
			assert.Equal(t, tt.expectRes, res)
		})
	}
}

type CallerMock struct {
	mock.Mock
}

func (m *CallerMock) Call() (string, error) {
	args := m.Called()
	var (
		ret1 string
		ret2 error
	)

	if returnFn, ok := args.Get(0).(func() (string, error)); ok {
		return returnFn()
	}

	if returnFn, ok := args.Get(0).(func() string); ok {
		ret1 = returnFn()
	} else {
		ret1 = args.String(0)
	}

	if returnFn, ok := args.Get(1).(func() error); ok {
		ret2 = returnFn()
	} else {
		ret2 = args.Error(1)
	}

	return ret1, ret2
}
