package utils

func CheckNullString(new *string, old string) string {
	if new == nil {
		return old
	} else {
		return *new
	}
}

func CheckNullInt(new *int, old int) int {
	if new == nil {
		return old
	} else {
		return *new
	}
}

func CheckUpdatedString(new_value string, old_value string) string {
	if new_value != old_value && new_value != "" {
		return new_value
	} else {
		return old_value
	}
}

func CheckUpdatedInt(new_value int, old_value int) int {
	if new_value != old_value && new_value != 0 {
		return new_value
	} else {
		return old_value
	}
}
