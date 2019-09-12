package types

import (
	"testing"
)

// ========= STARTS WITH ===========
func TestString_StartsWithBlank(t *testing.T) {
	//t.Error("This is a error message")
	resultVal := NewString("Vaibhav").StartsWith("")
	if true != resultVal {
		t.Errorf("Expected true. Got: %v", resultVal)
	}
}

func TestString_StartsWithBigger(t *testing.T) {
	//t.Error("This is a error message")
	resultVal := NewString("Vaibhav").StartsWith("Vaibhav1")
	if false != resultVal {
		t.Errorf("Expected true. Got: %v", resultVal)
	}
}

func TestString_StartsWithBiggerInsensitive(t *testing.T) {
	//t.Error("This is a error message")
	resultVal := NewString("Vaibhav").StartsWith("Vaibhav1", true)
	if false != resultVal {
		t.Errorf("Expected true. Got: %v", resultVal)
	}
}

func TestString_StartsWithFew(t *testing.T) {
	//t.Error("This is a error message")
	resultVal := NewString("Vaibhav").StartsWith("Vai")
	if true != resultVal {
		t.Errorf("Expected true. Got: %v", resultVal)
	}
}

func TestString_StartsWithFewInsensitive(t *testing.T) {
	//t.Error("This is a error message")
	resultVal := NewString("Vaibhav").StartsWith("Vai", true)
	if true != resultVal {
		t.Errorf("Expected true. Got: %v", resultVal)
	}
}

// ========= ENDS WITH ===========
func TestString_EndsWithBlank(t *testing.T) {
	//t.Error("This is a error message")
	resultVal := NewString("Vaibhav").EndsWith("")
	if true != resultVal {
		t.Errorf("Expected true. Got: %v", resultVal)
	}
}

func TestString_EndsWithBigger(t *testing.T) {
	//t.Error("This is a error message")
	resultVal := NewString("Vaibhav").EndsWith("Vaibhav1")
	if false != resultVal {
		t.Errorf("Expected true. Got: %v", resultVal)
	}
}

func TestString_EndsWithBiggerInsensitive(t *testing.T) {
	//t.Error("This is a error message")
	resultVal := NewString("Vaibhav").EndsWith("Vaibhav1", true)
	if false != resultVal {
		t.Errorf("Expected true. Got: %v", resultVal)
	}
}

func TestString_EndsWithFew(t *testing.T) {
	//t.Error("This is a error message")
	resultVal := NewString("Vaibhav").EndsWith("hav")
	if true != resultVal {
		t.Errorf("Expected true. Got: %v", resultVal)
	}
}

func TestString_EndsWithFewInsensitive(t *testing.T) {
	//t.Error("This is a error message")
	resultVal := NewString("Vaibhav").EndsWith("HAV", true)
	if true != resultVal {
		t.Errorf("Expected true. Got: %v", resultVal)
	}
}

