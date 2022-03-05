package handlers

import (
	"github.com/Pudgekim/domain/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	userRepo repository.UserRepository
}

func NewHandler(userRepo repository.UserRepository) *Handler {
	handler := &Handler{
		userRepo: userRepo,
	}

	return handler
}

func (h Handler) Routes() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLFiles("templates/index.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "main page",
		})
	})

	router.GET("/auth/google/login", h.GoogleLoginHandler)
	router.GET("/auth/google/callback", h.GoogleAuthCallBack)
	//router.GET("/check", func(c *gin.Context) {
	//	fmt.Println("id: ", os.Getenv("GOOGLE_CLIENT_ID"))
	//	fmt.Println("secret: ", os.Getenv("GOOGLE_SECRET_KEY"))
	//})

	router.GET("/auth/user/:id", h.GetUserById)

	return router
}
