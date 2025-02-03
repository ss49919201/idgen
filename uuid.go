package main

import (
	"fmt"

	"github.com/google/uuid"
)

func generateUUID() uuid.UUID {
	return uuid.New()
}

func generateUUIDV7() (uuid.UUID, error) {
	var id uuid.UUID
	id, err := uuid.NewV7()
	if err != nil {
		return id, fmt.Errorf("failed to generate uuid v7: %w", err)
	}

	return id, nil
}
