package utils

import (
	"context"
	"time"
)

func ContextTimoutNSecond(duration int) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(duration))
	return ctx, cancel
}
