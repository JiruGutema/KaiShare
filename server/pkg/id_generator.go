// Package pkg is a code which can be imported by another projects
package pkg

import "github.com/google/uuid"

func IDGenerator() (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	return id, err
}
