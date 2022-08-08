package errorhandling

import "strings"

type InvalidField struct {
	FieldName string      `json:"fieldName"`
	Path      string      `json:"path"`
	Reason    string      `json:"reason"`
	Value     interface{} `json:"value"`
}

// Adds path segment to start of path
func (field *InvalidField) PrependPath(paths ...string) {
	if len(paths) == 0 {
		return
	}

	// combine args with field's paths
	combinedPaths := append(paths, field.Path)

	// Filter empty paths
	nonEmptyPaths := []string{}
	for _, path := range combinedPaths {
		if path != "" {
			nonEmptyPaths = append(nonEmptyPaths, path)
		}
	}

	field.Path = strings.Join(nonEmptyPaths, ".")
}
