// Copyright 2015 Jonah Glover. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hungry

import (
	"github.com/streadway/amqp"
)

type Channel struct {
	Name        string
	queues      map[string]*Queue
	amqpChannel *amqp.Channel
	readPump    chan Message
	writePump   chan Message
}

func (c *Channel) DeclareQueue(name string) error {
	q, err := NewQueue(c.amqpChannel, name)
	failOnError(err, "Failed to create new Queue")
	c.queues[name] = q
	return err
}

func (c *Channel) listen() {
	for {
		msg := <-c.readPump
		log.Info(string(msg.Body))
	}
}

func (c *Channel) listenAndServe() {
	go c.listen()
}

func (c *Channel) getQueue(queueName string) *Queue {
	return c.queues[queueName]
}

func (c *Channel) queueProduce(q *Queue) {

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
	failOnError(err, "Failed to publish a message")
}

func (c *Channel) Consume(queueName string) chan Message {

	queue := c.getQueue(queueName)

	if queue == nil {
		log.Fatal("Could not find queue: " + queueName)
		return nil
	}

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

	queue.bind(msgs)
	return queue.readPump
}

func NewChannel(name string, conn *amqp.Connection) (*Channel, error) {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	c := Channel{Name: name, amqpChannel: ch, queues: make(map[string]*Queue), readPump: make(chan Message), writePump: make(chan Message)}

	c.listenAndServe()

	return &c, err
}
