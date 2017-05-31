package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// DatabaseType ...
type DatabaseType interface {
	AddUnit(unit *Unit) error
	GetUnit(id string) (*Unit, error)
	ListUnits() ([]Unit, error)

	AddGroup(group *Group) error
	GetGroup(id string) (*Group, error)
	ListGroups() ([]Group, error)
	AddUnitToGroup(unit *Unit, group *Group) error
}

// Server ...
type Server struct {
	Port string

	db   DatabaseType
	echo *echo.Echo
}

// NewServer ...
func NewServer(port string, db DatabaseType, production bool) *Server {
	// Initialize HTTP server
	e := echo.New()
	e.HideBanner = true
	server := &Server{port, db, e}

	e.GET("/unit", server.handleUnit)
	e.GET("/unit/:id", server.handleUnit)
	e.GET("/group", server.handleGroup)
	e.GET("/group/:id", server.handleGroup)

	return server
}

// Start ...
func (server *Server) Start() error {
	return server.echo.Start(server.Port)
}

// Shutdown ...
func (server *Server) Shutdown(ctx context.Context) error {
	if err := server.echo.Shutdown(ctx); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (server *Server) handleGroup(c echo.Context) error {
	var (
		result interface{}
		err    error
	)
	if id := c.Param("id"); len(id) > 0 {
		result, err = server.db.GetGroup(id)
	} else {
		result, err = server.db.ListGroups()
	}
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if result == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, result)
}

func (server *Server) handleUnit(c echo.Context) error {
	var (
		result interface{}
		err    error
	)
	if id := c.Param("id"); len(id) > 0 {
		result, err = server.db.GetUnit(id)
	} else {
		result, err = server.db.ListUnits()
	}
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if result == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, result)
}
