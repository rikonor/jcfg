package jcfg

import (
	"fmt"
	"testing"
)

type Settings struct {
	TestField string `json:"testField"`
}

func TestParse(t *testing.T) {
	// Should not panic
	cfg := Parse("./aux/goodConfig.json", &Settings{}).(*Settings)

	if cfg.TestField != "testValue" {
		t.Error("Failed to parse config")
	}
}

func ExampleParse() {
	cfg := Parse("./aux/goodConfig.json", &Settings{}).(*Settings)

	fmt.Println(cfg.TestField)
	// Output: testValue
}
