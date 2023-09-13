package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

func NewRouter(g *gin.Engine, list []domain.Ticket) Mhandler {
	//aca instanciamos los repositorios
	repo := tickets.NewRepository(list)

	servi := tickets.NewService(repo)
	mhdnler := handler.NewService(servi)

	return Mhandler{
		servhdnler: mhdnler,
		servicio:   servi,
		dta:        g,
	}
}

type Mhandler struct {
	servhdnler *handler.Service
	servicio   tickets.Service
	dta        *gin.Engine
}

func (h *Mhandler) MapRoutes() {

	mygroup := h.dta.Group("/ticket")
	mygroup.GET("/getByCountry/:dest", h.servhdnler.GetTicketsByCountry())
	mygroup.GET("/getAverage/:dest", h.servhdnler.AverageDestination())
}
