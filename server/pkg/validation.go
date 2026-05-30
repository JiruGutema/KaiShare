package pkg

func ValidatePasteLength(content string) bool {
	return len(content) <= 10000
}
