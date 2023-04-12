package database

import (
	"testing"

	"github.com/lccoronel/golang-full-cycle/apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	database, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	database.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "john@example.com", "123")
	userDB := NewUser(database)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = database.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, user.Password)
}

func TestFindByEmail(t *testing.T) {
	database, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	database.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "john@example.com", "123")
	userDB := NewUser(database)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, user.Password)
}
