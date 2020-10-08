package controllers

import (
	"encoding/json"
	"github.com/Dadard29/go-api-utils/auth"
	"github.com/Dadard29/go-todo/api"
	"github.com/Dadard29/go-todo/managers"
	"github.com/Dadard29/go-todo/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	idParam = "id"
)


func getId(w http.ResponseWriter, r *http.Request) (int, bool) {
	f := 0

	idStr := r.URL.Query().Get(idParam)
	if idStr == "" {
		api.Api.BuildMissingParameter(w)
		return f, false
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusBadRequest, "failed to parse id", w)
		return f, false
	}

	return id, true
}


func getOneTimeDtoBody(w http.ResponseWriter, r *http.Request) (models.OneTimeDto, bool) {
	var f models.OneTimeDto

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusBadRequest, "failed to parse body", w)
		return f, false
	}

	var o models.OneTimeDto
	err = json.Unmarshal(b, &o)
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusBadRequest, "failed to parse json", w)
		return f, false
	}

	return o, true
}

func OneTimePost(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	o, s := getOneTimeDtoBody(w, r)
	if !s {
		return
	}

	t, err := managers.OneTimeManagerCreate(o)
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "failed to create task", w)
		return
	}

	api.Api.BuildJsonResponse(true, "one-time task created", t, w)
}

func OneTimeGet(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	id, s := getId(w, r)
	if !s {
		return
	}

	t, err := managers.OneTimeManagerGet(id)
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "failed to get task", w)
		return
	}

	api.Api.BuildJsonResponse(true, "one-time task retrieved", t, w)
}

func OneTimePut(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	o, s := getOneTimeDtoBody(w, r)
	if !s {
		return
	}

	t, err := managers.OneTimeManagerUpdate(o)
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "failed to create task", w)
		return
	}

	api.Api.BuildJsonResponse(true, "one-time task created", t, w)
}

func OneTimeDelete(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	id, s := getId(w, r)
	if !s {
		return
	}

	t, err := managers.OneTimeManagerDelete(id)
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "failed to delete task", w)
		return
	}

	api.Api.BuildJsonResponse(true, "one-time task deleted", t, w)
}

func OneTimeListGet(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	l, err := managers.OneTimeManagerList()
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "failed to get list", w)
		return
	}

	api.Api.BuildJsonResponse(true, "one-time list retrieved", l, w)
}
