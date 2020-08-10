package usecase

import (
	"context"

	"github.com/kind84/shorty/pkg/hash"
)

type App struct {
	db DB
}

func NewApp(db DB) *App {
	return &App{db}
}

// MatchHash returns the URL matching the provided hash.
func (a *App) MatchHash(ctx context.Context, hash string) (string, error) {
	// retrieve hashed URL
	url, err := a.db.Find(ctx, hash)
	if err != nil {
		return "", err
	}

	// increment redirections counter
	// a.db.Incr(ctx, url)
	return url, nil
}

// CutURL stores the url/hash pair and returns the hasing result.
func (a *App) CutURL(ctx context.Context, url string) (string, error) {
	// run hash func
	hash := hash.Hash(url)

	// store k/v to repo
	err := a.db.Save(ctx, url, hash)
	if err != nil {
		return "", err
	}

	return hash, nil
}

// BurnURL deletes the provided long/short URL.
func (a *App) BurnURL(ctx context.Context, key string) error {
	return a.db.Delete(ctx, key)
}

// InflateURL returns the extenfed version of the provided short URL.
func (a *App) InflateURL(ctx context.Context, short string) (string, error) {
	return a.db.Find(ctx, short)
}
