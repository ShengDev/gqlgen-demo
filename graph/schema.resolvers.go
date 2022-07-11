package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gqlgen-demo/graph/generated"
	"gqlgen-demo/graph/model"
	"math/rand"
)

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	comment := &model.Comment{
		Id:             int64(rand.Int()),
		DoctorUin:      111,
		PatientUin:     0,
		OrderId:        0,
		IsShow:         0,
		CommentStars:   input.CommentStars,
		CommentContent: input.CommentContent,
		DbFrom:         "1",
	}
	r.comments = append(r.comments, comment)
	return comment, nil
}

// Doctor is the resolver for the doctor field.
func (r *queryResolver) Doctor(ctx context.Context, input model.DoctorQueryParam) (*model.Doctor, error) {
	//panic(fmt.Errorf("not implemented"))
	doctor, err := r.queryDoctor(ctx, input.Uin)
	if err != nil {
		return nil, err
	}
	return doctor, nil
}

// CommentLabels is the resolver for the commentLabels field.
func (r *queryResolver) CommentLabels(ctx context.Context) ([]*model.CommentLabel, error) {
	//panic(fmt.Errorf("not implemented"))
	commentLabels, err := r.queryCommentLabels(ctx)
	if err != nil {
		return nil, err
	}
	return commentLabels, nil
}

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context, commentQueryParam *model.CommentQueryParam) (*model.Comment, error) {
	comment, err := r.queryComment(ctx, commentQueryParam.Id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
