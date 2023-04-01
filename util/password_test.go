package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashhedpassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashhedpassword)

	err = CheckPassword(password, hashhedpassword)

	require.NoError(t, err)

	worngpassword := RandomString(6)

	err2 := CheckPassword(worngpassword, hashhedpassword)
	require.EqualError(t, err2, bcrypt.ErrMismatchedHashAndPassword.Error())

}
