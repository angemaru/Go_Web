package main

import (
	"log"
	"rest/cmd/routers"
	"rest/internal/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	// instances
	db := []domain.Product{
		{Id: 1, Name: "arepa", Is_published: false, Price: 350, Code_value: "ana01", Expiration: "01/03/2030", Quantity: 30},
		{Id: 2, Name: "pan", Is_published: true, Price: 350, Code_value: "ana09", Expiration: "01/03/2023", Quantity: 14},
		{Id: 3, Name: "huevo", Is_published: false, Price: 350, Code_value: "ana07", Expiration: "01/03/2023", Quantity: 10},
	}
	// rp := website.NewRepository(&db, 3)
	// sv := website.NewService(rp)

	// app
	// websites, _ := sv.Get()
	// fmt.Println("- websites:", websites)

	// ws, err := sv.Create("https://www.music.com", "amazon", "music", false)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	// fmt.Println("- website created:", ws)

	// exposure [std io]
	en := gin.Default()
	rt := routers.NewRouter(en, &db)
	rt.SetRoutes()

	if err := en.Run(); err != nil {
		log.Fatal(err)
	}
}
