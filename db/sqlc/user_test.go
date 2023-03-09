package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zahrou/ecommerce/util"
)

func CreateRandomUser(t *testing.T) User {
	args := CreateUserParams{
		Username:    util.RnadomName(),
		Email:       util.RandomEmail(),
		Password:    util.RnadomName(),
		PhoneNumber: util.RandomInteger(9, 100),
		ZipCode:     util.RandomInteger(9, 20),
	}

	user, err := testQueries.CreateUser(context.Background(), args)

	fmt.Println(user.PhoneNumber)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.Email, user.Email)
	require.Equal(t, args.Password, user.Password)
	require.Equal(t, args.PhoneNumber, user.PhoneNumber)

	require.NotZero(t, user.ID)

	return user

}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestDeleteUserUser(t *testing.T) {
	user := CreateRandomUser(t)

	err := testQueries.DeleteUsers(context.Background(), user.ID)

	require.NoError(t, err)

	user2, err := testQueries.GetUsers(context.Background(), user.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)

}

func TestUpdatUser(t *testing.T) {
	user1 := CreateRandomUser(t)

	args := UpdateUserParams{
		ID:       user1.ID,
		Username: util.RnadomName(),
	}

	user2, err := testQueries.UpdateUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, args.Username, user2.Username)
	require.Equal(t, args.ID, user2.ID)

	require.NotZero(t, user2)
}

func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t)

	user2, err := testQueries.GetUsers(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.Password, user2.Password)
	require.Equal(t, user.ID, user2.ID)

}

func TestListUsers(t *testing.T) {
	for i := 0; i < 7; i++ {
		CreateRandomUser(t)
	}

	args := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
