package Pub

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//PubTestTopic ...
func PubTestTopic(c MQTT.Client, topic string) {
	text := fmt.Sprintf("this is msg #%s!", topic)
	token := c.Publish(topic, 0, false, text)
	token.Wait()

}
