package main

import (
	"context"
	"fmt"
	"time"
)

func MainContext() {
	// Timed cancellation
	//timeout := 10 * time.Second
	//ctx, _  := context.WithTimeout(context.Background(), timeout)
	//fmt.Println("##___________Add() result:", add(ctx))

	// Automatic cancellation
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
		fmt.Println("##___________add() result:", add(ctx))
	}()

	select {}
}

func add(ctx context.Context) int {
	ctx = context.WithValue(ctx, "Hello", "Hello")
	ctx = context.WithValue(ctx, "World", "World")
	go fmt.Println("##___________bdd() result:", bdd(ctx))

	select {
	case <-ctx.Done():
		return -1
	default:
		fmt.Printf("##___________%s\n", "add() not end")
		return 0
	}
}

func bdd(ctx context.Context) int {
	fmt.Printf("##___________%s\n", ctx.Value("Hello"))
	fmt.Printf("##___________%s\n", ctx.Value("World"))

	ctx = context.WithValue(ctx, "boy", "zhe")
	go fmt.Println("##___________cdd() result:", cdd(ctx))

	select {
	case <-ctx.Done():
		return -2
	default:
		fmt.Printf("##___________%s\n", "bdd() not end")
		return 0
	}
}

func cdd(ctx context.Context) int {
	fmt.Printf("##___________%s\n", ctx.Value("boy"))
	select {
	case <-ctx.Done():
		return -3
	default:
		fmt.Printf("##___________%s\n", "cdd() not end")
		return 0
	}
}
