package queue

import "errors"

var (
	ErrQueueUnderflow = errors.New("Queue underflow")
	ErrPopFailure     = errors.New("could not pop")
)
