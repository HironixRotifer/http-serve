package utils

import (
	"context"
	"fmt"
	"math"
	"time"
)

// Split splits data to batches array with size less or equal to batchSize.
func Split[T ~[]U, U any](data T, batchSize int) (batches []T) {
	if batchSize <= 0 {
		panic("batchSize must be greater than 0")
	}

	total := len(data)
	batchesCount := int(math.Ceil(float64(total) / float64(batchSize)))

	batches = make([]T, 0, batchesCount)

	for i := 0; i < batchesCount; i++ {
		j := i * batchSize
		k := (i + 1) * batchSize

		if k > total {
			k = total
		}

		batch := data[j:k]

		batches = append(batches, batch)
	}

	return batches
}

// WriteBatch splits ch's data to batches and flushes them by the overflow or the period time.
// Flushing time is limited by the timeout.
// Loop is alive until the ch is opened.
func WriteBatch[T any](
	ch <-chan T,
	batchSize int,
	flushTimeout time.Duration,
	flushPeriod time.Duration,
	flushFn func(context.Context, []T) error,
) error {
	if ch == nil {
		panic("ch is nil!")
	}
	if batchSize <= 0 {
		panic("batch size must be greater than zero!")
	}
	if flushTimeout <= 0 {
		panic("flush timeout must be greater than zero!")
	}
	if flushPeriod <= 0 {
		panic("flush period must be greater than zero!")
	}
	if flushFn == nil {
		panic("flushFn is nil!")
	}

	batch := make([]T, 0, batchSize)
	ticker := time.NewTicker(flushPeriod)

	flushWithTimeout := func(batch []T, timeout time.Duration) error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		return flushFn(ctx, batch)
	}

writing:
	for {
		select {
		case value, ok := <-ch:
			if !ok {
				break writing
			}

			batch = append(batch, value)

			if len(batch) >= cap(batch) {
				err := flushWithTimeout(batch, flushTimeout)
				if err != nil {
					return fmt.Errorf("flush batch of values: %w", err)
				}
				batch = batch[:0]
			}

		case <-ticker.C:
			if len(batch) > 0 {
				err := flushWithTimeout(batch, flushTimeout)
				if err != nil {
					return fmt.Errorf("flush batch of values: %w", err)
				}
				batch = batch[:0]
			}
		}
	}

	if len(batch) > 0 {
		err := flushWithTimeout(batch, flushTimeout)
		if err != nil {
			return fmt.Errorf("flush batch of values: %w", err)
		}
	}

	return nil
}
