package uuid

import (
	uuid "github.com/satori/go.uuid"
)

func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}

type UUIDGenerator struct {
}

func (g *UUIDGenerator) Generate() uuid.UUID {
	return uuid.NewV4()
}
