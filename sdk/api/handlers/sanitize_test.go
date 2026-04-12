package handlers

import (
	"testing"
)

func TestDeepCleanUndefined(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "清理顶层 [undefined] 字段",
			input:    `{"model":"claude-opus-4-6","temperature":"[undefined]","top_k":"[undefined]","messages":[{"role":"user","content":"hello"}]}`,
			expected: `{"messages":[{"role":"user","content":"hello"}],"model":"claude-opus-4-6"}`,
		},
		{
			name:     "清理嵌套 [undefined] 字段",
			input:    `{"messages":[{"role":"user","content":[{"type":"text","text":"hello","cache_control":"[undefined]"}]}]}`,
			expected: `{"messages":[{"role":"user","content":[{"text":"hello","type":"text"}]}]}`,
		},
		{
			name:     "保留正常字段",
			input:    `{"model":"test","temperature":0.7,"messages":[]}`,
			expected: `{"messages":[],"model":"test","temperature":0.7}`,
		},
		{
			name:     "无效 JSON 原样返回",
			input:    `not json`,
			expected: `not json`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := string(DeepCleanUndefined([]byte(tt.input)))
			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}
