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

// RotateInfo contains information for naming rotated files.
type RotateInfo struct {
	BaseName  string
	Extension string
	Index     int
	Time      time.Time
}

// NamingFunc defines the function signature for naming rotated files.
type NamingFunc func(info RotateInfo) string

// NameWithIndex produces names like: app_1.log, app_2.log
func NameWithIndex(info RotateInfo) string {
	return fmt.Sprintf("%s_%d%s", info.BaseName, info.Index, info.Extension)
}

// NameWithDateAndIndex produces names like: app_20251021_1.log
func NameWithDateAndIndex(info RotateInfo) string {
	return fmt.Sprintf("%s_%s_%d%s", info.BaseName, info.Time.Format("20060102"), info.Index, info.Extension)
}

// NameWithTimestamp produces names like: app_20251021_153015.log
func NameWithTimestamp(info RotateInfo) string {
	return fmt.Sprintf("%s_%s%s", info.BaseName, info.Time.Format("20060102_150405"), info.Extension)
}

// NameWithTimestampAndIndex produces names like: app_20251021_153015_1.log
func NameWithTimestampAndIndex(info RotateInfo) string {
	return fmt.Sprintf("%s_%s_%d%s", info.BaseName, info.Time.Format("20060102_150405"), info.Index, info.Extension)
}
