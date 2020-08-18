package usecase

import "context"

type DB interface {
	Find(context.Context, string) (string, error)
	FindAndIncr(context.Context, string) (string, error)
	Save(context.Context, string, string) error
	Delete(context.Context, string) error
	Count(context.Context, string) (int, error)
	Incr(context.Context, string) error
}
