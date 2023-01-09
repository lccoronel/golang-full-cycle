package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(5 * time.Second):
		println("Hotel booked.")
	}
}
