package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "some key1", "some value1")
	ctx = context.WithValue(ctx, "some key2", "some value2")

	//	fmt.Println(ctx.Value("some key1"))
	//	fmt.Println(ctx.Value("some key2"))

	do(ctx)
}

func do(ctx context.Context) {
	fmt.Println(ctx.Value("some key1"))
	fmt.Println(ctx.Value("some key2"))
}
