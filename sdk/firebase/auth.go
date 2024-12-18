package firebase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"firebase.google.com/go/v4/auth"
)

var (
	ErrExpiredTokenCode = errors.New("firebase token has been expired")
	ErrRevokedTokenCode = errors.New("firebase token has been revoked")
)

type Auth struct {
	authClient *auth.Client
}

type UserInfo struct {
	SocialID    string
	DisplayName string `json:"displayName,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	PhotoURL    string `json:"photoUrl,omitempty"`
	ProviderID  string `json:"providerId,omitempty"`
	UID         string `json:"rawId,omitempty"`
	Metadata    []byte `json:"metadata"`
}

func NewAuthApp(cfg Config) (*Auth, error) {
	app, err := newApp(cfg)
	if err != nil {
		return nil, fmt.Errorf("initialize firebase app: %w", err)
	}

	a, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return &Auth{
		authClient: a,
	}, nil
}

func (a *Auth) VerifyToken(ctx context.Context, idToken string) (*UserInfo, error) {
	token, err := a.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "revoked"):
			return nil, fmt.Errorf("%w: %w", ErrRevokedTokenCode, err)
		case strings.Contains(err.Error(), "expired"):
			return nil, fmt.Errorf("%w: %w", ErrExpiredTokenCode, err)
		}
		return nil, fmt.Errorf("verify id token: %w", err)
	}

	userRecord, err := a.authClient.GetUser(ctx, token.UID)
	if err != nil {
		return nil, fmt.Errorf("get user from firebase: %w", err)
	}

	metadata, err := json.Marshal(userRecord)
	if err != nil {
		return nil, fmt.Errorf("marshal metadata: %w", err)
	}

	info := userRecord.ProviderUserInfo[0]
	return &UserInfo{
		SocialID:    token.UID,
		UID:         info.ProviderID,
		DisplayName: info.DisplayName,
		Email:       info.Email,
		PhoneNumber: info.PhoneNumber,
		PhotoURL:    info.PhotoURL,
		ProviderID:  info.ProviderID,
		Metadata:    metadata,
	}, nil
}
