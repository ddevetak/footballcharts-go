// Package footballcharts provides a zero-dependency Go client for the keyless
// Football Charts REST API (https://www.football-charts.com/developers).
//
// Football Charts covers 93 leagues worldwide (including English Premier League
// through National League North/South, German Bundesliga through Regionalliga,
// Spanish LaLiga through Segunda Federación, Italian Serie A through Serie C,
// and Brasileirão Série A & B) with:
//   - Dixon-Coles match probabilities (1X2, Over/Under, BTTS, exact scorelines)
//   - Daily 10,000-run Monte Carlo league table projections
//   - Goal-minute timing distributions and historical standings
//
// Homepage: https://www.football-charts.com
// Open Data (CC-BY-4.0): https://www.football-charts.com/data
// Interactive Calculators: https://www.football-charts.com/tools/dixon-coles-calculator
package footballcharts

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// DefaultBaseURL is the public keyless REST endpoint for Football Charts.
const DefaultBaseURL = "https://www.football-charts.com/api/public/v1"

// Client is an HTTP client for the Football Charts public API.
type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// NewClient returns a new Football Charts API client configured for the keyless public endpoint.
func NewClient() *Client {
	return &Client{
		BaseURL: DefaultBaseURL,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// GetJSON performs a GET request to the given endpoint path and decodes the JSON response into v.
func (c *Client) GetJSON(ctx context.Context, path string, query url.Values, v any) error {
	u := c.BaseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "footballcharts-go/0.1.0 (+https://www.football-charts.com/developers)")
	if c.APIKey != "" {
		req.Header.Set("X-API-Key", c.APIKey)
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return fmt.Errorf("footballcharts API HTTP %d: %s", resp.StatusCode, string(body))
	}
	return json.NewDecoder(resp.Body).Decode(v)
}

// GetLeagues fetches the list of all 93 supported football leagues.
func (c *Client) GetLeagues(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	err := c.GetJSON(ctx, "/meta/leagues", nil, &out)
	return out, err
}

// GetStandings fetches league standings for a given league key (e.g. "premier", "eng1", "germany3").
func (c *Client) GetStandings(ctx context.Context, leagueKey string) (map[string]any, error) {
	q := url.Values{"league": []string{leagueKey}}
	var out map[string]any
	err := c.GetJSON(ctx, "/standings", q, &out)
	return out, err
}

// GetProjection fetches the daily 10,000-run Monte Carlo season projection for a league.
func (c *Client) GetProjection(ctx context.Context, leagueKey string) (map[string]any, error) {
	q := url.Values{"league": []string{leagueKey}}
	var out map[string]any
	err := c.GetJSON(ctx, "/league-projection", q, &out)
	return out, err
}
