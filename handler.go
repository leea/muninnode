package muninnode

import (
	"errors"
	"sort"
	"strings"
)

var (
	errUnknownCmd = errors.New("Unknown command")
	errUnknownSvc = errors.New("Unknown service")
)

func handler(command string) (r string, err error) {
	fields := strings.Fields(command)
	if len(fields) == 0 {
		return "", errUnknownCmd
	}

	switch fields[0] {

	case "list":
		keys := []string{}
		for _, m := range registry {
			keys = append(keys, m.Name)
		}
		sort.Strings(keys)
		return strings.Join(keys, " "), nil

	case "nodes":
		return hostname + "\n.", nil

	case "fetch":
		if len(fields) < 2 {
			return "", errUnknownSvc
		}
		key := fields[1]

		for _, m := range registry {
			if m.Name == key {
				return m.fetch(), nil
			}
		}
		return "", errUnknownSvc

	case "config":
		if len(fields) < 2 {
			return "", errUnknownSvc
		}
		key := fields[1]

		for _, m := range registry {
			if m.Name == key {
				return m.config(), nil
			}
		}
		return "", errUnknownSvc

	case "cap":
		return "multigraph", nil

	case "help", "version":
		fallthrough

	default:
		return "", errUnknownCmd

	}
}
