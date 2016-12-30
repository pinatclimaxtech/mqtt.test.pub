package main

import (
	"fmt"
	"time"

	//import the Paho Go MQTT library

	"strconv"

	"climax.com/mqtt.test.pub/Pub"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("tcp://10.15.8.129:1883")

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	total := 80000
	start := time.Now()

	for i := 1; i <= total; i++ {
		Pub.PubTestTopic(c, strconv.Itoa(i))
	}
	fmt.Println("pubtest finished")
	fmt.Println("elapsed time: ", time.Since(start))
}
