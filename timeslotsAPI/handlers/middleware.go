package main

import (
    "context"
    "net/http"
    "github.com/myk4040okothogodo/ResourceSheduler/timeslotsAPI/data"

)


//MiddlewareValidateTimeslot validated the timeslot in the request and calls next if ok
func (t *TimeSlots) MiddlewareValidateTimeSlot(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
        rw.Header().Add("Content-Type", "applications/json")

        timslt := &data.TimeSlot{}

        err := data.FromJSON(tmslt, r.Body)
        if err != nil {
            t.l.Println("[ERROR] deserializing asset", err)
            rw.WriteHeader(http.StatusBadRequest)
            data.ToJSON(&GenericError{Message: err.Error()}, rw)
            return
        }

        // Validate the asset
        errs := t.v.Validate(timslt)
        if len(errs) != 0 {
            t.l.Println("[ERROR] validating asset ", errs)

            // return the validation messages as an array
            rw.WriteHeader(http.StatusUnprocessableEntity)
            data.ToJSON(&ValidationError{Message: errs.Errors()}, rw)
            return
        }

        ctx := context.WithValue(r.Context(), KeyTimeSlot{}, timslt)
        r = r.WithContext(ctx)

        //call the next hander, which can be another middleware in the in the chain or the final handler
        next.ServeHTTP(rw, r)
    })

}
