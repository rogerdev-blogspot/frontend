package post

import (
	// "frontend/delivery/controllers/common"

	"frontend/configs"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	// utils "todo-list-app/utils/aws_S3"
	// "github.com/aws/aws-sdk-go/aws/session"
)

type PostController struct {
	config    *configs.AppConfig
	templates *template.Template
	// conn *session.Session
}

func New(config *configs.AppConfig /*, S3 *session.Session*/) *PostController {
	return &PostController{
		config: config,
		// conn: S3,
	}
}

func (ac *PostController) IndexPosts() echo.HandlerFunc {
	return func(c echo.Context) error {

		type Post struct {
			Title       string
			Description string
		}
		var post Post
		var posts []Post

		post.Title = "judul"
		post.Description = "hellosfksdjfkdsjfds"
		posts = append(posts, post)
		posts = append(posts, post)

		return c.Render(http.StatusOK, "posts.html", map[string]interface{}{
			"name":    "RogerDev!",
			"success": true,
			"csrf":    c.Get("csrf"),
			"posts":   posts,
		})
	}
}
func (ac *PostController) GetPostDetail() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.Render(http.StatusOK, "post.html", map[string]interface{}{
			"name":    "RogerDev!",
			"success": true,
			"csrf":    c.Get("csrf"),
		})
	}
}
