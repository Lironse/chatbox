package main

import "errors"

var (
	ErrUpgradeFailure       = errors.New("upgrade to ws connection failed")
	ErrReadPacketFailure    = errors.New("packet could not be read")
	ErrPacketMarshalFailure = errors.New("packet could not be marshalled into a string")
	ErrSendPacketFailure    = errors.New("packet could not be sent")
	ErrClientNotFound       = errors.New("a client was not found")
	ErrClientAlreadyExists  = errors.New("username is already taken")
)
