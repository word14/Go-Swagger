package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "fmt"
	// "strconv"

	"./database"
	_ "./docs"
	"./middleware"
	"github.com/gin-gonic/gin"
)

// import (
// 	"./interfaces"
// 	"./model"
// )

type Data struct {
	Name string
}

// type Bebas map[string]interface{}

// @title Test Gin Restful Api
// @version 1.0
// @description This is Test Gin Restful Api.
// @termsOfService http://swagger.io/terms/

// @contact.name Yusuf PM Pangaribuan
// @contact.email yusufpangaribuan31@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
func main() {
	// // var name string
	// // name = "Yusuf"
	// // fmt.Println(name)
	// // data := Data{}
	// // data.Name = ""
	// // fmt.Println(data)
	// emp := model.Employee{}
	// emp.Name = "Yusuf"
	// fmt.Println(emp.GetName())

	// emp2 := &emp
	// emp2.Name = "Samuel"
	// fmt.Println(emp2.Name)

	// // angka := 9
	// // // var angka2 *int
	// // angka2 := &angka //& untuk reference
	// // fmt.Println(angka)
	// // angka = 13
	// // fmt.Println(*angka2) //* untuk pointer
	// nama, umur := model.Test()
	// fmt.Printf("%s %d\n", nama, umur)

	// _, err := strconv.Atoi("123")

	// if err != nil {
	// 	fmt.Println("Bukan angka")
	// } else {
	// 	fmt.Println("Angka")
	// }

	app := gin.Default()
	// app.Use(func(context *gin.Context) {
	// 	// context.Next()
	// 	fmt.Println("masok global!!")
	// 	// context.Next()
	// })
	// // app.Use(func(context *gin.Context) {
	// // 	context.Next()
	// // 	fmt.Println("masok global 2!!")
	// // 	// context.Next()
	// // })
	// name := "sam"
	database.Connect()
	// app.Use(middleware.TestMiddleware(name))
	app.Use(middleware.ErrorMiddleware)
	InitRoute(app)
	InitSwagger(app)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	secret := []byte("yusuftampan")
	tokenString, err := token.SignedString(secret)
	fmt.Println("=========================")
	fmt.Println(tokenString, err)
	fmt.Println("=========================")

	parsedToken, err2 := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		fmt.Println(claims["foo"], claims["exp"])
	} else {
		fmt.Println(err2)
	}
	app.Run(":3000")
	// var fuel interfaces.Fuel
	// pertamax := model.Pertamax{}
	// pertamax.Name = "Premium"
	// pertamax.Price = 1000000
	// pertamax.Octane = 90
	// fuel = pertamax

	// model.Car.FillUp(fuel)
}
