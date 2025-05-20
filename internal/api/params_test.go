package api_test

import (
	"net/url"
	"testing"

	"github.com/ryohidaka/go-zaim/internal/api"
	"github.com/stretchr/testify/assert"
)

// テスト用構造体
type testParams struct {
	ID    int    `url:"id"`
	Name  string `url:"name,omitempty"`
	Empty string `url:"empty,omitempty"`
}

func TestBuildQueryParams(t *testing.T) {
	tests := []struct {
		name     string
		input    []testParams
		expected url.Values
	}{
		{
			name:  "All fields present",
			input: []testParams{{ID: 123, Name: "Alice"}},
			expected: url.Values{
				"id":   []string{"123"},
				"name": []string{"Alice"},
			},
		},
		{
			name:  "Omit empty optional fields",
			input: []testParams{{ID: 456}},
			expected: url.Values{
				"id": []string{"456"},
			},
		},
		{
			name:     "No input provided",
			input:    []testParams{},
			expected: url.Values{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// スライスを展開して渡す（重要）
			values, err := api.BuildQueryParams(tt.input...)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, values)
		})
	}
}
