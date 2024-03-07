package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jooaos/pismo/internal/service"
	"github.com/jooaos/pismo/internal/utils"
)

type AccountController struct {
	accountService service.IAccountService
}

type IAccountController interface {
	Create(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

func NewAccountController(accountService service.IAccountService) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}

type AccountControllerCreateRequest struct {
	DocumentNumber string `json:"document_number"`
}

func (ac *AccountController) Create(c *fiber.Ctx) error {
	request := new(AccountControllerCreateRequest)

	err := c.BodyParser(request)
	if err != nil {
		log.Printf("[AccountController::Create] Error while parsing body: %s", err.Error())
		return c.Status(500).JSON(utils.NewApiErrorResponse("error while parsing body"))
	}

	if request.DocumentNumber == "" {
		log.Printf("[AccountController::Create] Request without document number")
		return c.Status(400).JSON(utils.NewApiErrorResponse("document_number is required"))
	}

	account, err := ac.accountService.Create(request.DocumentNumber)
	if err != nil {
		log.Printf("[AccountController::Create] Error at create account: %s", err.Error())
		return c.Status(400).JSON(utils.NewApiErrorResponse(err.Error()))
	}

	return c.Status(201).JSON(utils.NewApiSuccessResponse(account))
}

func (ac *AccountController) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		log.Printf("[AccountController::GetById] Request without account id")
		return c.Status(400).JSON(utils.NewApiErrorResponse("account id is required"))
	}

	accountId, _ := strconv.ParseUint(id, 10, 64)

	account, err := ac.accountService.GetById(uint32(accountId))
	if err != nil {
		log.Printf("[AccountController::GetById] Error while gettint account by id: %s", err.Error())
		return c.Status(400).JSON(utils.NewApiErrorResponse(err.Error()))
	}

	return c.Status(200).JSON(utils.NewApiSuccessResponse(account))
}
