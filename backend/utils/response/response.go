package response

import "github.com/gin-gonic/gin"

func errorResponse(err error, suffix string) gin.H {
	return gin.H{suffix: err.Error()}
}

func EndWithStatus(c *gin.Context, status int, body any) {
	c.JSON(status, body)
}

func EndWithStatusError(c *gin.Context, status int, suffix string, err error) {
	c.JSON(status, errorResponse(err, suffix))
}

func ExtractAllParams(c *gin.Context) map[string]string {
	request := c.Request

	queryParams := request.URL.Query()

	queryParamsMap := make(map[string]string)

	for key := range queryParams {
		queryParamsMap[key] = queryParams.Get(key)
	}

	return queryParamsMap
}
