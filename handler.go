package evmn

import (
	"errors"
	"expvar"
	"sort"
	"strings"
	"unicode"
)

var (
	errUnknownCmd = errors.New("Unknown command")
	errUnknownSvc = errors.New("Unknown service")
)

func kk(k string) string {
	if k == "" {
		return k
	}
	k = strings.Replace(k, ".", "_", -1)
	if !unicode.IsLetter(rune(k[0])) {
		return "_" + k
	}
	return k
}

func handler(command string) (r string, err error) {
	fields := strings.Fields(command)
	if len(fields) == 0 {
		return "", errUnknownCmd
	}

	switch fields[0] {

	case "list":
		keys := []string{}
		seen := map[string]bool{}
		do(func(kv expvar.KeyValue) {
			g := strings.Split(kv.Key, ":")[0]
			f := strings.Split(g, ".")[0]
			if seen[f] {
				return
			}
			keys = append(keys, f)
			seen[f] = true
		})
		sort.Strings(keys)
		return strings.Join(keys, " "), nil

	case "nodes":
		return hostname + "\n.", nil

	case "fetch":
		if len(fields) < 2 {
			return "", errUnknownSvc
		}
		key := fields[1]

		lines := []string{}
		fetch(key, func(k, v string) {
			lines = append(lines, kk(k)+".value "+v)
		})
		return strings.Join(lines, "\n") + "\n.", nil

	case "config":
		if len(fields) < 2 {
			return "", errUnknownSvc
		}
		key := fields[1]

		lines := []string{
			"graph_title " + key,
			"graph_category expvar",
			"graph_args --base 1000 --units=si",
		}

		fetch(key, func(k0, v string) {
			k := kk(k0)
			lines = append(lines,
				k+".label "+k0,
				k+".min 0",
				k+".type DERIVE",
			)
		})
		return strings.Join(lines, "\n") + "\n.", nil

	case "cap":
		return "multigraph", nil

	case "help", "version":
		fallthrough

	default:
		return "", errUnknownCmd

	}
}
