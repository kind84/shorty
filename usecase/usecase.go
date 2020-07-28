package usecase

import "context"

type App struct {
	db DB
}

func NewApp(db DB) *App {
	return &App{db}
}

// Match returns the URL matching the provided hash.
func (a *App) MatchHash(ctx context.Context, hash string) (string, error) {
	return a.db.FindHash(ctx, hash)
}
