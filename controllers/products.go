package controllers

import (
	"encoding/json"

	"github.com/1LCB/sal"
)

type PostProduct struct{
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
}

type ResponseProduct struct{
	Id uint `json:"id"`
	PostProduct
}

func AuthTokenMiddleware(next sal.SalHandlerFunc) sal.SalHandlerFunc {
	return func(c *sal.Ctx) {
		if token := c.Request.Header.Get("Authorization"); token == ""{
			c.Error("Unauthorized", 401)
			return
		}
		next.ServeHTTP(c)
	}
}

func registerProductRoutes() {
	router := sal.NewRouter("/api/v1/products", "Products")

	router.UseMiddleware(AuthTokenMiddleware)
	router.UseHeader("Authorization", true)

	router.GET("", sal.NewResponse([]string{}, 200), func(c *sal.Ctx) {
		c.Json([]string{"Apple", "Banana", "Coconut", "Dragon Fruit", "Eggplant"}, 200)
	})

	router.POST("", PostProduct{}, sal.NewResponse(ResponseProduct{}, 201), func(c *sal.Ctx) {
		var i PostProduct
		if err := json.NewDecoder(c.Request.Body).Decode(&i); err != nil{
			c.Error(err.Error(), 500)
			return
		}

		if i.Price < 0{
			c.Error("price must be greater than 0", 422)
			return
		}

		response := ResponseProduct{
			Id: 1,
			PostProduct: i,
		}
		c.Json(response, 201)
	})
}
