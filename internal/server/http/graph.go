package http

type Graphql struct {
}

func NewGraphql() (*Graphql, error) {
	return &Graphql{}, nil
}

func (g *Graphql) Handler() {

}

func (g *Graphql) Playground() {

}
