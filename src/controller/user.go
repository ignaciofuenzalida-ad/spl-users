package controller

import (
	"spl-users/src/dto"
	"spl-users/src/helpers"
	"spl-users/src/service"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *service.UserService
	Validator   *validator.Validate
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
		// TODO: Move validator to a global DI scope
		Validator: validator.New(),
	}
}

func (u *UserController) GetUsers(c *fiber.Ctx) error {
	queryParams := c.Queries()
	search := queryParams["search"]
	limitStr := queryParams["limit"]
	locationsStr := queryParams["locations"]

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "limit parameter must be a number"})
	}
	if limit <= 0 || limit > 20 {
		limit = 20
	}

	var locations []string
	if locationsStr != "" {
		locations = strings.Split(locationsStr, ",")
	}

	users, err := u.userService.GetAllUsers(search, limit, locations)
	if err != nil {
		return helpers.InternalError(c, err)
	}

	return c.JSON(fiber.Map{"data": users})
}

func (u *UserController) GetUserByRun(c *fiber.Ctx) error {
	runStr := c.Params("run")
	run, err := strconv.Atoi(runStr)
	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "run parameter must be a number"})
	}

	user, err := u.userService.GetUserByRun(run)
	if err != nil {
		return helpers.InternalError(c, err)
	}

	if user == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(fiber.Map{"data": user})
}

func (u *UserController) GetRandomUsers(c *fiber.Ctx) error {
	runs := u.userService.GetRandomUsers()

	return c.JSON(fiber.Map{"data": runs})
}

func (u *UserController) UpdateOrCreateUserByRun(c *fiber.Ctx) error {
	// Validate if Run is valid and exist
	runStr := c.Params("run")
	run, err := strconv.Atoi(runStr)
	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "run parameter must be a number"})
	}

	user, err := u.userService.GetUserQueueByRun(run)
	if err != nil {
		return helpers.InternalError(c, err)
	}

	if user == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Validate request body
	var userRequest dto.UpdateUserDto
	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	// Validate the request struct
	if err := u.Validator.Struct(userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = u.userService.UpdateOrCreateUser(run, userRequest)
	if err != nil {
		return helpers.InternalError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (u *UserController) GetQueueUsersStatistics(c *fiber.Ctx) error {
	data, err := u.userService.GetQueueUsersStatistics()
	if err != nil {
		return helpers.InternalError(c, err)
	}

	return c.JSON(fiber.Map{"data": data})
}
