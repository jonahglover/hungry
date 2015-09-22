// Copyright 2015 Jonah Glover. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE fedle.

package hungry

import (
	"github.com/streadway/amqp"
)

type Queue struct {
	Name      string
	Key       string
	durable   bool
	readPump  chan Message
	writePump chan Message
}

func (q *Queue) bindMessages(msgs <-chan amqp.Delivery) {
	log.Info("Binding messages to queue " + q.Name)
	for {
		for d := range msgs {
			q.readPump <- Message{Body: d.Body}
		}
	}
}

func (q *Queue) bind(msgs <-chan amqp.Delivery) {
	go q.bindMessages(msgs)
}

func NewQueue(ch *amqp.Channel, name string, options ...func(*Queue) error) (*Queue, error) {

	q := Queue{Name: name, readPump: make(chan Message), writePump: make(chan Message)}

	for _, option := range options {
		option(&q)
	}

	mq, err := ch.QueueDeclare(
		q.Name,    // name
		q.durable, // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	q.Key = mq.Name

	return &q, err
}
