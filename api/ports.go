package api

import "context"

type App interface {
	MatchHash(context.Context, string) (string, error)
	CutURL(context.Context, string) (string, error)
	BurnURL(context.Context, string) error
	InflateURL(context.Context, string) (string, error)
	CountHits(context.Context, string) (int, error)
}
