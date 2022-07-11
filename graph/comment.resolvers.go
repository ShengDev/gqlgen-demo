package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gqlgen-demo/graph/generated"
	"gqlgen-demo/graph/model"
)

// CommentLabels is the resolver for the CommentLabels field.
func (r *commentResolver) CommentLabels(ctx context.Context, obj *model.Comment) ([]*model.CommentLabel, error) {
	commentLabels, err := r.queryCommentLabelsByCommentId(ctx, obj.Id)
	//return nil, err
	if err != nil {
		return nil, err
	}
	return commentLabels, nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
