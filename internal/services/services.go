package services

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/siddarthtv/golang_assignment_kelp/internal/db"
	"golang.org/x/sync/singleflight"
)

type Service struct {
	conn  *db.DBConn
	group singleflight.Group
}

func InitService() *Service {
	service := new(Service)
	service.conn = db.InitDB()
	return service
}

func (s *Service) FinancialsHandler(c *gin.Context) {
	companyId := c.Query("companyId")
	if companyId == "" {
		c.String(http.StatusBadRequest, "Missing company ID in request")
		return
	}
	companyIdNum, convErr := strconv.Atoi(companyId)
	if convErr != nil {
		c.String(http.StatusBadRequest, "Bad Request: %s", convErr.Error())
		return
	}
	key := fmt.Sprintf("sales-%d", companyIdNum)
	obj, err := s.sendRequest(companyIdNum, key, s.conn.CompanyFinancials)
	if err == sql.ErrNoRows {
		c.String(http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, obj)
}

func (s *Service) SalesHandler(c *gin.Context) {
	companyId := c.Query("companyId")
	if companyId == "" {
		c.String(http.StatusBadRequest, "Missing company ID in request")
		return
	}
	companyIdNum, convErr := strconv.Atoi(companyId)
	if convErr != nil {
		c.String(http.StatusBadRequest, "Bad Request: %s", convErr.Error())
		return
	}
	key := fmt.Sprintf("sales-%d", companyIdNum)
	obj, err := s.sendRequest(companyIdNum, key, s.conn.CompanySales)
	if err == sql.ErrNoRows {
		c.String(http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, obj)
}

func (s *Service) StatsHandler(c *gin.Context) {
	companyId := c.Query("companyId")
	if companyId == "" {
		c.String(http.StatusBadRequest, "Missing company ID in request")
		return
	}
	companyIdNum, convErr := strconv.Atoi(companyId)
	if convErr != nil {
		c.String(http.StatusBadRequest, "Bad Request: %s", convErr.Error())
		return
	}
	key := fmt.Sprintf("sales-%d", companyIdNum)
	obj, err := s.sendRequest(companyIdNum, key, s.conn.CompanyEmployeeStats)
	if err == sql.ErrNoRows {
		c.String(http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, obj)
}

func (s *Service) sendRequest(companyID int, key string, fetchfunc func(int) (interface{}, error)) (interface{}, error) {
	result, err, _ := s.group.Do(key, func() (interface{}, error) {
		log.Printf("Executing the API logic - %s\n", key)
		// time.Sleep(2 * time.Second) // can be used to slow down the process and show the api calls getting grouped
		return fetchfunc(companyID)
	})
	return result, err
}
