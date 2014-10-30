package mock

import (
	"testing"
)

type Mock interface {
	InitMock()
	ExpectCall(string, ...interface{})
	ExpectedCallsCalled(*testing.T)
}

type MockHelper struct {
	calls map[string]call
}

type call struct {
	funcName string
	args     []interface{}
	expected int
	called   int
}

func (n *MockHelper) AddCall(funcName string, args ...interface{}) {
	if val, exists := n.calls[funcName]; exists {
		if val.Equal(call{args: args}) {
			val.called = val.called + 1
			n.calls[funcName] = val
			return
		}
	}
	n.calls[funcName] = call{funcName, args, 0, 1}
}

func (n *MockHelper) InitMock() {
	n.calls = map[string]call{}
}

func (n *MockHelper) ExpectCall(funcName string, args []interface{}) {
	if val, exists := n.calls[funcName]; exists {
		val.expected = val.expected + 1
		n.calls[funcName] = val
	} else {
		n.calls[funcName] = call{funcName, args, 1, 0}
	}
}

func (n *MockHelper) ExpectedCallsCalled(t *testing.T) {
	for _, elem := range n.calls {
		if elem.called != elem.expected {
			t.Errorf("%+v", elem)
		}
	}
}

func (a call) Equal(b call) bool {
	if len(a.args) != len(b.args) {
		return false
	}
	for i, arg := range a.args {
		if arg != b.args[i] {
			return false
		}
	}
	return true
}
