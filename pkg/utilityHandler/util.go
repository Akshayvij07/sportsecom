package utilityHandler

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GEtAdminIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("adminId")
	adminId, err := strconv.Atoi(fmt.Sprintf("%v", id))
	return adminId, err
}

func GetUserIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("userId")
	userId, err := strconv.Atoi(fmt.Sprintf("%v", id))
	return userId, err
}
func GetOrderIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("userId")
	userId, err := strconv.Atoi(fmt.Sprintf("%v", id))
	return userId, err
}

func GenerateSKU() string {
	sku := make([]byte, 10)

	rand.Read(sku)

	return hex.EncodeToString(sku)
}


func GenerateRandomString(length int) string {
	sku := make([]byte, length)

	rand.Read(sku)

	return hex.EncodeToString(sku)
}
