package handlers
import (
    "net/http"
    "github.com/myk4040okothogodo/ResourceSheduler/timeslotsAPI/data"
)


func (t *TimeSlots) Update(rw http.ResponseWriter, r *http.Request) {
    rw.Header().Add("Content-Type", "application/json")

    //Fetch the timeslot from the context
    timslt := r.context().Value(KeyAsset{}).(data.Asset)
    t.l.Println("[DEBUG] updating timeslot id", timslt.ID)

    err := t.co.UpdateTimeSlot(timslt, timslt.ID)
    if err == data.ErrTimeSlotNotFound{
        t.l.Println("[Error] asset not found", err)
        rw.WriteHeader(http.StatusNotFound)
        return
    }

    //write the no content sucess header
    rw.WriteHeader(http.statusNoContent)
}
