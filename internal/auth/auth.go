package auth

import "context"

type Provider interface {
	Login(ctx context.Context) error
	Verify(ctx context.Context) error
}
