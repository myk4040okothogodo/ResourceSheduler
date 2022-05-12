package handlers
import (
    "fmt"
    "context"
    "net/http" 
    "github.com/myk4040okothogodo/ResourceSheduler/timeslotsAPI/data"
)


//listall handlers GET requests and returns all current timeslots

func (t *TimeSlots) ListAll(rw http.ResponseWriter, r *http.Request) {
    t.l.Println("[DEBUG] get all records")
    rw.Header().Add("Content-Type", "application/json")

    timslts, _ := t.co.GetTimeslots()
    err := data.ToJSON(timslts, rw)
    if err != nil {
        //we should never be here but log the error just incase
        t.l.Println("[ERROR] serializing timeslots", err)
    }
}

//ListsSingle handle GET request
func (t *TimeSlots) ListSingle (rw http.ResponseWriter, r *http.Request){
    rw.Header().Add("Content-Type", "application/json")
    id := getTimeSlotID(r)

    t.l.Println("[DEBUG] get record id", id)
    timslts, err := t.co.GetTimeSlotByID(id)

    switch err {
    
    case nil:

    case data.ErrTimeSlotNotFound:
      t.l.Println("[ERROR] fetching timeslots", err)

      rw.WriteHeader(http.StatusNotFound)
      data.ToJSON(&GenericError{Message: err.Error()}, rw)
      return
    default:
      t.l.Println("[ERROR] fetching timeslots", err)
      rw.WriteHeader(http.StatusInternalServerError)
      data.ToJSON(&GenericError{Message: err.Error()}, rw)
      return
    }
    //
    err := data.ToJSON(timslts, rw)
    if err != nil {
        t.l.Println("[ERROR] serializing timeslots ")
    }
}


func (t *TimeSlots) ListOwnerActiveTimeSlots (rw http.ResponseWriter, r *http.Request){
    rw.Header().Add("Content-Type", "application/json")
    id := getAssetOwnerID(r)

    t.l.Println("[DEBUG get record id]", id)
    timslts, err := a.co.GetActiveTimeSlots(id)
  }
