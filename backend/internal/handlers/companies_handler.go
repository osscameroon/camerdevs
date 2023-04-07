package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/osscameroon/jobsika/internal/server"
	log "github.com/sirupsen/logrus"
)

// GetCompanies return a list of companies
func GetCompanies(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find companies"})
		return
	}

	companies, err := db.GetCompanies()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find companies"})
		return
	}

	c.JSON(http.StatusOK, companies)
}

// GetCompanyByID returns a company by id
func GetCompanyByID(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find salaries"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "failed to parse parameters"})
		return
	}

	company, err := db.GetCompanyByID(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound,
			gin.H{"error": "could not find company"})
		return
	}

	c.JSON(http.StatusOK, company)
}
