package user

import (
	// "frontend/delivery/controllers/common"

	"fmt"
	"frontend/configs"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	// utils "todo-list-app/utils/aws_S3"
	// "github.com/aws/aws-sdk-go/aws/session"
)

type UserController struct {
	config    *configs.AppConfig
	templates *template.Template
	// conn *session.Session
}

func New(config *configs.AppConfig /*, S3 *session.Session*/) *UserController {
	return &UserController{
		config: config,
		// conn: S3,
	}
}

func (ac *UserController) IndexRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "register.html", map[string]interface{}{
			"name": "Dolly!",
		})
	}
}
func (ac *UserController) IndexLogin() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"name":    "Dolly!",
			"success": true,
			"csrf":    c.Get("csrf"),
		})
	}
}
func (ac *UserController) Login() echo.HandlerFunc {
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

// func (ac *UserController) Register() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusCreated, "Success create user", response))

// 	}
// }

// func (ac *UserController) GetByUid() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		userUid := middlewares.ExtractTokenUserUid(c)

// 		res, err := ac.repo.GetByUid(userUid)

// 		if err != nil {
// 			statusCode := http.StatusInternalServerError
// 			errorMessage := "There is some problem from the server"
// 			if err.Error() == "record not found" {
// 				statusCode = http.StatusNotFound
// 				errorMessage = err.Error()
// 			}
// 			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
// 		}

// 		res.CreatedAt = ac.TimeToUser(res.CreatedAt.(int64))
// 		res.UpdatedAt = ac.TimeToUser(res.UpdatedAt.(int64))
// 		res.DeletedAt = ac.TimeToUser(res.DeletedAt.(int64))

// 		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get user", res))
// 	}
// }
// func (ac *UserController) Search() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		q := c.QueryParam("q")

// 		res, err := ac.repo.Search(q)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "internal server error for get all "+err.Error(), nil))
// 		}

// 		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success Get all Room", res))
// 	}
// }

// func (ac *UserController) Update() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		userUid := middlewares.ExtractTokenUserUid(c)
// 		var newUser = UpdateUserRequestFormat{}
// 		c.Bind(&newUser)

// 		err := c.Validate(&newUser)
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
// 		}

// 		// resGet, errGet := ac.repo.GetById(user_uid)
// 		// if errGet != nil {
// 		// 	log.Info(resGet)
// 		// }

// 		// file, errO := c.FormFile("image")
// 		// if errO != nil {
// 		// 	log.Info(errO)
// 		// } else if errO == nil {
// 		// 	src, _ := file.Open()
// 		// 	if resGet.Image != "" {
// 		// 		var updateImage = resGet.Image
// 		// 		updateImage = strings.Replace(updateImage, "https://airbnb-app.s3.ap-southeast-1.amazonaws.com/", "", -1)

// 		// 		var resUp = utils.UpdateUpload(ac.conn, updateImage, src, *file)
// 		// 		if resUp != "success to update image" {
// 		// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server"+resUp, nil))
// 		// 		}
// 		// 	} else if resGet.Image == "" {
// 		// 		var image, errUp = utils.Upload(ac.conn, src, *file)
// 		// 		if errUp != nil {
// 		// 			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "Upload Failed", nil))
// 		// 		}
// 		// 		newUser.Image = image
// 		// 	}
// 		// }

// 		res, err_repo := ac.repo.Update(userUid, entities.User{
// 			Name:     newUser.Name,
// 			Email:    newUser.Email,
// 			Password: newUser.Password,
// 			Address:  newUser.Address,
// 			Gender:   newUser.Gender,
// 			// Image:    newUser.Image,
// 		})

// 		if err_repo != nil {
// 			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
// 		}

// 		response := UserUpdateResponse{}
// 		response.UserUid = res.UserUid
// 		response.Name = res.Name
// 		response.Email = res.Email
// 		response.Address = res.Address
// 		response.Gender = res.Gender
// 		// response.Roles = res.Roles
// 		// response.Image = res.Image

// 		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success update user", response))
// 	}
// }

// func (ac *UserController) Delete() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		userUid := middlewares.ExtractTokenUserUid(c)
// 		err := ac.repo.Delete(userUid)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
// 		}

// 		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success delete user", nil))
// 	}
// }
// func (ac *UserController) Dummy() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		q, _ := strconv.Atoi(c.QueryParam("length"))

// 		result := ac.repo.Dummy(q)
// 		if !result {
// 			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
// 		}
// 		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success create user dummy", nil))

// 	}
// }
// func (c *UserController) TimeToUser(timeInt int64) string {
// 	if timeInt <= 0 {
// 		return ""
// 	}
// 	i, err := strconv.ParseInt(strconv.Itoa(int(timeInt)), 10, 64)
// 	if err != nil {
// 		panic(err)
// 	}
// 	tm := time.Unix(i, 0)
// 	year := strconv.Itoa(tm.Year())
// 	month := strconv.Itoa(int(tm.Month()))
// 	h, m, s := tm.Clock()
// 	hour := strconv.Itoa(h)
// 	if len(hour) == 1 {
// 		hour = "0" + hour
// 	}
// 	minute := strconv.Itoa(m)
// 	if len(minute) == 1 {
// 		minute = "0" + minute
// 	}
// 	second := strconv.Itoa(s)

// 	switch month {
// 	case "1":
// 		month = "Januari"
// 	case "2":
// 		month = "Februari"
// 	case "3":
// 		month = "Maret"
// 	case "4":
// 		month = "April"
// 	case "5":
// 		month = "Mei"
// 	case "6":
// 		month = "Juni"
// 	case "7":
// 		month = "Juli"
// 	case "8":
// 		month = "Agustus"
// 	case "9":
// 		month = "September"
// 	case "10":
// 		month = "Oktober"
// 	case "11":
// 		month = "November"
// 	case "12":
// 		month = "Desember"
// 	}
// 	day := strconv.Itoa(tm.Day())
// 	createdOnString := day + " " + month + " " + year + " Pukul " + hour + ":" + minute + ":" + second + " WIB"
// 	return createdOnString
// }
