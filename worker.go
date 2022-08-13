package utilz

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

func WorkerDo[K any](ctx context.Context, workerNum int, inputs []K,
	fn func(context.Context, K) error) error {
	jobCh := make(chan K, workerNum)
	waitg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < workerNum; i++ {
		waitg.Go(func() error {
			for input := range jobCh {
				err := fn(ctx, input)
				if err != nil {
					return err
				}
			}
			return nil
		})
	}
	for _, input := range inputs {
		jobCh <- input
	}
	close(jobCh)
	return waitg.Wait()
}

func WorkerResult[K, V any](ctx context.Context, workerNum int, inputs []K,
	fn func(context.Context, K) (V, error)) ([]V, error) {
	var resultLock sync.Mutex
	results := make([]V, 0, len(inputs))
	jobCh := make(chan K, workerNum)
	waitg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < workerNum; i++ {
		waitg.Go(func() error {
			for input := range jobCh {
				result, err := fn(ctx, input)
				if err != nil {
					return err
				} else {
					resultLock.Lock()
					results = append(results, result)
					resultLock.Unlock()
				}
			}
			return nil
		})
	}
	for _, input := range inputs {
		jobCh <- input
	}
	close(jobCh)
	err := waitg.Wait()
	return results, err
}
