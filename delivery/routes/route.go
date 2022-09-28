package route

import (
	"frontend/delivery/controllers/auth"
	"frontend/delivery/controllers/post"
	"frontend/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo,
	uc *user.UserController,
	ac *auth.AuthController,
	pc *post.PostController,

) {
	const CSRFTokenHeader = "X-CSRF-Token"
	const CSRFKey = "csrf"
	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + CSRFTokenHeader,
		ContextKey:  CSRFKey,
	}))

	//ROUTE REGISTER - LOGIN USERS
	e.GET("/register", ac.IndexRegister())
	e.POST("/register", ac.Register())
	e.GET("/login", ac.IndexLogin())
	e.POST("/login", ac.Login())

	e.GET("/posts", pc.IndexPosts())
	e.GET("/posts/:id", pc.GetPostDetail())

	// e.POST("users/login", aa.Login())

	//ROUTE USERS
	// e.GET("/users", uc.GetById())
	// e.PUT("/users", uc.Update())
	// e.DELETE("/users", uc.Delete())

}
