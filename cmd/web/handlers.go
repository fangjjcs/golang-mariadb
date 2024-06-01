package main

import (
	"app/internal/models"
	"encoding/json"
	"net/http"
	"time"
)

type JsonResp struct {
	OK bool `json:"ok"`
}

type MenuParser struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Memo       string `json:"memo"`
	FileString string `json:"fileString"`
}

func (app *application) CreateMenu(w http.ResponseWriter, r *http.Request) {
	var parser MenuParser
	err := json.NewDecoder(r.Body).Decode(&parser)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	var request models.CreateMenu
	request.Name = parser.Name
	request.Type = parser.Type
	request.Memo = parser.Memo
	request.FileString = parser.FileString
	request.CreatedAt = time.Now().Format("2006-01-02")
	request.UpdatedAt = time.Now().Format("2006-01-02")

	err = app.DB.CreateMenu(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) getAllMenu(w http.ResponseWriter, r *http.Request) {
	menu, err := app.DB.GetAllMenu()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.writeJSON(w, http.StatusOK, menu, "menu")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) getOpenMenu(w http.ResponseWriter, r *http.Request) {
	menu, err := app.DB.GetOpenMenu()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.writeJSON(w, http.StatusOK, menu, "menu")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

type OpenParser struct {
	ID         int    `json:"id"`
	CreateUser string `json:"createUser"`
	CreateDept string `json:"createDept"`
	CloseAt    string `json:"closeAt"`
}

func (app *application) OpenMenu(w http.ResponseWriter, r *http.Request) {
	var parser OpenParser
	err := json.NewDecoder(r.Body).Decode(&parser)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	var request models.OpenMenu
	request.ID = parser.ID
	request.CreateUser = parser.CreateUser
	request.CreateDept = parser.CreateDept
	request.CloseAt = parser.CloseAt

	err = app.DB.OpenMenu(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) updateMenu(w http.ResponseWriter, r *http.Request) {
	var menu models.Menu
	err := json.NewDecoder(r.Body).Decode(&menu)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	menu.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	err = app.DB.UpdateMenu(menu)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

type UpdateRating struct {
	ID    int     `json:"id"`
	Score float64 `json:"score"`
}

func (app *application) updateMenuRating(w http.ResponseWriter, r *http.Request) {
	var score UpdateRating
	err := json.NewDecoder(r.Body).Decode(&score)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.DB.UpdateMenuRating(score.ID, score.Score)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) addOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	order.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	err = app.DB.AddOrder(order)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

type OpenMenuID struct {
	OpenMenuID int `json:"openMenuId"`
}

func (app *application) getAllOrder(w http.ResponseWriter, r *http.Request) {
	var openMenuId OpenMenuID
	err := json.NewDecoder(r.Body).Decode(&openMenuId)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	orders, err := app.DB.AllOrder(openMenuId.OpenMenuID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.writeJSON(w, http.StatusOK, orders, "orders")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) updateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	order.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	err = app.DB.UpdateOrder(order)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

type OrderID struct {
	ID int `json:"id"`
}

func (app *application) deleteOrder(w http.ResponseWriter, r *http.Request) {

	var orderId OrderID
	err := json.NewDecoder(r.Body).Decode(&orderId)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.DB.DeleteOrder(orderId.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := JsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}
