package usecase

import (
	"context"
	"errors"

	"github.com/kind84/shorty/pkg/hash"
)

type App struct {
	db DB
}

func NewApp(db DB) *App {
	return &App{db}
}

// MatchHash returns the URL matching the provided hash.
// It also increase the counter of the matched URLs.
func (a *App) MatchHash(ctx context.Context, hash string) (string, error) {
	// retrieve hashed URL
	url, err := a.db.FindAndIncr(ctx, hash)
	if err != nil {
		return "", err
	}
	if url == "" {
		return "", errors.New("URL not found")
	}

	// increment redirections counter
	a.db.Incr(ctx, hash)
	return url, nil
}

// CutURL stores the url/hash pair and returns the hasing result.
func (a *App) CutURL(ctx context.Context, url string) (string, error) {
	// run hash func
	hash := hash.Hash(url)

	// store k/v to repo
	if err := a.db.Save(ctx, url, hash); err != nil {
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

// CountHits returns the number of times the given URL has been hit.
func (a *App) CountHits(ctx context.Context, url string) (int, error) {
	return a.db.Count(ctx, url)
}
