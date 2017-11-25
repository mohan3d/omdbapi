package omdbapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	movieAPI    = "https://www.omdbapi.com"
	posterAPI   = "https://img.omdbapi.com"
	searchParam = "s"
	titleParam  = "t"
	idParam     = "i"
)

// MovieInfo describes movie info.
type MovieInfo struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	DVD        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}

// SearchInfo describes search query results.
type SearchInfo struct {
	Search []struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	} `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response     string `json:"Response"`
}

// APIParam describes api query parameter used in client methods.
type APIParam struct {
	Name  string
	Value string
}

// Poster describes poster from omdb.
type Poster []byte

// Client describes omdb API client.
type Client struct {
	apiKey string
}

// Title finds movie info by movie title.
func (c *Client) Title(title string, params ...APIParam) (*MovieInfo, error) {
	params = append(params, APIParam{Name: titleParam, Value: title})
	return c.find(params...)
}

// ID finds movie info by imdb-movie-id.
func (c *Client) ID(id string, params ...APIParam) (*MovieInfo, error) {
	params = append(params, APIParam{Name: idParam, Value: id})
	return c.find(params...)
}

// Search searches for a movie by title.
func (c *Client) Search(title string, params ...APIParam) (*SearchInfo, error) {
	params = append(params, APIParam{Name: searchParam, Value: title})
	data, err := c.get(movieAPI, params...)

	if err != nil {
		return nil, err
	}
	var searchInfo SearchInfo

	if err := json.Unmarshal(data, &searchInfo); err != nil {
		return nil, err
	}
	return &searchInfo, nil
}

// PosterByTitle gets poster by movie title.
func (c *Client) PosterByTitle(title string) (Poster, error) {
	return c.poster(APIParam{Name: titleParam, Value: title})
}

// PosterByID gets movie poster by imdb-movie-id.
func (c *Client) PosterByID(id string) (Poster, error) {
	return c.poster(APIParam{Name: idParam, Value: id})
}

func (c *Client) poster(param APIParam) (Poster, error) {
	data, err := c.get(posterAPI, param)
	if err != nil {
		return nil, err
	}
	return Poster(data), nil
}

func (c *Client) find(params ...APIParam) (*MovieInfo, error) {
	data, err := c.get(movieAPI, params...)
	if err != nil {
		return nil, err
	}
	var movieInfo MovieInfo

	if err := json.Unmarshal(data, &movieInfo); err != nil {
		return nil, err
	}
	return &movieInfo, nil
}

func (c *Client) get(apiURL string, params ...APIParam) ([]byte, error) {
	URL, err := url.Parse(apiURL)

	if err != nil {
		return nil, err
	}
	query := URL.Query()
	query.Set("apikey", c.apiKey)

	for _, param := range params {
		query.Set(param.Name, param.Value)
	}
	URL.RawQuery = query.Encode()
	response, err := http.Get(URL.String())

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	if err := reponseError(body); err != nil {
		return nil, err
	}
	return body, nil
}

func reponseError(response []byte) error {
	var errorResponse struct {
		Response string `json:"Response"`
		Error    string `json:"Error"`
	}
	if err := json.Unmarshal(response, &errorResponse); err != nil {
		return err
	}
	if errorResponse.Error != "" {
		return errors.New(errorResponse.Error)
	}
	return nil
}

// New returns a new client reference.
func New(apiKey string) *Client {
	client := new(Client)
	client.apiKey = apiKey
	return client
}
