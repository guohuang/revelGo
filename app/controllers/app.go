package controllers

import (
	"encoding/json"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

/*
Index test
*/
func (c App) Index() revel.Result {
	return c.Render()
}

//MyName test
func (c App) MyName() revel.Result {
	nameModel := NameModel{Name: "guo", Email: "guo@abc.com"}
	return c.Render(nameModel)
}

//GetService test
func (c App) GetService() revel.Result {
	nameModel := NameModel{Name: "Justin", Email: "justin@abc.com"}
	nameModel.Age = 12
	return c.RenderJson(nameModel)
}

//
func (c App) SaveUser(data NameModel) revel.Result {
	//nameModel := NameModel{Name:"xxx"}
	var nameModel NameModel
	err := json.NewDecoder(c.Request.Body).Decode(&nameModel)
	if err != nil {

	}
	//nameModel := revel.Binder.Bind(c.Params, "data", reflect.type(data))
	return c.RenderJson(nameModel)
}

type NameModel struct {
	Name  string
	Email string
	Age   int
}
