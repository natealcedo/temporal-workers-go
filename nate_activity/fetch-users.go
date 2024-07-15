package nate_activity

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	} `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

func (a *GreetActivitiesImpl) FetchUsers(ctx context.Context) ([]User, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return users, nil
}
