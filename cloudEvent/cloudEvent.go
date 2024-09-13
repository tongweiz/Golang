package main

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func main() {
	event := cloudevents.NewEvent()
	event.SetExtension("test", "string value")
	print(event.String())
}
