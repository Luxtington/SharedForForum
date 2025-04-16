package http

import (
	"encoding/json"
	"fmt"
	"github.com/Luxtington/SharedForForum/models"
	"net/http"
)

type Client struct {
	BaseURL string
}

func (c *Client) GetUserByID(userID int) (*models.User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/users/%d", c.BaseURL, userID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
