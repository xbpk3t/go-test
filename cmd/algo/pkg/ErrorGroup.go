package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	errGroup, newCtx := errgroup.WithContext(ctx)

	done := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			errGroup.Go(func() error {
				time.Sleep(time.Second * 10)
				return nil
			})
		}
		if err := errGroup.Wait(); err != nil {
			fmt.Printf("do err:%v\n", err)
			return
		}
		done <- struct{}{}
	}()

	select {
	case <-newCtx.Done():
		fmt.Printf("err:%v ", newCtx.Err())
		return
	case <-done:
	}
	fmt.Println("success")
}
