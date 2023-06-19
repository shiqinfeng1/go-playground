package main

import (
	"context"
	"fmt"
	"go-playground/cmd/clib/hash"
)

func main() {
	ctx := context.Background()
	p := [32]byte{'1', '2', '3', '4', '1', '2', '3', '4', '1', '2', '3', '4', '1', '2', '3', '4'}
	s := [16]byte{'5', '6', '7', '8', '5', '6', '7', '8', '5', '6', '7', '8', '5', '6', '7', '8'}
	buf := []byte{'5', '6', '7', '8', '5', '6', '7', '5'}
	h := hash.GetHash(ctx, p[:], s[:], buf, uint64(len(buf)))
	fmt.Println("Get hash=", h)
}
