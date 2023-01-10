package internal

import (
	"encoding/json"
	"fmt"
	"matcher/pkg"
	"net/http"
)

//go:generate mockgen -source=repository.go -destination=./../test/mock/mock_repository.go -package=mock
type IRepository interface {
	GetNearestDriver(port int, lat, long float64) (*Location, error)
}

type Repository struct {
	Client *http.Client
}

func NewRepository() IRepository {
	return &Repository{
		Client: &http.Client{},
	}
}

// Send a http request to get nearest driver
func (r *Repository) GetNearestDriver(port int, lat, long float64) (*Location, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(fmt.Sprintf("http://localhost:%d/api/v1/drivers/nearest?lat=%f&long=%f", port, lat, long)), nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, pkg.ErrDriverNotFound
	}
	var location Location

	json.NewDecoder(resp.Body).Decode(&location)

	defer resp.Body.Close()

	return &location, nil
}
