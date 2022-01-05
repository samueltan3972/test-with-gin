package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"time"

	"test/gin-test/database"
	"test/gin-test/models"

	"github.com/gin-gonic/gin"
)

var fruitModel = new(models.FruitModel)
var dummyModel = new(models.DummyModel)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/database", func(c *gin.Context) {
		// Create
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(1000-1) + 1
		dt := time.Now()
		md5hash := GetMD5Hash(dt.String() + fmt.Sprint(randNum))

		err := dummyModel.Create(md5hash, "Test", "Test2")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
			return
		}

		// Update
		updateerr := dummyModel.Update(md5hash, "Test3")

		if updateerr != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": updateerr.Error()})
			return
		}

		// Delete
		deleteerr := dummyModel.Delete(md5hash)

		if deleteerr != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": deleteerr.Error()})
			return
		}

		// Select
		fruit, fruiterr := fruitModel.GetAll()

		if fruiterr != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": deleteerr.Error()})
			return
		}

		c.JSON(http.StatusOK, fruit)
	})

	r.GET("/fibonacci", func(c *gin.Context) {
		result := "[0, 1"
		count := 5000
		n1 := big.NewInt(0)
		n2 := big.NewInt(1)

		for i := 2; i < count; i++ {
			n1.Add(n1, n2)
			result = result + ", " + fmt.Sprint(n1)
			n1, n2 = n2, n1
		}

		result = result + "]"
		c.String(http.StatusOK, result)
	})

	return r
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	database.Init()

	router := setupRouter()

	router.Run("0.0.0.0:8080")
}
