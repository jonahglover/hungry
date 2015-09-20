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

func (c *Channel) Send(queueName string, body string) {
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

func NewChannel(name string, conn *amqp.Connection) (*Channel, error) {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	c := Channel{Name: name, amqpChannel: ch, queues: make(map[string]*Queue)}
	return &c, err
}
