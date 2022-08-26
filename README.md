# pubsub-provisioner

A simple program for provisioning Pub/Sub topics and subscriptions for the [firebase-tools](https://github.com/firebase/firebase-tools) emulator for local testing.

# How to use

See a working example of how to use it [here](https://github.com/JonnyOrman/pubsub-provisioner-example).

In a `docker-compose` file with a `firebase-tools` service, add another service that pulls the [pubsub-provisioner docker image](https://hub.docker.com/repository/docker/jonnyorman/pubsub-provisioner).

On your `pubsub-provisioner` service, set the `PUBSUB_EMULATOR_HOST` environment variable and pass the JSON config for your Pub/Sub in this format:
```
{
    "projectId": "pubsub-provisioner-example",
    "topics": [
        {
            "topicId": "my-topic",
            "subscriptions": [
                {
                    "type": "push",
                    "subscriptionId": "my-topic-subscription",
                    "ackDeadlineSeconds": 10,
                    "pushEndpoint": "http://my-push-endpoint:1234"
                }
            ]
        }
    ]
}
```