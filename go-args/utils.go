package main

import "strconv"

func toBool(v interface{}) bool {
	switch v.(type) {
	case bool:
		return v.(bool)
	case string:
		return v.(string) == "true"
	case int:
		return v.(int) != 0
	case int32:
		return v.(int32) != 0
	case int64:
		return v.(int64) != 0
	case float32:
		return v.(float32) != 0
	default:
		return false
	}
}

func toInt(v interface{}) int {
	switch v.(type) {
	case bool:
		return 0
	case string:
		i, _ := strconv.Atoi(v.(string))
		return i
	case int:
		return v.(int)
	case int32:
		return int(v.(int32))
	case int64:
		return int(v.(int64))
	case float32:
		return int(v.(float32))
	default:
		return 0
	}
}

func toString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	default:
		return "0"
	}
}
