package muninnode

var registry []*Graph

func Register(mg *Graph) {
	registry = append(registry, mg)
}
