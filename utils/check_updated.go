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
