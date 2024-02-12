package templates

import "context"

func isBoosted(ctx context.Context) bool {
	return ctx.Value("isBoosted") != nil
}
