package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sethvargo/go-retry"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	log := log.With().
		Uint64("target_epoch_counter", 1).
		Str("target_epoch_phase", "target_epoch_phase1").
		Logger()

	log.Info().Msg("starting dynamic startup - waiting until target epoch/phase to start...")
	count := 1
	backoff := retry.NewConstant(time.Duration(3))
	err := retry.Do(ctx, backoff, func(ctx context.Context) error {

		// wait then poll for latest snapshot again
		log.Info().
			Dur("time-waiting", time.Since(start)).
			Uint64("current-epoch", uint64(count)).
			Msgf("等待 epoch %d and phase %s", count, "urrent-epoch-phase")
		if count == 3 {
			count = count + 1
			return nil
		}
		if count == 2 {
			count = count + 1
			return fmt.Errorf("返回一般err")
		}
		if count == 1 {
			count = count + 1
			return retry.RetryableError(fmt.Errorf("需要重试 : %v", count))
		}
		return nil
	})
	if err != nil {
		fmt.Printf("重试报错: %v\n", err)
	}
}
