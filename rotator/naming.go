/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package rotator

import (
	"fmt"
	"time"
)

type RotateInfo struct {
	BaseName  string
	Extension string
	Index     int
	Time      time.Time
}

type NamingFunc func(info RotateInfo) string

func NameWithIndex(info RotateInfo) string {
	return fmt.Sprintf("%s_%d%s", info.BaseName, info.Index, info.Extension)
}

func NameWithDateAndIndex(info RotateInfo) string {
	return fmt.Sprintf("%s_%s_%d%s", info.BaseName, info.Time.Format("20060102"), info.Index, info.Extension)
}

func NameWithTimestamp(info RotateInfo) string {
	return fmt.Sprintf("%s_%s%s", info.BaseName, info.Time.Format("20060102_150405"), info.Extension)
}

func NameWithTimestampAndIndex(info RotateInfo) string {
	return fmt.Sprintf("%s_%s_%d%s", info.BaseName, info.Time.Format("20060102_150405"), info.Index, info.Extension)
}
