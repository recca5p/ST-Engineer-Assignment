package utils

import "database/sql"

// Convert sql.NullTime to string, return empty string if NullTime is invalid
func FormatNullTime(t sql.NullTime) string {
	if t.Valid {
		return t.Time.Format("2006-01-02 15:04:05") // Customize the date format as needed
	}
	return ""
}
