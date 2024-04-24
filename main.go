package main

import (
	"reflect"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Hold struct {
	X     int  `json:"X"`
	Y     int  `json:"Y"`
	Ligth bool `json:"Ligth"`
}
type Board struct {
	gorm.Model
	Title       string `json:"Title"`
	Grade       string `json:"Grade"`
	Description string `json:"Description"`
	Board       []Hold `json:"Board"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("can`t connect with the database")
	}
	db.AutoMigrate(&[]Board{})
	r := gin.Default()
	r.POST("/api/board", func(c *gin.Context) {
		var board Board
		NewBoard := &Board{}
		if err := c.ShouldBind(NewBoard); err != nil {
			panic(err)
		}
		boardCreate := db.Create(&NewBoard)
		if boardCreate.Error != nil {
			panic(boardCreate.Error)
		}
		db.First(&board, NewBoard.ID)
		c.JSON(200, board)
	})
	r.PATCH("/api/board/:id", func(c *gin.Context) {
		id := c.Param("id")
		if reflect.TypeOf(id) != nil {
			c.JSON(401, "invalid ID")
		}
		var board Board
		db.First(&board, id)
		Changes := &Board{}
		db.Save(Changes)
		c.JSON(200, board)
	})
	r.GET("/boards", func(c *gin.Context) {
		var boards []Board
		db.Find(&boards)
		c.JSON(200, boards)
	})
	r.GET("/board/:id", func(c *gin.Context) {
		id := c.Param("id")
		if reflect.TypeOf(id) != nil {
			c.JSON(401, "invalid ID")
		}
		var board Board
		db.First(&board, id)
		c.JSON(200, board)
	})
	r.Run(":8080")

}
