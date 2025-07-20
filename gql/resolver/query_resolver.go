package resolver

type queryResolver struct{}

func (r *queryResolver) Ping() string {
	return "pong"
}
