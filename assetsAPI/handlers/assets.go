package handlers

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    protos "github.com/myk4040okothogodo/ResourceSheduler/protos/assetsdefinition/protos"
    data "github.com/myk4040okothogodo/ResourceSheduler/assetsAPI/data"


)

// KeyAsset is a key used for the asset object in the context
type KeyAsset  struct{}

// Assets handler for getting and updating assets
//
type Assets struct {
    l *log.Logger
    v *data.Validation
    ac protos.AssetCheckClient
    co  data.Connecting
}

//NewAssets returns  a new asset handler with the given logger
func NewAssets(l *log.Logger, v *data.Validation, ac protos.AssetCheckClient, co data.Connecting) *Assets {
    return &Assets{l, v, ac, co}
}

// ErrInvalidAssetPath is an error message when the asset path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid path, should be /assets/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
    Message string `json:"message"`
}


//ValidationError is a collection of validation error messages
type ValidationError struct {
    Message []string `json:"messages"`
}


//getAssetID returns the asset ID from the url
//panics if cannot convert the id into an integer
//this should never happen as the router ensures that this is a valid number


func getAssetID(r *http.Request) int {
    // parse the asset id from the url
    vars := mux.Vars(r)

    //conver the id into an intenger an return
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        //should never happend
        panic(err)
    }    
    return id

}
