package services

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestJWTService_GenerateToken(t *testing.T) {
	err := os.Setenv("secret", "gnwnoigbiqgj")
	assert.Nil(t, err)

	jwtsvc := JWTService{}
	kwt, err := jwtsvc.GenerateToken("test username")
	assert.Nil(t, err)
	assert.Greater(t, len(kwt), 0)
}
