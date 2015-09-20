// Copyright 2015 Jonah Glover. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func Host(host string) func(*Feed) error {
	log.Info("Setting host: " + host)
	return func(f *Feed) error {
		return f.setHost(host)
	}
}

func (f *Feed) setHost(host string) error {
	f.host = host
	return nil
}

func Port(port string) func(*Feed) error {
	log.Info("Setting port: " + port)
	return func(f *Feed) error {
		return f.setPort(port)
	}
}

func (f *Feed) setPort(port string) error {
	f.port = port
	return nil
}

func User(user string) func(*Feed) error {
	log.Info("Setting user: " + user)
	return func(f *Feed) error {
		return f.setUser(user)
	}
}

func (f *Feed) setUser(user string) error {
	f.user = user
	return nil
}

func Password(password string) func(*Feed) error {
	log.Info("Setting password: " + password)
	return func(f *Feed) error {
		return f.setPassword(password)
	}
}

func (f *Feed) setPassword(password string) error {
	f.password = password
	return nil
}
