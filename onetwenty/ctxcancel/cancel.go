package ctxcancel

import (
	"context"
	"log"
	"time"
)

func Process(ctx context.Context, x, y int) error {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		return context.Cause(ctx)
	case <-time.After(3 * time.Second):
		return nil
	}
}
