package handler

import (
	"clean-arch/features/user"
	"clean-arch/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) GetAllUsers(c echo.Context) error {
	// panggil func di service layer
	results, errSelect := handler.userService.GetAll()
	if errSelect != nil {
		// return c.JSON(http.StatusInternalServerError, map[string]any{
		// 	"message": "error read data. " + errSelect.Error(),
		// })
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}
	// proses mapping dari core ke response
	usersResult := CoreToResponseList(results)
	// var usersResult []UserResponse
	// for _, value := range results {
	// 	usersResult = append(usersResult, UserResponse{
	// 		ID:    value.ID,
	// 		Name:  value.Name,
	// 		Email: value.Email,
	// 	})
	// }

	// return c.JSON(http.StatusOK, map[string]any{
	// 	"message": "success",
	// 	"data":    usersResult,
	// })
	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", usersResult))
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		// return c.JSON(http.StatusBadRequest, map[string]any{
		// 	"message": "error bind data. data not valid",
		// })
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	//mapping dari request ke core
	userCore := RequestToCore(newUser)
	// userCore := user.Core{
	// 	Name:        newUser.Name,
	// 	Email:       newUser.Email,
	// 	Password:    newUser.Password,
	// 	Address:     newUser.Address,
	// 	PhoneNumber: newUser.PhoneNumber,
	// 	Role:        newUser.Role,
	// }
	errInsert := handler.userService.Create(userCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data"+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *UserHandler) Update(c echo.Context) error {
	id := c.Param("user_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error convert id param", nil))
	}

	var userData = UserRequest{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	userCore := RequestToCore(userData)
	err := handler.userService.Update(idParam, userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("user_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error convert id param", nil))
	}

	err := handler.userService.Delete(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}
