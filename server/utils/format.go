package utils

import (
	"database/sql"
	"time"
)

func FormatNullTime(t sql.NullTime) string {
	if t.Valid {
		return t.Time.Format("2006-01-02 15:04:05")
	}
	return ""
}

func ConvertNullTime(nullTime sql.NullTime) (time.Time, bool) {
	if nullTime.Valid {
		return nullTime.Time, true
	}
	return time.Time{}, false
}
