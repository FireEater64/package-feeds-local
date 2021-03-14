package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/kafkapubsub"
)

func main() {
	time.Sleep(5 * time.Second)
	fmt.Print("Connecting...")
	endpoint := os.Getenv("PUBSUB_ENDPOINT")

	ctx := context.TODO()
	subs, err := pubsub.OpenSubscription(ctx, endpoint)

	if err != nil {
		panic(err)
	}

	fmt.Print("Connected")

	for {
		msg, err := subs.Receive(ctx)

		if err != nil {
			fmt.Print(err)
			break
		}

		fmt.Printf("Received %s\n", msg.Body)
		msg.Ack()
	}

	defer subs.Shutdown(ctx)
}