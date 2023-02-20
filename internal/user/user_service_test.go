package user

import (
	"errors"
	"gin-starter/pkg/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserService(t *testing.T) {
	dao := newMockUserDAO()
	s := NewUserService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestUserService_Get(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.Get(2)
	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "Ben", user.FirstName)
		assert.Equal(t, "Doe", user.LastName)
	}

	user, err = s.Get(100)
	assert.NotNil(t, err)
}

func newMockUserDAO() userDAO {
	return &mockUserDAO{
		records: []User{
			{Model: common.Model{ID: 1}, FirstName: "John", LastName: "Smith", Email: "john.smith@gmail.com", Address: "Dummy Value"},
			{Model: common.Model{ID: 2}, FirstName: "Ben", LastName: "Doe", Email: "ben.doe@gmail.com", Address: "Dummy Value"},
		},
	}
}

func (m *mockUserDAO) Get(id uint) (*User, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

type mockUserDAO struct {
	records []User
}