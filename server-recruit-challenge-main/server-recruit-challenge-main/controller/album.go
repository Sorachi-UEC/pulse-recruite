package controller

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"


    "github.com/pulse227/server-recruit-challenge-sample/model"
    "github.com/pulse227/server-recruit-challenge-sample/service"
)

type AlbumController struct {
    service service.AlbumService
}

func NewAlbumController(s service.AlbumService) *AlbumController {
    return &AlbumController{service: s}
}

func (c *AlbumController) GetAlbumListHandler(w http.ResponseWriter, r *http.Request) {
    albums, err := c.service.GetAlbumListService(r.Context())
    if err != nil {
        http.Error(w, fmt.Sprintf("failed to get albums: %v", err), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
    json.NewEncoder(w).Encode(albums)
}

func (c *AlbumController) GetAlbumDetailHandler(w http.ResponseWriter, r *http.Request) {
    //idString := r.PathValue("id")
	idString := r.URL.Path[len("/albums/"):]  // URLからID部分を取得
    albumID, err := strconv.Atoi(idString)
    if err != nil {
        http.Error(w, fmt.Sprintf("invalid id: %v", err), http.StatusBadRequest)
        return
    }

    album, err := c.service.GetAlbumService(r.Context(), model.AlbumID(albumID))
    if err != nil {
        http.Error(w, fmt.Sprintf("failed to get album: %v", err), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
    json.NewEncoder(w).Encode(album)
}

func (c *AlbumController) PostAlbumHandler(w http.ResponseWriter, r *http.Request) {
    var album model.Album
    if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
        http.Error(w, fmt.Sprintf("invalid request body: %v", err), http.StatusBadRequest)
        return
    }

    if err := c.service.PostAlbumService(r.Context(), &album); err != nil {
        http.Error(w, fmt.Sprintf("failed to add album: %v", err), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(album)
}

func (c *AlbumController) DeleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
    idString := r.PathValue("id")
    albumID, err := strconv.Atoi(idString)
    if err != nil {
        http.Error(w, fmt.Sprintf("invalid id: %v", err), http.StatusBadRequest)
        return
    }

    if err := c.service.DeleteAlbumService(r.Context(), model.AlbumID(albumID)); err != nil {
        http.Error(w, fmt.Sprintf("failed to delete album: %v", err), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}