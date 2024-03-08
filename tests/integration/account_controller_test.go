package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jooaos/pismo/internal/controller"
	"github.com/jooaos/pismo/internal/model"
	"github.com/jooaos/pismo/internal/repository/adapter"
	"github.com/jooaos/pismo/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestAccountControler_Create(t *testing.T) {
	SetUp()
	accountController := setUpAccountController()

	app := fiber.New()
	app.Post("/accounts", accountController.Create)

	t.Run("should create account with success", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/accounts", strings.NewReader(`{"document_number": "12345678900"}`))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)
		assert.Equal(t, 201, resp.StatusCode)
	})

	t.Run("should return error invalid document number", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/accounts", strings.NewReader(`{"document_number": ""}`))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("should return error account already exists", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/accounts", strings.NewReader(`{"document_number": "12345678900"}`))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("should return error body parse", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/accounts", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)
		assert.Equal(t, 400, resp.StatusCode)
	})

	SetDown()
}

func TestAccountControler_GetById(t *testing.T) {
	SetUp()
	accountController := setUpAccountController()

	app := fiber.New()
	app.Post("/accounts", accountController.Create)
	app.Get("/accounts/:id", accountController.GetById)

	t.Run("should return user with success ", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/accounts", strings.NewReader(`{"document_number": "12345678900"}`))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req, -1)

		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		fmt.Printf("body %s", string(body))

		data := struct {
			Data model.Account `json:"data"`
		}{}
		json.Unmarshal([]byte(body), &data)

		req = httptest.NewRequest("GET", fmt.Sprintf("/accounts/%d", data.Data.ID), nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ = app.Test(req, -1)

		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("should return error for user not found", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/accounts", strings.NewReader(`{"document_number": "12345678901"}`))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req, -1)

		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		fmt.Printf("body %s", string(body))

		data := struct {
			Data model.Account `json:"data"`
		}{}
		json.Unmarshal([]byte(body), &data)

		req = httptest.NewRequest("GET", fmt.Sprintf("/accounts/%d", data.Data.ID+100), nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ = app.Test(req, -1)

		assert.Equal(t, 400, resp.StatusCode)
	})

	SetDown()
}

func setUpAccountController() *controller.AccountController {
	accountRepository := adapter.NewAccountRepositoryMariaDB(BaseIntegrationTest.DB)
	accountService := service.NewAccountService(accountRepository)
	accountController := controller.NewAccountController(accountService)
	return accountController
}
