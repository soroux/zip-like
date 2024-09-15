package handlers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"zip-like/services/compressor"
)

func CompressHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	compressedLZ77, err := compressor.CompressLZ77(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	encodedHuffman, codes := compressor.CompressHuffman(compressedLZ77)
	c.JSON(http.StatusOK, gin.H{"compressed": encodedHuffman, "codes": codes})
}

func DecompressHandler(c *gin.Context) {
	var requestData struct {
		Compressed string          `json:"compressed"`
		Codes      map[byte]string `json:"codes"`
	}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	decodedLZ77, err := compressor.DecompressHuffman(requestData.Compressed, requestData.Codes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	decompressedData, err := compressor.DecompressLZ77(decodedLZ77)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", decompressedData)
}
