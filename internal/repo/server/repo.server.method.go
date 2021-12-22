package server

import (
	"fmt"

	ServerEntity "github.com/TheUnderdogFolks/budgetme-backend/internal/entity/server"
	"github.com/TheUnderdogFolks/budgetme-backend/internal/pkg/env"
)

func (r *Repo) GetHealth() (*ServerEntity.Health, error) {
	return &ServerEntity.Health{Message: fmt.Sprintf("Server running on %s", env.Get("version", "@latest"))}, nil
}