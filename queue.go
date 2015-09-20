// Copyright 2015 Jonah Glover. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/streadway/amqp"
)

type Queue struct {
	Name    string
	Key     string
	durable bool
}

func NewQueue(ch *amqp.Channel, name string, options ...func(*Queue) error) (*Queue, error) {

	q := Queue{Name: name}

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
