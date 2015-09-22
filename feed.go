// Copyright 2015 Jonah Glover. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hungry

import (
	"github.com/streadway/amqp"
)

type Feed struct {
	amqpConn *amqp.Connection

	channels map[string]*Channel

	host     string
	port     string
	user     string
	password string
}

func (f *Feed) Close() {
	f.amqpConn.Close()
	// f.channel.Close()
}

func (f *Feed) DeclareChannel(name string) (*Channel, error) {
	c, err := NewChannel(name, f.amqpConn)
	failOnError(err, "Failed Declaring Channel")
	f.channels[name] = c
	return c, err
}

func NewFeed(options ...func(*Feed) error) (*Feed, error) {

	f := Feed{host: "localhost", port: "5672", user: "guest", password: "guest", channels: make(map[string]*Channel)}

	for _, option := range options {
		option(&f)
	}

	conn, err := amqp.Dial("amqp://" + f.user + ":" + f.password + "@" + f.host + ":" + f.port + "/")

	failOnError(err, "Failed to open up connection")

	f.amqpConn = conn

	return &f, nil
}
