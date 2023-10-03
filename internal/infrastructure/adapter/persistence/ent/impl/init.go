package impl

import (
	"context"
	"fmt"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/tools/ent"

	_ "github.com/lib/pq"
)

func NewEntImplementation(dbHost, dbPort, dbUser, dbPassword, dbName string) (repository.UserRepository, error) {
	// Connect to the database
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword))
	if err != nil {
		return nil, codeerr.NewInternal("NewEntImplementation", err)
	}

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, codeerr.NewInternal("NewEntImplementation", err)
	}

	return NewUserRepository(client), nil
}
