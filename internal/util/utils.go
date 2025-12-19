package util

import "fmt"

func HumanSize(bytes int64) string {
	if bytes <= 0 {
		return "unknown size"
	}

	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div := float64(unit)
	exp := 0
	for n := float64(bytes) / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf(
		"%.2f %ciB",
		float64(bytes)/div,
		"KMGTPE"[exp],
	)
}