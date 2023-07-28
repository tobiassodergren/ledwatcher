// https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang

package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var verbose bool
var client mqtt.Client

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	if verbose {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	if verbose {
		fmt.Println("Connected")
	}
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	if verbose {
		fmt.Printf("Connect lost: %v", err)
	}
}

func Send(data string) {
	client.Publish("panna/led", 0, true, data)
}

func Init(isVerbose bool, broker string, port int, username, password string) {
	verbose = isVerbose
	var opts = mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("pelletspanna")
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
