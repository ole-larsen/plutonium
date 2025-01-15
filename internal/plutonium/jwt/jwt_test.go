package jwt_test

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/ole-larsen/plutonium/internal/plutonium/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testSecret = "testsecret"
	testLogin  = "testuser"
)

func TestCreateJWTToken(t *testing.T) {
	token, err := jwt.CreateToken(testLogin, testSecret)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Verify token structure and claims
	claims, err := jwt.Verify(token, testSecret)
	require.NoError(t, err)
	require.Equal(t, testLogin, claims["login"])
	require.NotEmpty(t, claims["timestamp"])
}

func TestVerifyJwt_Success(t *testing.T) {
	token, err := jwt.CreateToken(testLogin, testSecret)
	require.NoError(t, err)

	claims, err := jwt.Verify(token, testSecret)
	require.NoError(t, err)
	require.Equal(t, testLogin, claims["login"])
	require.NotEmpty(t, claims["timestamp"])
}

func TestVerifyJwt_InvalidToken(t *testing.T) {
	_, err := jwt.Verify("invalidtoken", testSecret)
	require.Error(t, err)
	require.Contains(t, err.Error(), "token contains an invalid number of segments")
}

func TestVerifyJwt_InvalidSecret(t *testing.T) {
	token, err := jwt.CreateToken(testLogin, testSecret)
	require.NoError(t, err)

	_, err = jwt.Verify(token, "wrongsecret")
	require.Error(t, err)
	require.Contains(t, err.Error(), "signature is invalid")
}

func TestVerifyJwt_InvalidSigningMethod(t *testing.T) {
	// Create a token with a different signing method
	claims := jwtgo.MapClaims{
		"login":     testLogin,
		"timestamp": time.Now(),
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
	_, err := token.SignedString([]byte(testSecret))
	require.Error(t, err)
	require.Contains(t, err.Error(), "key is invalid")
}

func TestSignJwt(t *testing.T) {
	claims := jwtgo.MapClaims{
		"login": testLogin,
	}
	token, err := jwt.Sign(claims, testSecret)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	parsedToken, err := jwtgo.Parse(token, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, jwtgo.ErrInvalidKey
		}

		return []byte(testSecret), nil
	})
	require.NoError(t, err)

	if err != nil {
		t.Fatal("there must be no errors")
	}

	require.True(t, parsedToken.Valid)

	parsedClaims, ok := parsedToken.Claims.(jwtgo.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, testLogin, parsedClaims["login"])
}

func TestCreateJWTToken_EmptySecret(t *testing.T) {
	_, err := jwt.CreateToken(testLogin, "")
	require.Error(t, err)
	require.Contains(t, err.Error(), "could not sign token") // Adjust based on actual error message
}

func TestVerifyJwt_EmptyToken(t *testing.T) {
	_, err := jwt.Verify("", testSecret)
	require.Error(t, err)
	require.Contains(t, err.Error(), "token contains an invalid number of segments")
}

func TestVerifyJwt(t *testing.T) {
	token, err := jwt.CreateToken(testLogin, testSecret)
	require.NoError(t, err)

	claims, err := jwt.Verify(token, testSecret)

	require.NoError(t, err)
	require.Equal(t, testLogin, claims["login"])

	_, err = jwt.Verify(token, "")

	require.Error(t, err)

	_, err = jwt.Verify("balblabla", testSecret)

	require.Error(t, err)

	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6InRlc3R1c2VyIiwidGltZXN0YW1wIjoiMjAyNC0wOC0zMFQyMDo0MjoyMS45MTU0NTI4NloifQ.5Qu-2qtl2gozvq-Axkuy3ChbXacQvT_CuMu18sQSF25"
	_, err = jwt.Verify(token, testSecret)

	require.Error(t, err)
	require.Equal(t, "signature is invalid", err.Error())

	claimsMap := jwtgo.MapClaims{
		"login":     testLogin,
		"timestamp": time.Now(),
	}
	// jwt.SigningMethodHS256
	tkn := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claimsMap)
	token, err = tkn.SignedString([]byte(testSecret))

	require.NoError(t, err)

	_, err = jwt.Verify(token, testSecret)

	require.NoError(t, err)

	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwtgo.StandardClaims
	}

	// Create the Claims
	customClaims := MyCustomClaims{
		"bar",
		jwtgo.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	customToken := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, customClaims)
	ss, err := customToken.SignedString(mySigningKey)
	require.NoError(t, err)
	_, err = jwt.Verify(ss, testSecret)
	require.Error(t, err)
	require.Equal(t, "signature is invalid", err.Error())
}

