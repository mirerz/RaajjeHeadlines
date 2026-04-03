package utils

import (
	"regexp"
	"strings"
)

// NormalizeThaana cleans and standardizes Maldivian Thaana text from scrapers
func NormalizeThaana(input string) string {
	// 1. Remove Zero Width Spaces (common in web copy-pasting)
	zws := regexp.MustCompile(`[\x{200B}-\x{200D}\x{FEFF}]`)
	result := zws.ReplaceAllString(input, "")

	// 2. Normalize multiple spaces
	spaces := regexp.MustCompile(`\s+`)
	result = spaces.ReplaceAllString(result, " ")

	// 3. Normalize common Dhivehi punctuations
	result = strings.ReplaceAll(result, "،", ",")
	result = strings.ReplaceAll(result, "؛", ";")
	result = strings.ReplaceAll(result, "؟", "?")

	// 4. Trim artifacts
	return strings.TrimSpace(result)
}
