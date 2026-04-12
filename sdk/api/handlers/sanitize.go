// Package handlers 通用请求清理工具。
package handlers

import (
	"encoding/json"
)

// DeepCleanUndefined 递归遍历 JSON 数据，移除值为 "[undefined]" 的字段。
// 主要解决 Cherry Studio 等客户端将未定义参数序列化为字符串 "[undefined]" 的问题，
// 导致上游 API（如 Google Cloud Code）返回 400 INVALID_ARGUMENT 错误。
func DeepCleanUndefined(rawJSON []byte) []byte {
	var data interface{}
	if err := json.Unmarshal(rawJSON, &data); err != nil {
		return rawJSON
	}
	cleanUndefinedRecursive(data, 0)
	cleaned, err := json.Marshal(data)
	if err != nil {
		return rawJSON
	}
	return cleaned
}

func cleanUndefinedRecursive(v interface{}, depth int) {
	if depth > 10 {
		return
	}
	switch val := v.(type) {
	case map[string]interface{}:
		// 移除值为 "[undefined]" 的键
		for k, v := range val {
			if s, ok := v.(string); ok && s == "[undefined]" {
				delete(val, k)
			} else {
				cleanUndefinedRecursive(v, depth+1)
			}
		}
	case []interface{}:
		for _, item := range val {
			cleanUndefinedRecursive(item, depth+1)
		}
	}
}
