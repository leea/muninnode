package main

import (
	"runtime"
	"time"

	"github.com/leea/muninnode"
)

func main() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	mg := &muninnode.Graph{
		Name:  "alloc",
		Title: "Memory Allocated",
		Args: map[string]interface{}{
			"base": 1024},
		Configs: map[string]interface{}{
			"alloc.min":  "0",
			"alloc.type": "GAUGE",
		},
		Values: map[string]interface{}{
			"alloc": &ms.Alloc,
		},
		Gather: func() { runtime.ReadMemStats(&ms) },
	}
	muninnode.Register(mg)

	for {
		time.Sleep(10 * time.Second)
	}
}
