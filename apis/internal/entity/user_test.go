package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Lucas Coronel", "lc@gmail.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Lucas Coronel", user.Name)
	assert.Equal(t, "lc@gmail.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Lucas Coronel", "lc@gmail.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
}
