package graph

import (
	"context"
	"gqlgen-demo/graph/model"
)

// This file will not be regenerated automatically.
// DB/cache...
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	doctors []*model.Doctor
	commentLabels []*model.CommentLabel
	comments []*model.Comment
}

func (*Resolver) queryDoctor(ctx context.Context, uin int64) (*model.Doctor, error) {
	doctors := []model.Doctor{
		{
			Uin:        1,
			Name:       "default",
			Department: "default",
			Title:      "default",
			Supplier:   "default",
		},
		{
			Uin:        111,
			Name:       "doctor111",
			Department: "儿科",
			Title:      "主任医师",
			Supplier:   "腾讯医生",
		},
		{
			Uin:        112,
			Name:       "doctor112",
			Department: "内科",
			Title:      "副主任医师",
			Supplier:   "杏仁",
		},
	}
	resp := doctors[0]
	for _, doctor := range doctors{
		if doctor.Uin == uin {
			resp = doctor
		}
	}
	return &resp, nil
}

func (*Resolver) queryComment(ctx context.Context, id int64) (*model.Comment, error) {
	comments := []model.Comment{
		{
			Id: 1,
			DoctorUin: 111,
			PatientUin: 0,
			OrderId: 2022,
			OperatorUin: "2333",
			IsShow: 0,
			CommentStars: 5,
			CommentContent: "医生很好",
			CreateTime: 32354534354,
			UpdateTime: 34535345577,
			DbFrom: "0",
		},
		{
			Id: 2,
			DoctorUin: 111,
			PatientUin: 0,
			OrderId: 2022,
			OperatorUin: "",
			IsShow: 0,
			CommentStars: 5,
			CommentContent: "医生很好",
			CreateTime: 32354534354,
			UpdateTime: 34535345577,
			DbFrom: "0",
		},
	}
	resp := comments[0]
	for _, comment := range comments{
		if comment.Id == id {
			resp = comment
		}
	}
	return &resp, nil
}

func (*Resolver) queryCommentLabels(ctx context.Context) ([]*model.CommentLabel, error) {
	commentLabels := []*model.CommentLabel{
		{
			Id: 1,
			CommentLabel: "热情大方",
			StarRating: 5,
		},
		{
			Id: 2,
			CommentLabel: "医术高明",
			StarRating: 4,
		},
		{
			Id: 3,
			CommentLabel: "关怀备至",
			StarRating: 3,
		},
		{
			Id: 4,
			CommentLabel: "态度不好",
			StarRating: 2,
		},
		{
			Id: 5,
			CommentLabel: "不解决问题",
			StarRating: 1,
		},
	}
	return commentLabels, nil
}

func (*Resolver) queryCommentLabelsByCommentId(ctx context.Context, commentId int64) ([]*model.CommentLabel, error) {
	commentLabels := []*model.CommentLabel{
		{
			Id: 2,
			CommentLabel: "医术高明",
			StarRating: 4,
		},
		{
			Id: 3,
			CommentLabel: "关怀备至",
			StarRating: 3,
		},
	}
	return commentLabels, nil
}