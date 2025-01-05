package controllers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/helper"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
)

type motorcycleController struct{}

func newMotorcycleController() *motorcycleController {
	return &motorcycleController{}
}

func (mc *motorcycleController) CreateMotorcycle(w http.ResponseWriter, r *http.Request) {
	var motorcycle *model.MotorcycleDTO
	if err := helper.ReadJson(w, r, &motorcycle); err != nil {
		helper.ErrorJson(w, err, http.StatusUnprocessableEntity)
		return
	}
	if err := srv.MotorcycleSvc.CreateMotorcycle(motorcycle); err != nil {
		helper.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteJson(w, http.StatusCreated, nil)
}

func (mc *motorcycleController) UpdateMotorcycle(w http.ResponseWriter, r *http.Request) {
	var motorcycle model.MotorcycleDTO
	if err := helper.ReadJson(w, r, &motorcycle); err != nil {
		helper.ErrorJson(w, err, http.StatusUnprocessableEntity)
		return
	}
	if err := srv.MotorcycleSvc.UpdateMotorcycle(&motorcycle); err != nil {
		helper.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteJson(w, http.StatusOK, nil)
}

func (mc *motorcycleController) GetAllMotorcycles(w http.ResponseWriter, r *http.Request) {
	motorcycles, err := srv.MotorcycleSvc.GetAllMotorcycles()
	if err != nil {
		helper.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteJson(w, http.StatusOK, motorcycles)
}

func (ms *motorcycleController) GetMotorcycleById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ErrorJson(w, err, http.StatusBadRequest)
		return
	}
	motorcycle, err := srv.MotorcycleSvc.GetMotorcycleById(int64(id))
	if err != nil {
		helper.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteJson(w, http.StatusOK, motorcycle)
}

func (mc *motorcycleController) GetMotorcyclesByYear(w http.ResponseWriter, r *http.Request) {
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	if _, ok := params["year"]; !ok {
		helper.ErrorJson(w, err)
		return
	}
	year, err := strconv.Atoi(params["year"][0])
	motorcycles, err := srv.MotorcycleSvc.GetMotorcyclesByYear(int64(year))
	if err != nil {
		helper.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteJson(w, http.StatusOK, motorcycles)
}

func (mc *motorcycleController) GetMotorcyclesByModel(w http.ResponseWriter, r *http.Request) {
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	if _, ok := params["model"]; !ok {
		helper.ErrorJson(w, err)
		return
	}
	motorcycles, err := srv.MotorcycleSvc.GetMotorcyclesByModel(params["model"][0])
	if err != nil {
		helper.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteJson(w, http.StatusOK, motorcycles)
}
