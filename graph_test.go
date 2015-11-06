package muninnode

import (
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mg *Graph
)

func init() {
	mg = NewGraph(
		"test",
		"TestGraph",
		map[string]interface{}{
			"graph_args": "--base 1024",
		},
		nil)

}

func TestGraph(t *testing.T) {
	var ms runtime.MemStats

	mg.AddValue("alloc",
		func() interface{} { return ms.Alloc },
		map[string]interface{}{
			"min":  "0",
			"type": "GAUGE",
		})

	assert.Contains(t, mg.fetch(), "alloc.value", "no value?")
	assert.True(t, strings.HasSuffix(mg.fetch(), "\n."))
	assert.Contains(t, mg.config(), "alloc.label alloc")
	assert.Contains(t, mg.config(), "alloc.min 0")
	assert.Contains(t, mg.config(), "alloc.type GAUGE")
	assert.Contains(t, mg.config(), "graph_args --base 1024")
	assert.True(t, strings.HasSuffix(mg.config(), "\n."))
}

func TestNumName(t *testing.T) {

	mg.AddValue("50.5",
		func() interface{} { return 42 },
		map[string]interface{}{
			"min":  "0",
			"type": "GAUGE",
		})

	assert.Contains(t, mg.fetch(), "_50_5.value")
	assert.Contains(t, mg.config(), "_50_5.label _50_5")

}

func TestPrefetch(t *testing.T) {
	prefetch := 0
	mg1 := NewGraph(
		"test",
		"TestGraph",
		map[string]interface{}{
			"graph_args": "--base 1024",
		},
		func() { prefetch = 1 })

	mg1.fetch()
	assert.Equal(t, prefetch, 1)
}
