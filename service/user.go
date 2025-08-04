package service

import (
	"context"
	"ent/ent"
	"ent/ent/user"
)

// CreateUser tạo mới một user
func CreateUser(ctx context.Context, client *ent.Client, username, email, passwordHash, fullName string) (*ent.User, error) {
	return client.User.
		Create().
		SetUsername(username).
		SetEmail(email).
		SetPasswordHash(passwordHash).
		SetFullName(fullName).
		Save(ctx)
}

// GetUserByID đọc user theo ID
func GetUserByID(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {
	return client.User.
		Query().
		Where(user.IDEQ(id)).
		Only(ctx)
}

// GetUserByUsername đọc user theo username
func GetUserByUsername(ctx context.Context, client *ent.Client, username string) (*ent.User, error) {
	return client.User.
		Query().
		Where(user.UsernameEQ(username)).
		Only(ctx)
}

// GetUserByEmail đọc user theo email
func GetUserByEmail(ctx context.Context, client *ent.Client, email string) (*ent.User, error) {
	return client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
}
