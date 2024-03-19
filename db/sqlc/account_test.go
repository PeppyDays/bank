package db

import (
	"context"
	"strconv"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
)

func getRandomCreateAccountParams() CreateAccountParams {
	owner := faker.LastName()
	amountWithCurrency := strings.Split(faker.AmountWithCurrency(), " ")
	currency := amountWithCurrency[0]
	balance, _ := strconv.ParseInt(amountWithCurrency[1], 10, 64)

	params := CreateAccountParams{
		Owner:    owner,
		Balance:  balance,
		Currency: currency,
	}

	return params
}

func getRandomUpdateAccountParams(id int64) UpdateAccountParams {
	amountWithCurrency := strings.Split(faker.AmountWithCurrency(), " ")
	balance, _ := strconv.ParseInt(amountWithCurrency[1], 10, 64)

	params := UpdateAccountParams{
		ID:      id,
		Balance: balance,
	}

	return params
}

func arrangeAccount() Account {
	createAccountParams := getRandomCreateAccountParams()
	account, _ := queries.CreateAccount(context.Background(), createAccountParams)
	return account
}

func TestCreateAccount(t *testing.T) {
	// given
	params := getRandomCreateAccountParams()

	// when
	account, err := queries.CreateAccount(context.Background(), params)

	// then
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, params.Owner, account.Owner)
	require.Equal(t, params.Balance, account.Balance)
	require.Equal(t, params.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	// given
	account := arrangeAccount()
	params := getRandomUpdateAccountParams(account.ID)

	// when
	updatedAccount, err := queries.UpdateAccount(context.Background(), params)

	// then
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, params.Balance, updatedAccount.Balance)
	require.Equal(t, account.Currency, updatedAccount.Currency)
}

func TestGetAccount(t *testing.T) {
	// given
	account := arrangeAccount()

	// when
	acquiredAccount, err := queries.GetAccount(context.Background(), account.ID)

	// then
	require.NoError(t, err)
	require.NotEmpty(t, acquiredAccount)
	require.Equal(t, account, acquiredAccount)
}

func TestDeleteAccount(t *testing.T) {
	// given
	account := arrangeAccount()

	// when
	err := queries.DeleteAccount(context.Background(), account.ID)

	// then
	require.NoError(t, err)

	// when
	acquiredAccount, err := queries.GetAccount(context.Background(), account.ID)

	// then
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, acquiredAccount)
}
