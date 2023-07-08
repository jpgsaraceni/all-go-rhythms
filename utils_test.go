package main

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, name string, expect, got any) {
	t.Helper()

	if !reflect.DeepEqual(expect, got) {
		t.Logf("Test %s expected %v but got %v", name, expect, got)
		t.Fail()
	}
}
