package v1

import (
	"crypto/x509"
	"github.com/gin-cli/init/jwk"
	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2"
)

func TestPublicKey(c *gin.Context) {
	publicKey := jose.JSONWebKey{
		Key:          jwk.PublicKey,
		Certificates: []*x509.Certificate{},
		KeyID:        "melody",
		Algorithm:    "RS256",
		Use:          "sig",
	}
	c.JSON(200, publicKey)
}

func TestPrivateKey(c *gin.Context) {
	privateKey := jose.JSONWebKey{
		Key:          jwk.PrivateKey,
		Certificates: []*x509.Certificate{},
		KeyID:        "melody",
		Algorithm:    "RS256",
		Use:          "sig",
	}
	c.JSON(200, privateKey)
}
