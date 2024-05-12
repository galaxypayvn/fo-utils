package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type Config struct {
	GoogleAppCreds    []byte
	FirebaseProjectID string
}

func newApp(cfg Config) (*firebase.App, error) {
	opts := option.WithCredentialsJSON(cfg.GoogleAppCreds)
	config := &firebase.Config{
		ProjectID: cfg.FirebaseProjectID,
	}

	app, err := firebase.NewApp(context.Background(), config, opts)

	return app, err
}
