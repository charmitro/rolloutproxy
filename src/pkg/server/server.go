package server

import (
	"log"
	"net/http"
	"os"

	"github.com/charmitro/rolloutproxy/src/pkg/cluster"
	"github.com/gin-gonic/gin"
)

func basicRollout(c *gin.Context) {
	if c.Param("deployment") != "" && c.Param("namespace") != "" {
		if err := cluster.RolloutRestart(
			c.Param("deployment"),
			c.Param("namespace"),
		); err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"deployemnt": c.Param("deployment"),
				"namespace":  c.Param("namespace"),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"deployemnt": c.Param("deployment"),
				"namespace":  c.Param("namespace"),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters",
		})
	}
}

func Init() {
	cluster.Init()
	r := gin.Default()

	auth := r.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("USERNAME"): os.Getenv("PASSWORD"),
	}))

	auth.POST("/:deployment/:namespace", basicRollout)

	r.Run()
}
