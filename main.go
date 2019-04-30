package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		c.File("index.htm")
	})

	r.POST("/upload", Upload)
	r.GET("/image/:fileName", GetImage)

	err := r.Run(":3000")
	if err != nil {
		log.Fatal("server start error: ", err)
	}
}

func Upload(c *gin.Context) {

	file, header, err := c.Request.FormFile("upload")
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"status": gin.H{
				"status_code": http.StatusBadRequest,
				"description": err,
			},
		})
		return
	}

	filename := header.Filename

	log.Println(file, err, filename)

	//prefix = date format(yyyyMMdd_)
	prefix := time.Now().Format("20060102") + "_"

	f, err := os.Create("./upload/orign/" + prefix + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"status_code": http.StatusOK,
			"file_name":   filename,
			"description": "upload successfully",
		},
	})

}

func GetImage(c *gin.Context) {

	//image name
	name := c.Param("fileName")

	//query string (resize if any)
	width := c.Query("width")
	height := c.Query("height")

	//resize if any
	if width != "" && height != "" {

		log.Println("prepare a csutom image")

		new_width := 0
		new_height := 0

		new_width, _ = strconv.Atoi(width)

		new_height, _ = strconv.Atoi(height)

		src, err := imaging.Open("./upload/orign/" + name)
		if err != nil {
			log.Println("fail with open : ", err)
			return
		}

		//resze
		src = imaging.Resize(src, new_width, new_height, imaging.Lanczos)

		//save new image
		//prefix = img new size
		size := fmt.Sprintf("%d*%d_", new_width, new_height)
		err = imaging.Save(src, "./upload/custom/"+size+name)
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}

		//download
		c.File("./upload/custom/" + size + name)

	} else {

		log.Println("prepare an orign image")

		//download
		c.File("./upload/orign/" + name)
	}

}
