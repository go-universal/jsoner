package jsoner

import (
	"slices"
	"strings"
)

// newFilter creates a new filter instance, excluding any empty or whitespace-only strings
// from the provided filters.
func newFilter(filters ...string) *filter {
	// Remove empty or whitespace-only strings from the filters slice.
	cleanedFilters := slices.DeleteFunc(filters, func(item string) bool {
		return strings.TrimSpace(item) == ""
	})

	return &filter{
		exclude: cleanedFilters,
	}
}

// filter represents a structure that holds a list of fields to exclude.
type filter struct {
	exclude []string
}

// shouldSkip checks if a given field should be skipped based on the exclude list.
// It returns true if the field is empty or is present in the exclude list.
func (f *filter) shouldSkip(field string) bool {
	return field == "" || slices.Contains(f.exclude, field)
}
