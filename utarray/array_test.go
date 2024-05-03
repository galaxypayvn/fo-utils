package utarray

import (
	"reflect"
	"testing"
)

func TestArrayToMap(t *testing.T) {
	type testCase[T comparable] struct {
		name     string
		input    []T
		expected map[T]struct{}
	}
	tests := []testCase[int]{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: map[int]struct{}{},
		},
		{
			name:     "Slice with unique elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}},
		},
		{
			name:     "Slice with duplicates",
			input:    []int{1, 2, 3, 3, 4, 5, 5},
			expected: map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayToMap(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ArrayToMap() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type testCase[T comparable] struct {
		name     string
		input    []T
		expected []T
	}
	tests := []testCase[int]{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Slice with single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Slice with unique elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Slice with duplicates",
			input:    []int{1, 2, 3, 3, 4, 5, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Unique() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type testCase[T comparable] struct {
		name     string
		slice    []T
		value    T
		expected bool
	}
	tests := []testCase[int]{
		{
			name:     "Empty slice",
			slice:    []int{},
			value:    42,
			expected: false,
		},
		{
			name:     "Value not present",
			slice:    []int{1, 2, 3, 4, 5},
			value:    42,
			expected: false,
		},
		{
			name:     "Value present",
			slice:    []int{1, 2, 3, 4, 5},
			value:    3,
			expected: true,
		},
		{
			name:     "Duplicate values",
			slice:    []int{1, 2, 3, 3, 4, 5},
			value:    3,
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.slice, tt.value); got != tt.expected {
				t.Errorf("Contains() = %v, want %v", got, tt.expected)
			}
		})
	}
}
