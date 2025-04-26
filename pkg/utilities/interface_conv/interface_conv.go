package interface_conv

import (
	"errors"
	"strconv"
	"strings"
)

func ToUint(iParem interface{}) (uint64, error) {
	switch v := iParem.(type) {
	case int:
		return uint64(v), nil
	case int32:
		return uint64(v), nil
	case int64:
		return uint64(v), nil
	case float32:
		return uint64(v), nil
	case float64:
		return uint64(v), nil
	case string:
		intVal, err := strconv.ParseUint(strings.TrimSpace(v), 10, 32)
		if err != nil {
			return 0, err
		}
		return intVal, nil
	default:
		return 0, errors.New("can't convert")
	}
}

func ToFloat(fParem interface{}) (float64, error) {
	var floatVal float64
	var err error
	switch v := fParem.(type) {
	case int:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return float64(v), nil
	case string:
		floatVal, err = strconv.ParseFloat(strings.TrimSpace(fParem.(string)), 64)
		if err != nil {
			return 0, err
		}
		return floatVal, nil
	default:
		return 0, errors.New("can't convert")
	}
}
