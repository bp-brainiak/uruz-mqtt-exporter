/*
SPDX-License-Identifier: GPL-2.0-only
Copyright © 2023 Ulises Ruz Puga ulises.ruz@gmail.com

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package mqttlogic

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"os/signal"
	"time"
)

var ConfigData MqttData

func SetConfigData(configData MqttData) {
	ConfigData = configData
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	if ConfigData.Verbose {
		log.Println("Connected")
	}
}
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
	if ConfigData.Verbose {
		log.Println(err)
	}
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	if ConfigData.Verbose {
		log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	}
}

func Connect() {

	fmt.Printf("Connecting with %s please wait...\n", fmt.Sprintf("%s:%d", ConfigData.Server, ConfigData.Port))
	ops := mqtt.NewClientOptions()
	ops.AddBroker(fmt.Sprintf("%s:%d", ConfigData.Server, ConfigData.Port))
	ops.SetClientID("bpbrainiak-mqtt-exporter")
	ops.SetUsername(ConfigData.User)
	ops.SetPassword(ConfigData.Pass)
	ops.OnConnect = connectHandler
	ops.OnConnectionLost = connectLostHandler
	ops.SetDefaultPublishHandler(messagePubHandler)
	client := mqtt.NewClient(ops)
	defer client.Disconnect(1000 * 3600)
	// throw an error if the connection isn't successfull
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	topics := ConfigData.Topics
	for _, topic := range topics {
		subscribe(client, topic)
	}
	go func() {
		for {
			time.Sleep(1 * time.Second)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func subscribe(client mqtt.Client, topic string) {
	// subscribe to the same topic, that was published to, to receive the messages
	token := client.Subscribe(topic, 1, messagePubHandler)
	token.Wait()
	// Check for errors during subscribe (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		log.Printf("Failed to subscribe to topic: %s\n", topic)
		panic(token.Error())
	}
	log.Printf("Subscribed to topic: %s\n", topic)
}
