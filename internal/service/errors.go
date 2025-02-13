package service

import "errors"

var ErrEmptyTitle = errors.New("title is empty")
var ErrEmptyDescription = errors.New("description is empty")
var ErrInvalidPriority = errors.New("priority is invalid")
