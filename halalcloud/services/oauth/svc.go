package oauth

import (
	"context"
	"log"
	"time"

	"github.com/halalcloud/golang-sdk-lite/halalcloud/apiclient"
	"github.com/halalcloud/golang-sdk-lite/halalcloud/utils"
)

type OAuthService struct {
	client *apiclient.Client
}

func NewOAuthService(client *apiclient.Client) *OAuthService {
	return &OAuthService{
		client: client,
	}
}

func (s *OAuthService) TokenAuthorize() error {
	codeVerifier := utils.CreateRandomString(32)
	codeChallenge := utils.Sha256HashString(codeVerifier)
	state := utils.CreateRandomString(16)
	operationContext, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	request := &AuthorizeRequest{
		ClientId:            "123",
		ResponseType:        "code",
		RedirectUri:         "https://example.com/callback",
		Scope:               "read write",
		State:               state,
		CodeChallenge:       codeChallenge,
		CodeChallengeMethod: "S256",
	}
	response, err := s.Authorize(operationContext, request)
	if err != nil {
		return err
	}
	log.Printf("Authorization successful: Code: %s, Redirect URI: %s, State: %s, Status: %s, Code Verifier: %s",
		response.Code, response.RedirectUri, response.State, response.Status, codeVerifier)
	return nil
}

func (s *OAuthService) Authorize(ctx context.Context, req *AuthorizeRequest) (*AuthorizeResponse, error) {
	data := &AuthorizeResponse{}
	err := s.client.Post(ctx, "/v6/oauth/authorize", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *OAuthService) GetToken(ctx context.Context, req *TokenRequest) (*TokenResponse, error) {
	data := &TokenResponse{}
	err := s.client.Post(ctx, "/v6/oauth/get_token", nil, req, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
