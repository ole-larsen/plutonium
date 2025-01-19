package oauth2client

import (
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"strings"

	"github.com/ole-larsen/plutonium/internal/plutonium/settings"
	"golang.org/x/oauth2"
)

type Oauth2 struct {
	Client *Client
	Config map[string]oauth2.Config
}
type Client struct {
	settings *settings.Settings
}

func genCodeChallengeS256(s string) string {
	s256 := sha256.Sum256([]byte(s))
	return base64.URLEncoding.EncodeToString(s256[:])
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SetSettings(cfg *settings.Settings) *Client {
	c.settings = cfg
	return c
}

func (c *Client) Config(clientID, clientSecret string) oauth2.Config {
	return oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"all"},
		RedirectURL:  c.settings.OAUTH2.Callback,
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.settings.OAUTH2.Provider + "/api/v1/authorize",
			TokenURL: c.settings.OAUTH2.Provider + "/api/v1/token",
		},
	}
}

func (c *Client) AuthorizeURL(config *oauth2.Config) string {
	return config.AuthCodeURL("xyz",
		oauth2.SetAuthURLParam("code_challenge", genCodeChallengeS256("s256example")),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("grant_type", "client_credentials"))
}

func (c *Client) GetClientIDFromReferer(refererURL string) (string, error) {
	clientID := ""

	refererParameters := strings.Split(refererURL, "?")
	if len(refererParameters) > 1 {
		for _, parameter := range strings.Split(refererParameters[1], "&") {
			if strings.Contains(parameter, "client_id") {
				clientID = strings.Split(parameter, "=")[1]

				decodedValue, err := url.QueryUnescape(clientID)
				if err != nil {
					return "", err
				}

				clientID = decodedValue
			}
		}
	}

	return clientID, nil
}
