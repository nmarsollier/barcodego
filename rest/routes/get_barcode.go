package routes

import (
	"github.com/gin-gonic/gin"

	qrcode "github.com/skip2/go-qrcode"

	"github.com/nmarsollier/go_barcode/err"
)

/**
 * @api {get} /barcode/:barcode Obtener Barcode
 * @apiName Obtener Barcode
 * @apiGroup Barcode
 *
 * @apiDescription Obtiene un barcode
 *
 * @apiUse SizeHeader
 *
 * @apiSuccessExample {png} Respuesta
 *    Imagen en formato png
 *
 * @apiUse CustomError
 */
func init() {
	router().GET(
		"/barcode/:barcode",
		validateBarcode,
		getBarcode,
	)
}

func validateBarcode(c *gin.Context) {
	barcodeStr := c.Param("barcode")

	if len(barcodeStr) <= 0 || len(barcodeStr) > 1024 {
		c.Error(err.NewCustom(400, "Invalid Barcode"))
		c.Abort()
		return
	}
}

func getBarcode(c *gin.Context) {
	barcodeStr := c.Param("barcode")
	scale := getSizeParam(c)

	png, err := qrcode.Encode(barcodeStr, qrcode.Medium, scale)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.Data(200, "image/png", png)
}

/**
 * @apiDefine SizeHeader
 *
 * @apiExample {String} size : Parametro url o header
 *    size=[100|200|400|600|800]
 */

func getSizeParam(c *gin.Context) int {
	headerSize, ok := c.GetQuery("size")
	if !ok {
		headerSize = c.GetHeader("size")
	}

	return normalizeParamSize(headerSize)
}

// normalizeParamSize retorna el tamaño a partir del header
func normalizeParamSize(sizeHeader string) int {
	switch sizeHeader {
	case "100":
		return 100
	case "200":
		return 200
	case "400":
		return 400
	case "600":
		return 600
	case "800":
		return 800
	}
	return 200
}
