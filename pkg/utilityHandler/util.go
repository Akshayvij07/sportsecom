package utilityHandler

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"

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

func GenerateInvoiceNumber() string {
	// Get the current timestamp in milliseconds
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// Generate a random number between 100 and 999 (3 digits)
	randomNum := rand.Intn(900) + 100

	// Combine the timestamp and random number to create the invoice number
	invoiceNumber := fmt.Sprintf("INV%d%03d", timestamp, randomNum)

	return invoiceNumber
}
