package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	//s := c.Request.Header.Get("Authorization")

	tokenString, err := c.Cookie("UserAuth")
	fmt.Println("test1",tokenString,err)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId, err := ValidateToken(tokenString)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("userId", userId)
	c.Next()
}

/*func LoginHandler(c *gin.Context) {
	// implement login logic here
	// user := c.PostForm("user")
	// pass := c.PostForm("pass")

	// // Throws Unauthorized error
	// if user != "john" || pass != "lark" {
	// 	return c.AbortWithStatus(http.StatusUnauthorized)
	// }

	// Create the Claims
	// claims := jwt.MapClaims{
	// 	"name":  "John Lark",
	// 	"admin": true,
	// 	"exp":   time.Now().Add(time.Hour * 72).Unix(),
	// }

	// Create token
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}*/