func TestCreateJWTToken_EmptyLogin(t *testing.T) {
	token, err := jwt.CreateToken("", testSecret)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Verify token structure and claims
	claims, err := jwt.Verify(token, testSecret)
	require.NoError(t, err)
	require.Empty(t, claims["login"]) // Expect empty login in claims
	require.NotEmpty(t, claims["timestamp"])
}

func TestVerifyJwt_InvalidSigningMethod2(t *testing.T) {
	// Create a token with an unsupported signing method
	claims := jwtgo.MapClaims{
		"login":     testLogin,
		"timestamp": time.Now(),
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodNone, claims)
	signedToken, err := token.SignedString([]byte(testSecret))
	require.Error(t, err)

	_, err = jwt.Verify(signedToken, testSecret)
	require.Error(t, err)
	require.Contains(t, err.Error(), "token contains an invalid number of segment")
}

func TestVerifyJwt_InvalidTokenFormat(t *testing.T) {
	_, err := jwt.Verify("not.a.valid.token.format", testSecret)
	require.Error(t, err)
	require.Contains(t, err.Error(), "token contains an invalid number of segments")
}

func TestVerifyJwt_ExpiredToken(t *testing.T) {
	// Create a token with an expiration time in the past
	claims := jwtgo.MapClaims{
		"login":     testLogin,
		"timestamp": time.Now(),
		"exp":       time.Now().Add(-time.Hour).Unix(), // Expired one hour ago
	}
	token, err := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims).SignedString([]byte(testSecret))
	require.NoError(t, err)

	_, err = jwt.Verify(token, testSecret)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Token is expired")
}
func TestVerifyJwt_InvalidSigningMethod3(t *testing.T) {
	// Generate an RSA private key for testing
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Create a token with a different signing method (e.g., SigningMethodRS256)
	claims := jwtgo.MapClaims{
		"login":     testLogin,
		"timestamp": time.Now(),
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	require.NoError(t, err)

	// VerifyJwt should fail due to invalid signing method
	_, err = jwt.Verify(signedToken, testSecret)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid token")
}

func TestGetBearerToken_ValidHeader(t *testing.T) {
	// Test valid Bearer token
	header := "Bearer validToken123"
	token, err := jwt.GetBearerToken(header)
	require.NoError(t, err)
	assert.Equal(t, "validToken123", token)
}

func TestGetBearerToken_MalformedHeader(t *testing.T) {
	// Test malformed header without "Bearer" keyword
	header := "InvalidHeader validToken123"
	token, err := jwt.GetBearerToken(header)
	require.NoError(t, err)
	assert.NotEmpty(t, token)

	// Test malformed header with too many parts
	header = "Bearer token extraPart"
	token, err = jwt.GetBearerToken(header)
	require.Error(t, err)
	assert.Empty(t, token)
	assert.Contains(t, err.Error(), "malformed bearer token")
}

func TestGetBearerToken_EmptyHeader(t *testing.T) {
	// Test empty header
	header := ""
	token, err := jwt.GetBearerToken(header)
	require.Error(t, err)
	assert.Empty(t, token)
	assert.Contains(t, err.Error(), "an authorization header is required")
}
