package muninnode

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

type Graph struct {
	name      string
	title     string
	configs   map[string]interface{}
	values    []*Value
	prefetchf func()
}

func NewGraph(name string, title string, config map[string]interface{}, pf func()) *Graph {
	return &Graph{
		name,
		title,
		config,
		[]*Value{},
		pf,
	}
}
func (m *Graph) AddValue(name string, getf func() interface{}, configs map[string]interface{}) {

	m.values = append(m.values, &Value{muninKey(name), getf, configs})
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
	if m.prefetchf != nil {
		m.prefetchf()
	}
	lines := make([]string, 0, 10)

	for _, v := range m.values {
		lines = append(lines, v.fetch())
	}

	return strings.Join(lines, "\n") + "\n."
}

func (m *Graph) config() string {
	lines := make([]string, 0, 10)

	if m.title != "" {
		lines = append(lines, fmt.Sprintf("graph_title %s", m.title))
	} else {
		lines = append(lines, fmt.Sprintf("graph_title %s", m.name))
	}

	for k, v := range m.configs {
		lines = append(lines, fmt.Sprintf("%s %v", k, v))
	}

	for _, v := range m.values {
		lines = append(lines, v.config()...)
	}

	return strings.Join(lines, "\n") + "\n."
}

type Value struct {
	name     string
	getValue func() interface{} //Must return builtin numeric type: int, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64
	configs  map[string]interface{}
}

func (v *Value) fetch() string {

	value := v.getValue()
	switch value.(type) {
	case int, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		return fmt.Sprintf("%s.value %v", v.name, value)
	default:
		log.Println("value must be a pointer to builtin numeric type")
		return "\n."
	}

}

func (v *Value) config() []string {
	lines := make([]string, 0, 5)
	lines = append(lines, fmt.Sprintf("%s.label %v", v.name, v.name))

	for k, vv := range v.configs {
		lines = append(lines, fmt.Sprintf("%s.%s %v", v.name, k, vv))
	}

	return lines
}
