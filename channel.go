// Copyright 2015 Jonah Glover. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/streadway/amqp"
)

type Channel struct {
	Name        string
	queues      map[string]*Queue
	amqpChannel *amqp.Channel
}

func (c *Channel) DeclareQueue(name string) error {
	q, err := NewQueue(c.amqpChannel, name)
	failOnError(err, "Failed to create new Queue")
	c.queues[name] = q
	return err
}

func (c *Channel) Publish(queueName string, body string) {
	q := c.queues[queueName]

	if q == nil {
		log.Critical("Could not find queue")
		return
	}

	err := c.amqpChannel.Publish(
		"",    // exchange
		q.Key, // Queue Key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Info(" [x] Sent %s to %s", body, queueName)
	failOnError(err, "Failed to publish a message")
}

func (c *Channel) Consume(queueName string) {

	msgs, err := c.amqpChannel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)

	failOnError(err, "Failed to register a consumer.")
	go func() {
		for d := range msgs {
			log.Info("Received a message: " + string(d.Body))
		}
	}()
}

func NewChannel(name string, conn *amqp.Connection) (*Channel, error) {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	c := Channel{Name: name, amqpChannel: ch, queues: make(map[string]*Queue)}
	return &c, err
}
