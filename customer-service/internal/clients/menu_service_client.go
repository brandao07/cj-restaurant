package clients

import (
	"fmt"
	"strings"

	"github.com/brandao07/cj-restaurant/customer-service/internal/dtos"
	"github.com/go-resty/resty/v2"
)

type MenuServiceClient interface {
	GetMenus(ids []int) ([]dtos.Menu, error)
}

type menuServiceClient struct {
	client *resty.Client
}

func NewMenuServiceClient(baseURL string) MenuServiceClient {
	client := resty.New().
		SetBaseURL(baseURL).
		SetHeader("Accept", "application/json")

	return &menuServiceClient{
		client: client,
	}
}

func (c *menuServiceClient) GetMenus(ids []int) ([]dtos.Menu, error) {
	// Convert []int to comma-separated string
	var idStrings []string
	for _, id := range ids {
		idStrings = append(idStrings, fmt.Sprintf("%d", id))
	}
	idParam := strings.Join(idStrings, ",")

	// Make the GET request with query param ?ids=1,2,3
	resp, err := c.client.R().
		SetQueryParam("ids", idParam).
		SetResult(&[]dtos.Menu{}).
		Get("/menus")

	if err != nil {
		return nil, fmt.Errorf("error calling menu service: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("menu service returned status %d", resp.StatusCode())
	}

	menus := *resp.Result().(*[]dtos.Menu)
	return menus, nil
}
