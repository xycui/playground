package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/xycui/playground/container"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	task1 := NewTask(func(ctx context.Context) (interface{}, error) {
		future := time.Now().Unix() + 10
		breakTag := false

		for time.Now().Unix() < future && !breakTag {
			select {
			case <-ctx.Done():
				breakTag = true
			default:
				fmt.Println("------task 1-------")
				time.Sleep(time.Second)
			}
		}

		return nil, nil
	}, ctx)

	task2 := NewTask(func(ctx context.Context) (interface{}, error) {
		future := time.Now().Unix() + 7
		for time.Now().Unix() < future {
			fmt.Println(">>>>>>>task 2<<<<<")
			time.Sleep(time.Millisecond * 200)
		}

		return nil, nil
	}, ctx)

	WhenAll(task1, task2)
	return

	builder := container.Level0Builder
	cmd := builder.Build(os.Args[1:])
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

type ExecFunc func(context.Context) (interface{}, error)

type Task struct {
	Ctx      context.Context
	ExecFunc ExecFunc
	result   interface{}
	err      error
}

func (t *Task) Result() interface{} {
	return t.result
}

func (t *Task) Error() error {
	return t.err
}

func NewTask(execFunc ExecFunc, ctx context.Context) *Task {
	return &Task{
		Ctx:      ctx,
		ExecFunc: execFunc,
	}
}

func WhenAll(tasks ...*Task) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func(t *Task) {
			ret, err := t.ExecFunc(t.Ctx)
			t.result = ret
			t.err = err
			wg.Done()
		}(task)
	}

	wg.Wait()
}

func WhenAny(tasks ...*Task) {

}
