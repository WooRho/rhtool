package rgroup

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestErrGroup(T *testing.T) {
	ctx := context.TODO()
	err := Finish(ctx, 2, func(ctx context.Context) error {
		err := errors.New("1111err")
		time.Sleep(time.Second)
		fmt.Println(111111)
		return err
	}, func(ctx context.Context) error {
		time.Sleep(time.Second)
		fmt.Println(222222)
		return nil
	}, func(ctx context.Context) error {
		time.Sleep(1 * time.Second)
		fmt.Println(333333)
		return nil
	}, func(ctx context.Context) error {
		time.Sleep(time.Second)
		fmt.Println(44444)
		return nil
	}, func(ctx context.Context) error {
		time.Sleep(time.Second)
		fmt.Println(55555)
		return nil
	},
	)

	fmt.Println("-----------")
	fmt.Println(err)

	// 由于是并发有一些拦不住还是会执行到
	// test run 1
	//111111
	//ctx.Done()
	//ctx.Done()
	//333333
	//44444
	//-----------
	//	1111err

	// test run 2
	//111111
	//ctx.Done()
	//ctx.Done()
	//ctx.Done()
	//ctx.Done()
	//-----------
	//	1111err
}

func TestErrGroupV(T *testing.T) {
	ctx := context.TODO()
	err := FinishVoidErr(ctx, 5, func(ctx context.Context) error {
		fmt.Println(111111)
		err := errors.New("dddd")
		return err
	}, func(ctx context.Context) error {
		time.Sleep(time.Second)
		fmt.Println(222222)
		return nil
	}, func(ctx context.Context) error {
		time.Sleep(time.Second)
		fmt.Println(333333)
		return nil
	}, func(ctx context.Context) error {
		time.Sleep(time.Second)
		fmt.Println(44444)
		return nil
	}, func(ctx context.Context) error {
		time.Sleep(time.Second)
		fmt.Println(55555)
		return nil
	})

	fmt.Println(err)
	//=== RUN   TestErrGroupV
	//111111
	//222222
	//333333
	//44444
	//55555
	//--- PASS: TestErrGroupV (0.00s)
	//PASS
}
