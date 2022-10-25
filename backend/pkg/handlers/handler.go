package handlers

import (
    "net/url"
    "os"

	"github.com/courselab/pollex/pollex-backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

type handler struct {
	user controllers.User
}

type Params struct {
	Router *gin.Engine
	User   controllers.User
}

func NewHandler(p *Params) *handler {
	h := &handler{
		user: p.User,
	}

	h.routePing(p.Router)
	h.authRoutes(p.Router)

	return h
}

func (h *handler) routePing(router *gin.Engine) {
	router.GET("/ping", h.ping)
}

func (h *handler) authRoutes(router *gin.Engine) {
    //TODO: figure out a better place for this configuration
    base := os.Getenv("AUTH_SERVICE_URL")
    baseUrl, err := url.Parse(base)
    if len(base) == 0 || err != nil {
        panic("Invalid or missing auth service base url")
    }

    group := router.Group("/")
    group.Use(checkAuth(baseUrl))

    h.routeUsers(group)
}

func (h *handler) routeUsers(router *gin.RouterGroup) {
	router.GET("/users", h.getUsers)
	router.GET("/users/:id", h.getUser)
	router.POST("/users", h.createUser)
	router.PUT("/users/:id", h.updateUser)
	router.DELETE("/users/:id", h.deleteUser)
	router.PATCH("/users/:id", h.patchUser)
}
