/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

type Field struct {
	Key   string
	Value any
}

func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

func Int64(key string, value int64) Field {
	return Field{Key: key, Value: value}
}

func Float64(key string, value float64) Field {
	return Field{Key: key, Value: value}
}

func Bool(key string, value bool) Field {
	return Field{Key: key, Value: value}
}

func Any(key string, value any) Field {
	return Field{Key: key, Value: value}
}

func Err(err error) Field {
	return Field{Key: "error", Value: err}
}
