package isogram

import "strings"

func IsIsogram(word string) bool {
	lowerWord := strings.ToLower(word)

	// هر کاراکتر را با تمام کاراکترهای بعد از آن مقایسه می‌کنیم
	for i := 0; i < len(lowerWord); i++ {
		r := lowerWord[i]

		// فاصله و خط تیره نادیده گرفته می‌شوند
		if r == ' ' || r == '-' {
			continue
		}

		for j := i + 1; j < len(lowerWord); j++ {
			if r == lowerWord[j] {
				return false
			}
		}
	}
	return true
}
