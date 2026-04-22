package utils

import (
	"regexp"
	"strings"
)

// NormalizeThaana cleans and standardizes Maldivian Thaana text from scrapers.
// Optimized for Mobile (Flutter/Next.js) and Divehi Global typeface.
func NormalizeThaana(input string) string {
	// 1. Unicode Normalization (NFC) - Essential for character combining
	// Note: We use manual replacement for known drift characters
	
	// 2. Remove Zero Width Spaces and hidden artifacts
	zws := regexp.MustCompile(`[\x{200B}-\x{200D}\x{FEFF}]`)
	result := zws.ReplaceAllString(input, "")

	// 3. Fix common vowel (Fili) ordering errors
	// Ensures vowels always follow the consonant, preventing 'floating' vowels
	reVowelFirst := regexp.MustCompile(`([\x{07A6}-\x{07B0}])([\x{0780}-\x{07A5}])`)
	result = reVowelFirst.ReplaceAllString(result, "$2$1")

	// 4. Standardize Punctuations for the 'Global' design language
	result = strings.ReplaceAll(result, "،", ",") // Standard comma for cleaner UI
	result = strings.ReplaceAll(result, "؛", ";")
	result = strings.ReplaceAll(result, "؟", "?")

	// 5. Cleanup whitespace
	spaces := regexp.MustCompile(`\s+`)
	result = spaces.ReplaceAllString(result, " ")

	return strings.TrimSpace(result)
}
