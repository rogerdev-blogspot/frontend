package main

import (
	"fmt"
	"frontend/configs"
	"html/template"
	"io"
	"log"

	_authController "frontend/delivery/controllers/auth"
	_postController "frontend/delivery/controllers/post"
	_userController "frontend/delivery/controllers/user"

	route "frontend/delivery/routes"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func main() {
	config := configs.GetConfig()

	// authRepo := _authRepo.New(db)

	//CONTROLLER
	// authController := _authController.New(authRepo)
	userController := _userController.New(config)
	authController := _authController.New(config)
	postController := _postController.New(config)

	e := echo.New()
	e.Static("/", "assets")
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer

	// e.Validator = &CustomValidator{validator: validator.New()}

	route.RegisterPath(e,
		userController,
		authController,
		postController,
	)
	fmt.Println(config.Port)
	// log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))

}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
