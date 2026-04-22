package utils

import (
	"regexp"
	"strings"
)

// SanitizeHTML strips all HTML tags and noise, leaving only clean text for the AI.
func SanitizeHTML(input string) string {
	// 1. Remove scripts and styles entirely
	reScript := regexp.MustCompile(`(?s)<script.*?>.*?</script>`)
	input = reScript.ReplaceAllString(input, "")
	
	reStyle := regexp.MustCompile(`(?s)<style.*?>.*?</style>`)
	input = reStyle.ReplaceAllString(input, "")

	// 2. Remove all remaining HTML tags
	reTags := regexp.MustCompile(`<.*?>`)
	input = reTags.ReplaceAllString(input, " ")

	// 3. Normalize whitespace and special characters
	input = strings.ReplaceAll(input, "&nbsp;", " ")
	input = strings.ReplaceAll(input, "&quot;", "\"")
	input = strings.ReplaceAll(input, "&amp;", "&")
	
	fields := strings.Fields(input)
	return strings.Join(fields, " ")
}
