package types

import (
	poststypes "github.com/desmos-labs/desmos/x/staging/posts/types"
	reportstypes "github.com/desmos-labs/desmos/x/staging/reports/types"
)

type Post struct {
	*poststypes.Post
	Height int64
}

func NewPost(post *poststypes.Post, height int64) Post {
	return Post{
		Post:   post,
		Height: height,
	}
}

// --------------------------------------------------------------------------------------------------------------------

type UserPollAnswer struct {
	poststypes.UserAnswer
	PostID string
	Height int64
}

func NewUserPollAnswer(postID string, answer poststypes.UserAnswer, height int64) UserPollAnswer {
	return UserPollAnswer{
		UserAnswer: answer,
		PostID:     postID,
		Height:     height,
	}
}

// --------------------------------------------------------------------------------------------------------------------

type RegisteredReaction struct {
	poststypes.RegisteredReaction
	Height int64
}

func NewRegisteredReaction(reaction poststypes.RegisteredReaction, height int64) RegisteredReaction {
	return RegisteredReaction{
		RegisteredReaction: reaction,
		Height:             height,
	}
}

// --------------------------------------------------------------------------------------------------------------------

type PostReaction struct {
	poststypes.PostReaction
	PostID string
	Height int64
}

func NewPostReaction(postID string, reaction poststypes.PostReaction, height int64) PostReaction {
	return PostReaction{
		PostID:       postID,
		PostReaction: reaction,
		Height:       height,
	}
}

// --------------------------------------------------------------------------------------------------------------------

type Report struct {
	reportstypes.Report
	Height int64
}

func NewReport(report reportstypes.Report, height int64) Report {
	return Report{
		Report: report,
		Height: height,
	}
}