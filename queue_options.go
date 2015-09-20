// Copyright 2015 Jonah Glover. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// add opts here as needed

func Durable(durable bool) func(*Queue) error {
	log.Info("Setting durable")
	return func(q *Queue) error {
		return q.setDurable(durable)
	}
}

func (q *Queue) setDurable(durable bool) error {
	q.durable = durable
	return nil
}
