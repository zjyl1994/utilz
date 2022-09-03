package utilz

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

func WorkerDo[K any](ctx context.Context, workerNum int, inputs []K,
	fn func(context.Context, K) error) error {
	if workerNum <= 0 {
		return AsyncDo(ctx, inputs, fn)
	}
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
	if workerNum <= 0 {
		return AsyncResult(ctx, inputs, fn)
	}
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

func AsyncDo[K any](ctx context.Context, inputs []K,
	fn func(context.Context, K) error) error {
	worker := func(ctx context.Context, jobInput K) func() error {
		return func() error { return fn(ctx, jobInput) }
	}
	waitg, ctx := errgroup.WithContext(ctx)
	for _, input := range inputs {
		waitg.Go(worker(ctx, input))
	}
	return waitg.Wait()
}

func AsyncResult[K, V any](ctx context.Context, inputs []K,
	fn func(context.Context, K) (V, error)) ([]V, error) {
	var resultLock sync.Mutex
	results := make([]V, 0, len(inputs))
	worker := func(ctx context.Context, jobInput K) func() error {
		return func() error {
			if result, err := fn(ctx, jobInput); err != nil {
				return err
			} else {
				resultLock.Lock()
				results = append(results, result)
				resultLock.Unlock()
				return nil
			}
		}
	}
	waitg, ctx := errgroup.WithContext(ctx)
	for _, input := range inputs {
		waitg.Go(worker(ctx, input))
	}
	err := waitg.Wait()
	return results, err
}
