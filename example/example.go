package main

import (
	"runtime"
	"time"

	"github.com/leea/muninnode"
)

func main() {
	var ms runtime.MemStats

	mg := muninnode.NewGraph(
		"alloc",
		"Memory Allocated",
		map[string]interface{}{
			"graph_args": "--base 1024",
		},
		func() {
			runtime.ReadMemStats(&ms)
		})
	mg.AddValue("alloc",
		func() interface{} { return ms.Alloc },
		map[string]interface{}{
			"min":  "0",
			"type": "GAUGE",
		})

	muninnode.Register(mg)

	for {
		time.Sleep(10 * time.Second)
	}
}
