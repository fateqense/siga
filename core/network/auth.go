package network

import "net/http"

type AuthenticatedClient struct {
	client  *http.Client
	session string
}

func NewAuthenticatedClient(session string) *AuthenticatedClient {
	client := NewClient()
	return &AuthenticatedClient{client, session}
}

func (ac *AuthenticatedClient) Do(req *http.Request) (*http.Response, error) {
	sessionCookie := http.Cookie{}
	sessionCookie.Name = COOKIE_SESSION_KEY
	sessionCookie.Value = ac.session

	req.AddCookie(&sessionCookie)
	return ac.client.Do(req)
}

func (ac *AuthenticatedClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return ac.Do(req)
}
