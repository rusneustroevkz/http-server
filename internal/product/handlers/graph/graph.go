package graph

type ProductsGraph struct {
}

func NewProductsGraph() *ProductsGraph {
	return &ProductsGraph{}
}

func (*ProductsGraph) Pattern() string {
	return "/graph"
}

func (*ProductsGraph) PlaygroundPattern() string {
	return "/graph-playground"
}
