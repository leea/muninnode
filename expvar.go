package evmn

import (
	"expvar"
	"strings"
)

func do(f func(expvar.KeyValue)) {
	expvar.Do(func(kv expvar.KeyValue) {
		switch kv.Key {
		case "cmdline",
			"memstats":
			return
		}
		if _, ok := kv.Value.(*expvar.Int); !ok {
			return
		}
		f(kv)
	})
}

func fetch(name string, f func(k, v string)) {
	do(func(kv expvar.KeyValue) {
		if !strings.HasPrefix(kv.Key, name+":") {
			return
		}
		lst := strings.Split(kv.Key, ":")
		f(lst[1], kv.Value.String())
	})
}
