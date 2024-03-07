package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jooaos/pismo/internal/service"
	"github.com/jooaos/pismo/internal/utils"
)

type TransactionController struct {
	transactionService service.ITransactionService
}

type ITransactionController interface {
	Create(c *fiber.Ctx) error
}

func NewTransactionController(transactionService service.ITransactionService) *TransactionController {
	return &TransactionController{
		transactionService: transactionService,
	}
}

type TransactionControllerCreateRequest struct {
	AccountId       int     `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Ammount         float32 `json:"amount"`
}

func (tr *TransactionController) Create(c *fiber.Ctx) error {
	request := new(TransactionControllerCreateRequest)

	err := c.BodyParser(request)
	if err != nil {
		log.Printf("[TransactionController::Create] Error while parsing body: %s", err.Error())
		return c.Status(500).JSON(utils.NewApiErrorResponse("error while parsing body"))
	}

	if request.AccountId == 0 || request.OperationTypeId == 0 || request.Ammount == 0 {
		log.Printf("[TransactionController::Create] Request without required parameters")
		return c.Status(400).JSON(utils.NewApiErrorResponse("account_id, operation_type_id and amount are required"))
	}

	transaction, err := tr.transactionService.CreateTransaction(
		request.AccountId,
		request.OperationTypeId,
		request.Ammount,
	)
	if err != nil {
		log.Printf("[TransactionController::Create] Error at create transaction: %s", err.Error())
		return c.Status(400).JSON(utils.NewApiErrorResponse(err.Error()))
	}

	return c.Status(201).JSON(utils.NewApiSuccessResponse(transaction))
}
