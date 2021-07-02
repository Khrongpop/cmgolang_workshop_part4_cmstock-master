package backup

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

type todoResponse struct {
	Status string `json:"status"`
	Result todo   `json:"result"`
}

type todosResponse struct {
	Status string `json:"status"`
	Result []todo `json:"result"`
}

type todo struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

const (
	dbHost   = "localhost"
	port     = 5432
	username = "root"
	password = "root"
	dbName   = "todoapp"
)

var db *gorm.DB
var err error

func main() {

	databaseURL := fmt.Sprintf("host=%v user=%v dbname=%v sslmode=disable password=%v", dbHost, username, dbName, password)
	fmt.Println(databaseURL)

	conn, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		fmt.Print(err)
	}
	db = conn
	defer db.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! ")
	})

	todo := e.Group("/auth")
	todo.GET("", getTodos)
	todo.POST("", createTodo)
	todo.PATCH("/:id", updateTodo)
	todo.DELETE("/:id", deleteTodo)

	e.Logger.Fatal(e.Start(":1323"))
}

func getTodos(c echo.Context) error {
	todos := []todo{}
	if err := db.Find(&todos).Error; gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusOK, `booking`)
	}

	response := todosResponse{
		Result: todos,
		Status: "suscess",
	}

	return c.JSON(http.StatusOK, response)
}

func createTodo(c echo.Context) error {
	todo := todo{}
	name := c.FormValue("name")
	todo.Name = name
	db.Create(&todo)

	response := todoResponse{
		Result: todo,
		Status: "suscess",
	}

	return c.JSON(http.StatusOK, response)
}

func updateTodo(c echo.Context) error {
	todo := todo{}
	db.First(&todo, c.Param("id"))
	name := c.FormValue("name")

	todo.Name = name
	db.Model(&todo).Update("Name", name)

	response := todoResponse{
		Result: todo,
		Status: "suscess",
	}

	return c.JSON(http.StatusOK, response)
}

func deleteTodo(c echo.Context) error {
	todo := todo{}
	name := c.FormValue("name")
	todo.Name = name
	db.Create(&todo)

	response := todoResponse{
		Result: todo,
		Status: "suscess",
	}

	return c.JSON(http.StatusOK, response)
}
