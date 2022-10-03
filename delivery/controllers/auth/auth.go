package auth

import (
	// "frontend/delivery/controllers/common"

	"bytes"
	"encoding/json"
	"fmt"
	"frontend/configs"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	// utils "todo-list-app/utils/aws_S3"
	// "github.com/aws/aws-sdk-go/aws/session"
)

type AuthController struct {
	config    *configs.AppConfig
	templates *template.Template
	// conn *session.Session
}

func New(config *configs.AppConfig /*, S3 *session.Session*/) *AuthController {
	return &AuthController{
		config: config,
		// conn: S3,
	}
}

func (ac *AuthController) IndexRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "register.html", map[string]interface{}{
			"name":    "RogerDev!",
			"success": true,
			"csrf":    c.Get("csrf"),
		})
	}
}
func (ac *AuthController) IndexLogin() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"name":    "RogerDev!",
			"success": true,
			"csrf":    c.Get("csrf"),
		})
	}
}
func (ac *AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := make(map[string]interface{})
		if err := c.Bind(&data); err != nil {
			return err
		}
		fmt.Println(data)
		message := fmt.Sprintf("hello %s", data["name"])
		return c.JSON(http.StatusOK, message)
		// return c.JSON(http.StatusOK, message)
	}
}
func (ac *AuthController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

		data := make(map[string]interface{})
		if err := c.Bind(&data); err != nil {
			return c.Redirect(400, "/register")
		}
		json_data, err := json.Marshal(&data)
		var client = &http.Client{}
		baseURL := ac.config.UserService.Url
		request, err := http.NewRequest("POST", baseURL+"/register", bytes.NewBuffer(json_data))
		fmt.Println(baseURL)
		if err != nil {
			fmt.Println(err.Error())
			// return c.JSON(http.StatusBadRequest, "failed register")
			return c.Redirect(400, "/register")

		}
		request.Header.Set("Content-Type", "application/json")

		// request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

		res, err2 := client.Do(request)
		if err2 != nil {
			fmt.Println(err.Error())

			// return c.JSON(http.StatusBadRequest, "failed register")
			return c.Redirect(400, "/register")

		}
		type DataResponse struct {
			UserUid string `json:"user_uid"`
			Name    string `json:"name"`
			Email   string `json:"email"`
			Address string `json:"address"`
			Gender  string `json:"gender"`
		}
		type ResponseBackend struct {
			Code    interface{}  `json:"code"`
			Message interface{}  `json:"message"`
			Data    DataResponse `json:"data"`
		}

		var dataRegister ResponseBackend

		defer res.Body.Close()
		err3 := json.NewDecoder(res.Body).Decode(&dataRegister)
		if err3 != nil {
			fmt.Println(err.Error())
			return c.Redirect(400, "/register")
		}
		fmt.Println(dataRegister)

		// message := fmt.Sprintf("hello %s", data["name"])
		// return c.JSON(http.StatusOK, message)

		return c.Redirect(200, "/login")
	}
}
