package db

import (
	"context"
	"testing"

	"github.com/sonymimic1/go-transfer/pkg/util"
	"github.com/stretchr/testify/require"
)

func addUser(t *testing.T) User {

	hashedPassword, err := util.HashedPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	var user User

	_, err = TestQuri.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	user, err = TestQuri.GetUser(context.Background(), arg.Username)
	require.NoError(t, err)

	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangeAt.Valid)

	return user
}
func TestQueries_CreateUser(t *testing.T) {
	addUser(t)
}
