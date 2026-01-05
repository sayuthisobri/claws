package view

import (
	"slices"
	"strings"
)

// fuzzyMatch checks if pattern characters appear in order in str (case insensitive)
// Supports special prefix patterns:
// - "^pattern" - Exact substring match from start (prefix match)
// - "*pattern" - Contains substring anywhere (contains match)
// - "$pattern" - Exact substring match at end (suffix match)
// - "pattern" - Default fuzzy subsequence matching
func fuzzyMatch(str, pattern string) bool {
	str = strings.ToLower(str)

	// Check for special prefix patterns
	if len(pattern) > 0 {
		switch pattern[0] {
		case '^':
			// Prefix match: pattern must appear at the start of str
			// If pattern is just "^", it matches empty string at start (always true)
			if len(pattern) == 1 {
				return true
			}
			return strings.HasPrefix(str, pattern[1:])
		case '*':
			// Contains match: pattern must appear anywhere in str
			// If pattern is just "*", it matches empty string anywhere (always true)
			if len(pattern) == 1 {
				return true
			}
			return strings.Contains(str, pattern[1:])
		case '$':
			// Suffix match: pattern must appear at the end of str
			// If pattern is just "$", it matches empty string at end (always true)
			if len(pattern) == 1 {
				return true
			}
			return strings.HasSuffix(str, pattern[1:])
		}
	}

	// Default fuzzy subsequence matching
	pi := 0
	for i := 0; i < len(str) && pi < len(pattern); i++ {
		if str[i] == pattern[pi] {
			pi++
		}
	}
	return pi == len(pattern)
}

// matchNamesWithFallback returns names matching the pattern.
// It first tries prefix matching, then falls back to fuzzy matching if no prefix matches.
func matchNamesWithFallback(names []string, pattern string) []string {
	if pattern == "" {
		result := slices.Clone(names)
		slices.Sort(result)
		return result
	}

	pattern = strings.ToLower(pattern)

	var prefixMatches []string
	for _, name := range names {
		if strings.HasPrefix(strings.ToLower(name), pattern) {
			prefixMatches = append(prefixMatches, name)
		}
	}
	if len(prefixMatches) > 0 {
		slices.Sort(prefixMatches)
		return prefixMatches
	}

	var fuzzyMatches []string
	for _, name := range names {
		if fuzzyMatch(name, pattern) {
			fuzzyMatches = append(fuzzyMatches, name)
		}
	}
	slices.Sort(fuzzyMatches)
	return fuzzyMatches
}
