package main

import (
	"testing"

	"github.com/Rajeevnita1993/json-parser/validator"
)

func TestParseJSON_ValidJSON(t *testing.T) {
	data := []byte(`{"name": "John", "age": 30}`)
	err := validator.ParseJSON(data)
	if err != nil {
		t.Errorf("Expected no error for valid JSON, got: %v", err)
	}

}

func TestParseJSON_InvalidJSON(t *testing.T) {
	data := []byte(`{"name": "John", "age":}`)
	err := validator.ParseJSON(data)
	if err == nil {
		t.Error("Expected error for invalid JSON, got none")
	}
}

func TestParseJSON_EmptyJSON(t *testing.T) {
	data := []byte(`{}`)
	err := validator.ParseJSON(data)
	if err != nil {
		t.Errorf("Expected no error for empty JSON, got: %v", err)
	}
}
