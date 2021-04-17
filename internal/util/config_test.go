package util

import (
	"testing"
)

func TestGetNestedConfigValueAccessor(t *testing.T) {
	expected := "top.gladys.key"
	actual := GetNestedConfigValueAccessor("top", "gladys", "key")

	if actual != expected {
		t.Fatalf("GetNestedConfigValueAccessor did not return the expected value: %s", expected)
	}
}
