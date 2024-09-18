package mqtt

import (
	"goparkin_service/config"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MQTTClient mqtt.Client

func InitMQTT() {
    opts := mqtt.NewClientOptions().AddBroker(config.AppConfig.MQTTBroker)
    opts.SetClientID("mqtt_client")
    MQTTClient = mqtt.NewClient(opts)
    if token := MQTTClient.Connect(); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }
}