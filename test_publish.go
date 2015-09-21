// Copyright 2015 Jonah Glover. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"
)

func main() {
	feed, err := NewFeed()
	if err != nil {
		fmt.Println("something went wrong creating feed")
	}

	ch, err := feed.DeclareChannel("test")

	if ch == nil {
		fmt.Println("something went wrong creating channel")
	}

	if err != nil {
		fmt.Println("something went wrong creating channel")
	}

	qErr := ch.DeclareQueue("testQ")

	if qErr != nil {
		fmt.Println("something went wrong creating queue")
	}

	for {
		ch.Publish("testQ", "Hello, world")
		time.Sleep(100 * time.Millisecond)
	}

	feed.Close()

}
