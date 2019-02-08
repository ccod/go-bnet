package bnet

// AccountService has account-related APIs. See Client.
type AccountService struct {
	client *Client
}

// User calls the /account/user endpoint. See Battle.net docs.
func (s *AccountService) User() (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "account/user", nil)
	if err != nil {
		return nil, nil, err
	}

	var user User
	resp, err := s.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}
