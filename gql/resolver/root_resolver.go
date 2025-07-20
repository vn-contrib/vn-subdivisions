package resolver

func NewRootResolver() *rootResolver {
	return &rootResolver{
		query: &queryResolver{},
	}
}

type rootResolver struct {
	query *queryResolver
}

func (r *rootResolver) Query() *queryResolver {
	return r.query
}
