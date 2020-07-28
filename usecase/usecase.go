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
	return a.db.Find(ctx, hash)
}

// CutURL stores the url/hash pair and returns the hasing result.
func (a *App) CutURL(ctx context.Context, url string) (string, error) {
	// run hash func
	hash := hash.Hash(url)

	//store k/v to repo
	err := a.db.Save(ctx, hash)
	if err != nil {
		return "", err
	}

	return hash, nil
}

func (a *App) BurnURL(context.Context, string) error

func (a *App) InflateURL(context.Context, string) (string, error)
