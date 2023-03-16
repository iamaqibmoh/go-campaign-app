package test

import (
	"bwa-campaign-app/app"
	"bwa-campaign-app/model/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBConnection(t *testing.T) {
	var users []domain.User

	db := app.DBConnection()
	db.Find(&users)

	var emails []string
	for _, user := range users {
		emails = append(emails, user.Email)
	}

	assert.Equal(t, "ucup@test.com", emails[1])
	assert.Equal(t, "aqib@test.com", emails[0])
}
