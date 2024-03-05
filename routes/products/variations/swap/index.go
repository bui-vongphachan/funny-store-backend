package product_variation_swap

type SwapProductVariationBody struct {
	FromProductId string `json:"fromProductId" binding:"required"`
	ToProductId   string `json:"toProductId" binding:"required"`
}

// func SwapProductVariation(db *mongo.Database, r *gin.Engine) {
// 	r.PATCH("/product/variation/swap", func(c *gin.Context) {
// 		result := gin.H{
// 			"status":  400,
// 			"isError": true,
// 			"data":    nil,
// 			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
// 		}

// 		var requstBody SwapProductVariationBody

// 		if err := c.BindJSON(&requstBody); err != nil {
// 			fmt.Println(err.Error())
// 			result["message"] = "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ"
// 			c.JSON(http.StatusBadRequest, result)
// 			return
// 		}

// 		fromProductObjectId, err := primitive.ObjectIDFromHex(requstBody.FromProductId)

// 		if err != nil {
// 			fmt.Println(err.Error())
// 			result["message"] = "FromProductId ບໍ່ຖືກຕ້ອງ"
// 			c.JSON(http.StatusBadRequest, result)
// 			return
// 		}

// 		toProductObjectId, err := primitive.ObjectIDFromHex(requstBody.ToProductId)

// 		if err != nil {
// 			fmt.Println(err.Error())
// 			result["message"] = "ToProductId ບໍ່ຖືກຕ້ອງ"
// 			c.JSON(http.StatusBadRequest, result)
// 			return
// 		}

// 		session, err := db.Client().StartSession()

// 		if err != nil {
// 			panic(err)
// 		}

// 		swapProductVariation.SwapProductVariation(db, &session, &fromProductObjectId, &toProductObjectId)

// 		c.JSON(http.StatusOK, result)
// 	})
// }
