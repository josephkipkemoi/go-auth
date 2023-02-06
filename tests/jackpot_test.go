package tests

import (
	"bytes"
	"encoding/json"
	"go-auth-api/go-auth/controllers"
	"go-auth-api/go-auth/database/factory"
	"go-auth-api/go-auth/env"
	"go-auth-api/go-auth/models"
	"go-auth-api/go-auth/routes"
	"log"
	"strconv"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPostJackpotMarkets(t *testing.T) {
	url := env.GetDevAppUrl()
	r := routes.SetupRouter()
	w := httptest.NewRecorder()

	i := controllers.JackpotMarketInput{
		Market: "Mega Jackpot",
		JackpotMarketID: models.MegaJackpotMarketId,
	}

	b,e := json.MarshalIndent(i, "", " ")
	if e != nil {
		log.Fatal(e)
	}

	bodyReader := bytes.NewReader(b)

	req := httptest.NewRequest("POST", url + "api/v1/jackpots", bodyReader)
	r.ServeHTTP(w, req)

	assert.Contains(t, w.Header().Get("Content-Type"), "application/json:charset=utf-8", "Should have content-type data format set to json")
	assert.Contains(t, w.Header().Get("Authorization"), "", "Should have valid jwt authorization header")
	assert.JSONEq(t,`{"status":"201 Created", "market": "Mega Jackpot"}`,w.Body.String(), "Should have JSON BODY")
	assert.Equal(t,http.StatusCreated, w.Code, "Should return resource created http status code")
}

func TestCanPostJackpotGames(t *testing.T) {
	url := env.GetDevAppUrl()
	r := routes.SetupRouter()
	w := httptest.NewRecorder()

	i := controllers.JackpotGamesInput{
		JackpotMarketID: 1,
		HomeTeam: "Team A",
		AwayTeam: "Team B",
		HomeOdds: 3.35,
		DrawOdds: 4.55,
		AwayOdds: 3.85,
	}
	b,e := json.MarshalIndent(i, "", " ")
	if e != nil {
		log.Fatal(e)
	}

	bodyReader := bytes.NewReader(b)

	req := httptest.NewRequest("POST", url + "api/v1/jackpots/games", bodyReader)
	r.ServeHTTP(w, req)

	assert.Contains(t, w.Header().Get("Content-Type"), "application/json:charset=utf-8", "Should have content-type data format set to json")
	assert.Contains(t, w.Header().Get("Authorization"), "", "Should have valid jwt authorization header")
	assert.Equal(t,http.StatusCreated, w.Code, "Should return resource created http status code")
}

func TestCanGetJackpotGamesByJackpotMarketId(t *testing.T) {
	url := env.GetDevAppUrl()
	r := routes.SetupRouter()
	w := httptest.NewRecorder()

	m := factory.MakeJackpotMarket()
	factory.MakeJackpotGames()
	mId := strconv.Itoa(int(m.JackpotMarketID))

	req := httptest.NewRequest("GET", url + "api/v1/jackpots/games?jp_id=" + mId, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json:charset=utf-8", w.Header().Get("Content-Type"), "Should have content-type header set to application/json")
	// assert.Contains(t, w.Body.String(), `{"TeamA"}`, "Should have right JSON Body", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code, "Should return success status code")
}

func TestCanUpdateJackpotGame(t *testing.T) {
	url := env.GetDevAppUrl()
	r := routes.SetupRouter()
	w := httptest.NewRecorder()

	i := controllers.UpdateJackpotGameInput{
		HomeTeam: "A",
		AwayTeam: "B",
		HomeOdds: 1.2,
		DrawOdds: 1.3,
		AwayOdds: 1.5,
	}

	b,e := json.MarshalIndent(i, "", " ")
	if e != nil {
		log.Fatal(e)
	}
	bodyReader := bytes.NewReader(b)

	req := httptest.NewRequest("PATCH", url + "api/v1/jackpots/games/patch?id=1", bodyReader)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json:charset=utf-8", w.Header().Get("Content-Type"), "Should have content-type header set to application/json")
	assert.Equal(t, http.StatusOK, w.Code, "Should return success status code")
}