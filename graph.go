package muninnode

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

type Graph struct {
	Name    string
	Title   string
	Globals map[string]interface{}
	Configs map[string]interface{}
	Values  map[string]interface{}

	Gather func() // Call this function prior to calling fetch
}

func muninKey(s string) string {
	var key string
	// Munin does not like keys that start with numbers
	if !unicode.IsLetter(rune(s[0])) {
		key = "_" + s
	} else {
		key = s
	}
	// No dots in key names either
	return strings.Replace(key, ".", "_", -1)
}

func (m *Graph) fetch() string {
	if m.Gather != nil {
		m.Gather()
	}

	lines := make([]string, 0, 10)

	for k, v := range m.Values {
		switch v.(type) {
		case *int:
			lines = append(lines, fmt.Sprintf("%s.value %v", k, *v.(*int)))
		case *uint8:
			lines = append(lines, fmt.Sprintf("%s.value %v", k, *v.(*uint8)))
		case *uint16:
			lines = append(lines, fmt.Sprintf("%s.value %v", k, *v.(*uint16)))
		case *uint32:
			lines = append(lines, fmt.Sprintf("%s.value %v", k, *v.(*uint32)))
		case *uint64:
			lines = append(lines, fmt.Sprintf("%s.value %v", k, *v.(*uint64)))
		case *int8:
			lines = append(lines, fmt.Sprintf("%s.value %v", k, *v.(*int8)))
		case *int16:
			lines = append(lines, fmt.Sprintf("%s.value %v", k, *v.(*int16)))
		case *int32:
			lines = append(lines, fmt.Sprintf("%s.value %v", k, *v.(*int32)))
		case *int64:
			lines = append(lines, fmt.Sprintf("%s.value %v", k, *v.(*int64)))
		case *float32:
			lines = append(lines, fmt.Sprintf("%s.value %v", k, *v.(*float32)))
		case *float64:
			lines = append(lines, fmt.Sprintf("%s.value %v", k, *v.(*float64)))
		default:
			log.Println("value must be a pointer to builtin numeric type")
		}
	}

	return strings.Join(lines, "\n") + "\n."
}

func (m *Graph) config() string {
	lines := make([]string, 0, 10)

	if m.Title != "" {
		lines = append(lines, fmt.Sprintf("graph_title \"%s\"", m.Title))
	} else {
		lines = append(lines, fmt.Sprintf("graph_title \"%s\"", m.Name))
	}

	for k, v := range m.Globals {
		lines = append(lines, fmt.Sprintf("%s %v", k, v))
	}

	for k, v := range m.Configs {
		lines = append(lines, fmt.Sprintf("%s %v", k, v))
	}

	return strings.Join(lines, "\n") + "\n."
}
