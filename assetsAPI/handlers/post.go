package handlers

import (
    "net/http"
    "github.com/myk4040okothogodo/ResourceSheduler/assetsAPI/data"
)

func (a *Assets) Create(rw http.ResponseWriter, r *http.Request) {
    // fetch the product from the context
    asst := r.Context().Value(KeyAsset{}).(data.Asset)
    a.l.Printf("[DEBUG] Inserting asset: %#v\n", asst)
    a.co.AddAsset(asst)

}
