package types

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetUser(ctx context.Context, uid uuid.UUID)
}
