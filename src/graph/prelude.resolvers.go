package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/introspection"
)

func (r *__DirectiveResolver) IsRepeatable(ctx context.Context, obj *introspection.Directive) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// __Directive returns generated.__DirectiveResolver implementation.
func (r *Resolver) __Directive() generated.__DirectiveResolver { return &__DirectiveResolver{r} }

type __DirectiveResolver struct{ *Resolver }
