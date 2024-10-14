package services

import (
	"MovieRest/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
	"net/url"
)

func (s *MoviesService) SearchInfoForMovie(title string) (*models.Movie, error) {
	const op = "services.SearchInfoForMovie"

	s.l.With(slog.String("op", op))

	apiKey := "f4d65d0a865879453d9c7421ae7a67bd"
	baseURL := "https://api.themoviedb.org/3/search/movie"

	queryTitle := url.QueryEscape(title)
	requestURL := fmt.Sprintf("%s?api_key=%s&query=%s", baseURL, apiKey, queryTitle)

	resp, err := http.Get(requestURL)
	if err != nil {
		s.l.Error("Ошибка при запросе к TMDb API", slog.Any("error", err))
		return nil, fmt.Errorf("ошибка при запросе к TMDb API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.l.Error("Неверный ответ от TMDb API", slog.Int("status", resp.StatusCode))
		return nil, fmt.Errorf("неверный ответ от TMDb API: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.l.Error("Ошибка при чтении ответа от TMDb API", slog.Any("error", err))
		return nil, fmt.Errorf("ошибка при чтении ответа от TMDb API: %w", err)
	}

	var searchResults struct {
		Results []struct {
			ID          int     `json:"id"`
			Title       string  `json:"title"`
			Overview    string  `json:"overview"`
			Country     string  `json:"country"`
			ReleaseDate string  `json:"release_date"`
			PosterPath  string  `json:"poster_path"`
			VoteAverage float32 `json:"vote_average"`
		} `json:"results"`
	}

	if err := json.Unmarshal(body, &searchResults); err != nil {
		s.l.Error("Ошибка при парсинге ответа от TMDb API", slog.Any("error", err))
		return nil, fmt.Errorf("ошибка при парсинге ответа от TMDb API: %w", err)
	}

	if len(searchResults.Results) == 0 {
		s.l.Error("Фильм не найден", slog.String("title", title))
		return nil, fmt.Errorf("фильм с названием '%s' не найден", title)
	}

	tmdbMovie := searchResults.Results[0]

	movie := &models.Movie{
		Title:       tmdbMovie.Title,
		Description: tmdbMovie.Overview,
		Year:        parseYear(tmdbMovie.ReleaseDate),
		PosterURL:   constructPosterURL(tmdbMovie.PosterPath),
		Rating:      tmdbMovie.VoteAverage,
		Country:     tmdbMovie.Country,
	}

	return movie, nil
}

func parseYear(releaseDate string) int32 {
	if len(releaseDate) >= 4 {
		var year int
		fmt.Sscanf(releaseDate, "%4d", &year)
		return int32(year)
	}
	return 0
}

func constructPosterURL(posterPath string) string {
	if posterPath != "" {
		return "https://image.tmdb.org/t/p/w500" + posterPath
	}
	return ""
}
