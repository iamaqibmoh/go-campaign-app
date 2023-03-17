package test

import (
	"bwa-campaign-app/auth"
	"fmt"
	"log"
	"testing"
)

func TestJWTAuth(t *testing.T) {
	jwtAuth := auth.NewJWTAuth()
	token, err := jwtAuth.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMn0.6eadlCpRK3ySqNsrnWTHm2A7UGZ4A8_QBizBPO00Tmc")
	if err != nil {
		fmt.Println(err.Error())
	}

	if token.Valid {
		log.Println("VALID")
	} else {
		fmt.Println("INVALID")
	}
}
