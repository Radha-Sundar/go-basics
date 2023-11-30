// greeter_test.go

package greeter_test

import (
	"github.com/Radha-Sundar/godep-app/pkg/greeter"
	"testing"
)

func TestGetGreeting(t *testing.T) {
	expected := "Hello from the greeter package!"
	result := greeter.GetGreeting()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
