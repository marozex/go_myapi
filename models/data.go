package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}

	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "second comment",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:          1,
		Title:       "first article",
		Contents:    "this is test article",
		UserName:    "testuser",
		NiceNum:     345,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}
)
