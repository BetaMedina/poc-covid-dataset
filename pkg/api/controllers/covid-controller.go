package controllers

import (
	"dataset-covid/pkg/api/repositories"
	"dataset-covid/pkg/api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const LIMIT = 100

func FindByState(w http.ResponseWriter, r *http.Request) {
	state, _ := mux.Vars(r)["state"]
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	covidRows, total, _ := repositories.GetCovidPerRegion(state, LIMIT, int64(page))
	if covidRows == nil {
		utils.Json(w, http.StatusNotFound, map[string]interface{}{
			"status":  "error",
			"message": "NOT_FOUND",
		})
		return
	}
	utils.Json(w, http.StatusOK, map[string]interface{}{
		"status": "success",
		"rows":   covidRows,
		"total":  total,
		"page":   page,
	})
	return
}

func FindByCity(w http.ResponseWriter, r *http.Request) {
	city, _ := mux.Vars(r)["city"]
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	covidRows, total, _ := repositories.GetCovidPerCity(city, LIMIT, int64(page))
	if covidRows == nil {
		utils.Json(w, http.StatusNotFound, map[string]interface{}{
			"status":  "error",
			"message": "NOT_FOUND",
		})
		return
	}
	utils.Json(w, http.StatusOK, map[string]interface{}{
		"status": "success",
		"rows":   covidRows,
		"total":  total,
		"page":   page,
	})
	return
}
