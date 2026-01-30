package logging

import "testing"

func TestInit(t *testing.T) {
	if Logger != nil {
		t.Log("Logger already initialized (unexpected but not fatal)")
	}

	Init("test-service")

	if Logger == nil {
		t.Fatal("expected Logger to be non-nil after Init")
	}

	Logger.Println("test log message")
}
