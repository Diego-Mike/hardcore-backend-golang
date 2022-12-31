package db

import (
	"context"
	"testing"
	"time"

	"github.com/Diego-Mike/hardcore-backend-golang/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	// basic, check if no error | empty object
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// check if data is equal to arguments provided
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	// date and id generated correctly ?
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {

	// create account
	user1 := createRandomUser(t)

	getUser, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, getUser)

	require.Equal(t, user1.Username, getUser.Username)
	require.Equal(t, user1.HashedPassword, getUser.HashedPassword)
	require.Equal(t, user1.FullName, getUser.FullName)
	require.Equal(t, user1.Email, getUser.Email)
	require.WithinDuration(t, user1.CreatedAt, getUser.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, getUser.PasswordChangedAt, time.Second)

}
