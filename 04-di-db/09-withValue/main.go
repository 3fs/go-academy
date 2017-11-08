package main

import (
	"context"
	"fmt"
)

type contextKey string

const authCtxKey = contextKey("AUTH")
const userIDCtxKey = contextKey("USERID")

func main() {
	ctx := context.WithValue(context.Background(), authCtxKey, "authenticationToken=")
	ctx = fakeExternalCall(ctx)
	serveData(ctx)
}

func fakeExternalCall(ctx context.Context) context.Context {
	v := ctx.Value(authCtxKey).(string)
	fmt.Printf("Received token: %s\n", v)
	return context.WithValue(ctx, userIDCtxKey, "SOME-UUID")
}

func serveData(ctx context.Context) {
	userID := ctx.Value(userIDCtxKey).(string)
	fmt.Printf("Welcome %s\n", userID)
}
