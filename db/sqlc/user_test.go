package db

import (
	"context"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	acc1 := createRandomUser(t)
	acc2, err := testQueries.GetUser(context.Background(), acc1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc1.Username, acc2.Username)
	require.Equal(t, acc1.FullName, acc2.FullName)
	require.Equal(t, acc1.HashedPassword, acc2.HashedPassword)
	require.Equal(t, acc1.Email, acc2.Email)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
	require.WithinDuration(t, acc1.PasswordChangedAt, acc2.PasswordChangedAt, time.Second)
}

// func TestUpdateAccount(t *testing.T) {
// 	acc := createRandomAccount(t)
//
// 	arg := UpdateAccountParams{
// 		ID:      acc.ID,
// 		Balance: faker.UnixTime(),
// 	}
//
// 	acc2, err := testQueries.UpdateAccount(context.Background(), arg)
// 	require.NoError(t, err)
//
// 	require.NoError(t, err)
// 	require.NotEmpty(t, acc2)
// 	require.Equal(t, acc.ID, acc2.ID)
// 	require.Equal(t, acc.Owner, acc2.Owner)
// 	require.Equal(t, acc2.Balance, arg.Balance)
// 	require.Equal(t, acc.Currency, acc2.Currency)
// }
//
// func TestDeleteAccount(t *testing.T) {
// 	acc := createRandomAccount(t)
// 	err := testQueries.DeleteAccount(context.Background(), acc.ID)
// 	require.NoError(t, err)
//
// 	acc2, err := testQueries.GetAccount(context.Background(), acc.ID)
// 	require.Error(t, err)
// 	require.Empty(t, acc2)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// }
//
// func TestListAccounts(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomAccount(t)
// 	}
//
// 	accounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{
// 		Limit:  5,
// 		Offset: 5,
// 	})
//
// 	require.NoError(t, err)
//
// 	for _, acc := range accounts {
// 		require.NotEmpty(t, acc)
// 	}
// }

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       faker.Name(),
		FullName:       faker.FirstName() + " " + faker.LastName(),
		Email:          faker.Email(),
		HashedPassword: faker.Password(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}
