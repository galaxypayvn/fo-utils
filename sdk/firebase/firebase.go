package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type Config struct {
	GoogleAppCredFilepath string
	FirebaseProjectID     string
}

func newApp(cfg Config) (*firebase.App, error) {
	opts := option.WithCredentialsFile(cfg.GoogleAppCredFilepath)
	config := &firebase.Config{
		ProjectID: cfg.FirebaseProjectID,
	}

	app, err := firebase.NewApp(context.Background(), config, opts)

	return app, err
}
