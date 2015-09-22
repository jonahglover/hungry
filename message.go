// Copyright 2015 Jonah Glover. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hungry

type Message struct {
	Body      []byte
	QueueName string
}
