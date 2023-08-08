package storage

import (
	"context"
	"fmt"
	"github.com/shawkyelshalawy/Daily_Brief/model"
	"math/rand"
)

func (d *Database) SignupForNewsletter(ctx context.Context, email model.Email) (string, error) {
	token, err := createSecret()
	if err != nil {
		return "", err
	}
	query := `
		insert into newsletter_subscribers (email, token)
		values ($1, $2)
		on conflict (email) do update set
			token = excluded.token,
			updated = now()`
	_, err = d.DB.ExecContext(ctx, query, email, token)
	return token, err
}

func createSecret() (string, error) {
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", secret), nil
}
