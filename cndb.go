// Package chucknorris is a golang implementation of the ICNDB API
package chucknorris

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	baseURL = "http://api.icndb.com"
)

// Joke struct represents the Joke returned by ICNDB API
type Joke struct {
	ID   int    `json:"id"`
	Joke string `json:"joke"`
}

// Stringer interface for Joke struct
func (j *Joke) String() string {
	return j.Joke
}

// CNDBResponse is the type for API Result
type CNDBResponse struct {
	Type  string `json:"type"`
	Value *Joke  `json:"value"`
}

// CNDBRequest is used to specify request parameters for the API
type CNDBRequest struct {
	FirstName string
	LastName  string
}

// URLValues translates properities of Request struct to URL Parameters for API
func (r *CNDBRequest) URLValues() url.Values {
	params := url.Values{}
	if len(r.FirstName) > 0 {
		params.Set("firstName", r.FirstName)
	}

	if len(r.LastName) > 0 {
		params.Set("lastName", r.LastName)
	}
	return params
}

// GetRandomJoke returns a random joke from ICNDB API
func GetRandomJoke(req *CNDBRequest) (*Joke, error) {
	if req == nil {
		req = &CNDBRequest{}
	}
	resp, err := request("jokes/random", req.URLValues())
	if err != nil {
		return nil, err
	}
	return resp.Value, nil
}

// GetJokeByID can be used to get a joke specified by its ID
func GetJokeByID(id int, req *CNDBRequest) (*Joke, error) {
	if req == nil {
		req = &CNDBRequest{}
	}
	resp, err := request(fmt.Sprintf("jokes/%d", id), req.URLValues())
	if err != nil {
		return nil, err
	}
	return resp.Value, nil
}

// request is a helper method used to form the URL and make the actual request
func request(op string, p url.Values) (*CNDBResponse, error) {
	url := fmt.Sprintf("%s/%s?%s", baseURL, op, p.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var v CNDBResponse
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
