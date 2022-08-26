package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

type PubSubConfig struct {
	ProjectId string
	Topics    []PubSubTopicConfig
}

type PubSubTopicConfig struct {
	TopicId       string
	Subscriptions []TopicSubscriptionConfig
}

type TopicSubscriptionConfig struct {
	SubscriptionId     string
	AckDeadlineSeconds int
	PushEndpoint       string
}

func main() {
	pubSubConfigJson := os.Getenv("PUBSUB_CONFIG")
	fmt.Println("pubsub-provisioner config: " + pubSubConfigJson)

	pubSubConfigByteArray := []byte(pubSubConfigJson)

	var pubSubConfig PubSubConfig
	_ = json.Unmarshal(pubSubConfigByteArray, &pubSubConfig)

	ctx := context.Background()

	client, _ := pubsub.NewClient(ctx, pubSubConfig.ProjectId)

	defer client.Close()

	for _, topicConfig := range pubSubConfig.Topics {
		topic, _ := client.CreateTopic(ctx, topicConfig.TopicId)
		fmt.Println("Created topic: " + topicConfig.TopicId)

		for _, subscriptionConfig := range topicConfig.Subscriptions {
			client.CreateSubscription(ctx, subscriptionConfig.SubscriptionId, pubsub.SubscriptionConfig{
				Topic:       topic,
				AckDeadline: time.Duration(subscriptionConfig.AckDeadlineSeconds) * time.Second,
				PushConfig:  pubsub.PushConfig{Endpoint: subscriptionConfig.PushEndpoint},
			})
			fmt.Println("Created subscription: " + subscriptionConfig.SubscriptionId + " for topic + " + topicConfig.TopicId)
		}
	}

	fmt.Println("Pub/Sub provisioning complete!")
}
