package oauth

import (
    "context"
    "errors"
    "net/http"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var (
    googleOAuthConfig = &oauth2.Config{
        ClientID:     "YOUR_CLIENT_ID",
        ClientSecret: "YOUR_CLIENT_SECRET",
        RedirectURL:  "YOUR_REDIRECT_URL",
        Scopes: []string{
            "email",
            "profile",
        },
        Endpoint: google.Endpoint,
    }
)

// GetAuthURL generates the URL for the OAuth2 authentication request.
func GetAuthURL() string {
    return googleOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

// ExchangeCode exchanges the authorization code for an access token.
func ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
    if code == "" {
        return nil, errors.New("authorization code is required")
    }
    token, err := googleOAuthConfig.Exchange(ctx, code)
    if err != nil {
        return nil, err
    }
    return token, nil
}

// GetUserInfo retrieves user information from the OAuth2 provider.
func GetUserInfo(ctx context.Context, token *oauth2.Token) (map[string]interface{}, error) {
    client := googleOAuthConfig.Client(ctx, token)
    resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("failed to get user info")
    }

    var userInfo map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        return nil, err
    }
    return userInfo, nil
}