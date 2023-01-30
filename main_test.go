package main

import (
	"sync"
	"testing"
)

func TestUpdateMessage(t *testing.T) {
	msg = "Hello World"
	updatedMessage := "Goodbye, cruel world!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("x", &mutex)
	go updateMessage(updatedMessage, &mutex)
	wg.Wait()

	if msg != updatedMessage {
		t.Errorf("Expected %q but got %q", updatedMessage, msg)
	}
}
