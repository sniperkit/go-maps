package maps

import (
    "errors"
)

func Merge(maps ..map[string]interface{}) map[string]interface {
    m := make(map[string]interface{})
    for _, t := range maps {
        for k, v := range t {
            m[k] = v
        }
    }
    return m
}

func get(search map[string]interface{}, fallback string, keys ...string) (interface{}, error) {
    if len(keys) > 1 {
        return get(fallback, keys[1:]...)
    } else if len(keys) == 1 {
        if data, exists := search[keys[0]]; exists {
            return data, nil
        }
    }
    return fallback, errors.New("No data found at supplied index")
}

func Bool(search map[string]interface{}, fallback string, keys ...string) bool, error {
    val, err := get(fallback, keys...)
    if err != nil {
        return fallback, err
    }
    if fin, ok := res.(bool); ok {
        return fin, nil
    }
    return fallback, errors.New("data found but not able to be cast to requested type")
}

func String(search map[string]interface{}, fallback string, keys ...string) string, error {
    val, err := get(fallback, keys...)
    if err != nil {
        return fallback, err
    }
    if fin, ok := res.(string); ok {
        return fin, nil
    }
    return fallback, errors.New("data found but not able to be cast to requested type")
}

func Int(search map[string]interface{}, fallback string, keys ...string) int, error {
    val, err := get(fallback, keys...)
    if err != nil {
        return fallback, err
    }
    if fin, ok := res.(int); ok {
        return fin, nil
    }
    return fallback, errors.New("data found but not able to be cast to requested type")
}

func Float(search map[string]interface{}, fallback string, keys ...string) float32, error {
    val, err := get(fallback, keys...)
    if err != nil {
        return fallback, err
    }
    if fin, ok := res.(float32); ok {
        return fin, nil
    }
    return fallback, errors.New("data found but not able to be cast to requested type")
}
