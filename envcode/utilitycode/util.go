package utilitycode

import "strings"

//SubsAfer func
func SubsAfer(value string, after string) string {
	var lenght int
	var pos int
	if len(after) == 1 {
		pos = strings.Index(value, after) + 1
		if pos == -1 {
			return ""
		}
		lenght = len(value)

	} else {
		pos = strings.Index(value, after) + len(after)
		if pos == -1 {
			return ""
		}
		lenght = len(value)
	}
	return value[pos:lenght]
}

//SubsBefore func
func SubsBefore(value string, before string) string {
	pos := strings.Index(value, before)
	if pos == -1 {
		return value
	}
	return value[0:pos]
}
