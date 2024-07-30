/*
 * Copyright (C) 2018 Aurélien Chabot <aurelien@chabot.fr>
 *
 * SPDX-License-Identifier: MIT
 */

package main

import (
	"context"
	"log"

	"github.com/hekmon/transmissionrpc/v2"
)

// Transmission handle the transmission api request
type Transmission struct {
	client *transmissionrpc.Client
}

// NewTransmission return a new Transmission object
func NewTransmission(host string, port uint16, user string, pswd string) *Transmission {

	t, err := transmissionrpc.New(host, user, pswd,
		&transmissionrpc.AdvancedConfig{
			Port: port,
		})
	if err != nil {
		log.Fatal(err)
	}
	return &Transmission{t}
}

// Add add a new magnet link to the transmission server
func (t *Transmission) Add(magnet string) error {
	_, err := t.client.TorrentAdd(context.TODO(), transmissionrpc.TorrentAddPayload{
		Filename: &magnet,
	})
	if err != nil {
		return err
	}
	return nil
}
