package usecase

import "context"

type DB interface {
	FindHash(context.Context, string) (string, error)
	Save(context.Context, string) (string, error)
	Delete(context.Context, string) error
}
