package handlers

import (
    "context"
    "net/http"
    "github.com/myk4040okothogodo/ResourceSheduler/assetsAPI/data"
  )

//MiddlewareValidateAsset validated the asset in the request and calls next if ok
func (a *Assets) MiddlewareValidateAsset(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
        rw.Header().Add("Content-Type", "application/json")


        asst := &data.Asset{}

        err := data.FromJSON(asst, r.Body)
        if err != nil {
            a.l.Println("[ERROR] deserializing  asset", err)
            rw.WriteHeader(http.StatusBadRequest)
            data.ToJSON(&GenericError{Message: err.Error()}, rw)
            return
        }

        // Validate the asset
        errs := a.v.Validate(asst)
        if len(errs) != 0 {
            a.l.Println("[ERROR] validating asset",errs)

            //return the validation messages as an array
            rw.WriteHeader(http.StatusUnprocessableEntity)
            data.ToJSON(&ValidationError{Message: errs.Errors()}, rw)
            return

        }

        // add the asset to the context
        ctx := context.WithValue(r.Context(), KeyAsset{}, asst)
        r = r.WithContext(ctx)

        //call the next  handler, which can be another middleware in the chain, or te final handler
        next.ServeHTTP(rw, r)
    })

}
