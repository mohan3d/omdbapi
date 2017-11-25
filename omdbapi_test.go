package omdbapi

import (
	"os"
	"testing"
)

const ()

var titles = []string{
	"The Shawshank Redemption",
	"The Godfather",
	"The Dark Knight",
	"The Good, the Bad and the Ugly",
}

var IDs = []string{
	"tt0111161",
	"tt0068646",
	"tt0468569",
	"tt0060196",
}

func APIKey() string {
	return os.Getenv("OMDBAPI_KEY")
}

func TestInvalidAPIKey(t *testing.T) {
	client := New("INVALID_API_KEY")
	_, err := client.Title("MOVIE_TITLE")
	if err == nil || err.Error() != invalidAPIKeyError {
		t.Errorf("expected %v error got %v", invalidAPIKeyError, err)
	}
}

func TestValidMovieTitle(t *testing.T) {
	client := New(APIKey())
	for _, title := range titles {
		_, err := client.Title(title)
		if err != nil {
			t.Errorf("expected no errors with title '%v' got %v", title, err)
		}
	}
}

func TestInvalidMovieTitle(t *testing.T) {
	client := New(APIKey())
	_, err := client.Title("INVALID_MOVIE_TITLE")
	if err == nil || err.Error() != movieNotFoundError {
		t.Errorf("expected %v error got %v", movieNotFoundError, err)
	}
}

func TestValidMovieId(t *testing.T) {
	client := New(APIKey())
	for _, movieID := range IDs {
		_, err := client.ID(movieID)
		if err != nil {
			t.Errorf("expected no errors with id '%v' got %v", movieID, err)
		}
	}
}

func TestInvalidMovieId(t *testing.T) {
	client := New(APIKey())
	_, err := client.ID("INVALID_MOVIE_ID")
	if err == nil || err.Error() != incorrectIMDbIDError {
		t.Errorf("expected %v error got `%v`", incorrectIMDbIDError, err)
	}
}

func TestSearchValidMovie(t *testing.T) {
	client := New(APIKey())
	_, err := client.Search("Shawshank")
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
}

func TestSearchInvalidMovie(t *testing.T) {
	client := New(APIKey())
	_, err := client.Search("INVALID_MOVIE_TITLE")
	if err == nil || err.Error() != movieNotFoundError {
		t.Errorf("expected %v error got %v", movieNotFoundError, err)
	}
}
func TestPosterValidMovieID(t *testing.T) {
	client := New(APIKey())
	for _, movieID := range IDs {
		_, err := client.Poster(movieID)
		if err != nil {
			t.Errorf("expected no errors with id '%v' got %v", movieID, err)
		}
	}
}
func TestPosterInvalidMovieID(t *testing.T) {
	client := New(APIKey())
	_, err := client.Poster("INVALID_MOVIE_ID")
	if err == nil || err.Error() != incorrectIMDbIDError {
		t.Errorf("expected no errors with id '%v' got %v", incorrectIMDbIDError, err)
	}
}
