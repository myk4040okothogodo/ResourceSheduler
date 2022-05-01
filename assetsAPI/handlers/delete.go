package handlers

import (
    "net/http"
    "github.com/myk4040okothogodo/ResourceSheduler/assetsAPI/data"
)


func (a *Assets) Delete (rw http.ResponseWriter, r *http.Request){
    rw.Header().Add("Content-Type", "application/json")
    id := getAssetID(r)

    a.l.Println("[DEBUG] deleting record id", id)
    err := a.co.DeleteAsset(id)
    if err == data.ErrAssetNotFound {
        a.l.Println("[Error] deleting record id does not exist")

        rw.WriteHeader(http.StatusNotFound)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return

    }

    if err != nil {
        a.l.Println("[Error] deleting record", err)
        rw.WriteHeader(http.StatusInternalServerError)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return

    }

    rw.WriteHeader(http.StatusNoContent)

}
