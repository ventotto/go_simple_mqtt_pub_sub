package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"net/url"
	"os"
	"time"
)

const MQTT_URL_KEY = "MQTT_URL"


func connection(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}


func Listen(topic string) {
	uri, err := url.Parse(os.Getenv(MQTT_URL_KEY))
	if err != nil {
		log.Fatal(err)
	}
	client := connection("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func Publisher(topic, message string){
	uri, err := url.Parse(os.Getenv(MQTT_URL_KEY))
	if err != nil {
		log.Fatal(err)
	}
	client := connection("pub", uri)
	client.Publish(topic, 0, false, message)
}
