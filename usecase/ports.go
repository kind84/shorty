package usecase

import "context"

type DB interface {
	Find(context.Context, string) (string, error)
	Save(context.Context, string) error
	Delete(context.Context, string) error
}
