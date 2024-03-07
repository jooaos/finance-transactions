package integration

import (
	"bytes"
	"encoding/json"
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

func TestTransactionController_Create_Success(t *testing.T) {
	SetUp()
	account := accountSetup()
	transactionController := setUpTransactionController()

	app := fiber.New()
	app.Post("/transactions", transactionController.Create)

	cases := []struct {
		request controller.TransactionControllerCreateRequest
		code    int
	}{
		{
			request: controller.TransactionControllerCreateRequest{
				AccountId:       int(account.ID),
				OperationTypeId: int(model.CASH_PURCHASE),
				Ammount:         -20.20,
			},
			code: 201,
		},
		{
			request: controller.TransactionControllerCreateRequest{
				AccountId:       int(account.ID),
				OperationTypeId: int(model.INSTALLMENT_PURCHASE),
				Ammount:         -20.20,
			},
			code: 201,
		},
		{
			request: controller.TransactionControllerCreateRequest{
				AccountId:       int(account.ID),
				OperationTypeId: int(model.WITHDRAWAL),
				Ammount:         -20.20,
			},
			code: 201,
		},
		{
			request: controller.TransactionControllerCreateRequest{
				AccountId:       int(account.ID),
				OperationTypeId: int(model.PAYMENT),
				Ammount:         20.20,
			},
			code: 201,
		},
	}

	for _, item := range cases {
		t.Run("shoul create transaction with success", func(t *testing.T) {
			requestBody, _ := json.Marshal(&item.request)
			req := httptest.NewRequest("POST", "/transactions", bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req, -1)
			assert.Equal(t, item.code, resp.StatusCode)
		})
	}

	SetDown()
}

func TestTransactionController_Create_FailAmmount(t *testing.T) {
	SetUp()
	account := accountSetup()
	transactionController := setUpTransactionController()

	app := fiber.New()
	app.Post("/transactions", transactionController.Create)

	cases := []struct {
		request controller.TransactionControllerCreateRequest
		code    int
	}{
		{
			request: controller.TransactionControllerCreateRequest{
				AccountId:       int(account.ID),
				OperationTypeId: int(model.CASH_PURCHASE),
				Ammount:         20.20,
			},
			code: 400,
		},
		{
			request: controller.TransactionControllerCreateRequest{
				AccountId:       int(account.ID),
				OperationTypeId: int(model.INSTALLMENT_PURCHASE),
				Ammount:         20.20,
			},
			code: 400,
		},
		{
			request: controller.TransactionControllerCreateRequest{
				AccountId:       int(account.ID),
				OperationTypeId: int(model.WITHDRAWAL),
				Ammount:         20.20,
			},
			code: 400,
		},
		{
			request: controller.TransactionControllerCreateRequest{
				AccountId:       int(account.ID),
				OperationTypeId: int(model.PAYMENT),
				Ammount:         -20.20,
			},
			code: 400,
		},
	}

	for _, item := range cases {
		t.Run("should return amount error", func(t *testing.T) {
			requestBody, _ := json.Marshal(&item.request)
			req := httptest.NewRequest("POST", "/transactions", bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req, -1)
			assert.Equal(t, item.code, resp.StatusCode)
		})
	}

	SetDown()
}

func TestTransactionController_Create_General(t *testing.T) {
	SetUp()
	transactionController := setUpTransactionController()
	account := accountSetup()

	app := fiber.New()
	app.Post("/transactions", transactionController.Create)

	t.Run("should return params validation error", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/transactions", strings.NewReader(`{"tst": "tst"}`))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("should return operation type invalid error", func(t *testing.T) {
		body := controller.TransactionControllerCreateRequest{
			AccountId:       int(account.ID),
			OperationTypeId: 9999,
			Ammount:         20.20,
		}
		requestBody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/transactions", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)
		assert.Equal(t, 400, resp.StatusCode)
	})

	SetDown()
}

func setUpTransactionController() *controller.TransactionController {
	accountRepository := adapter.NewAccountRepositoryMariaDB(BaseIntegrationTest.DB)
	transactionRepository := adapter.NewTransactionRepositoryMariaDB(BaseIntegrationTest.DB)
	transactionService := service.NewTransactionService(transactionRepository, accountRepository)
	transactionController := controller.NewTransactionController(transactionService)
	return transactionController
}

func accountSetup() *model.Account {
	var account *model.Account
	_ = BaseIntegrationTest.DB.Raw("INSERT INTO accounts (document_number) VALUES (12345678900) RETURNING *").Scan(&account)
	return account
}
