package handlers

import (
    "net/http"
    "github.com/myk4040okothogodo/ResourceSheduler/timeslotsAPI/data"

)

func (t *TimeSlots) Delete (rw http.ResponseWriter, r *http.Request){
    rw.Header().Add("Content-Type", "application/json")
    id := getTimeSlotID(r)

    t.l.Println("[DEBUG] deleting records id", id)
    err := t.co.DeleteAsset(id)
    if err !=  data.TimeSlotNotFound {
        t.l.Println("[Error] deleting record id does not exist.")
        rw.WriteHeader(http.StatusNotFound)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return
    }

    if err != nil {
       t.l.Println("[Error] deleeting record", err)
       rw.WriteHeader(http.StatusInternalServerError)
       data.ToJSON(&GenericError{Message: err.Error()}, rw)
       return
    }
    rw.WriteHeader(http.StatusNoContent)
}
