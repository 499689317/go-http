package httpproxy

import (
	// "net/http"
	// "encoding/json"

	"github.com/499689317/go-log"
	"github.com/gin-gonic/gin"
)

var ()

func Init(engine *gin.Engine) bool {

	if engine == nil {
		log.Error().Msg("Init Router Failed")
		return false
	}

	engine.Any("healthz", healthz)

	// TODO
	test := engine.Group("/test")
	{
		v1 := test.Group("/v1")
		{

			// test/v1/:userId
			v1.GET(":userId", test1)
			// test/v1/:serverId/:userId
			v1.POST(":serverId/:userId", test2)
		}
	}

	log.Info().Msg("Init Router ok")
	return true
}

func healthz(c *gin.Context) {
	c.String(200, "ok")
}

// TODO
func test1(c *gin.Context) {

	userId := c.Param("userId")
	c.JSON(200, gin.H{"errCode": 0, "errDesc": "", "data": userId})
}
func test2(c *gin.Context) {

	serverId := c.Param("serverId")
	userId := c.Param("userId")

	c.JSON(200, gin.H{"errCode": 0, "errDesc": "", "data": serverId + userId})
}
