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

//Keytimeslot is a kkey used for the TimeSlot asset object in the context
type KeyTimeSlot struct{}

//TimeSlot handler for getting and updating timeslots

type TimeSlots struct {
    l *log.Logger
    v *data.Validation
    co data.Connecting
} 



//NewTimeSlots returns a new TimeSlot handler with the given logger
func NewTimeSlots(l *log.Logger, v *data.validation,co data.Connecting) *TimeSlots {
    return &TimeSlots{l, v, co}

}

//ErrInvalidTimeSlotPath is an error message when the asset path is not valid
var ErrInvalidTimeSlotPath = fmt.Errorf("Invalid path, should be /timeslots/[id]")


// The below Error is a generic Errorr message returned by a server
type GenericError struct {
    Message string `json: "message"`
}


//ValidationError is a collection of validation error messages
type ValidationError struct {
    Message []string `json: "messages"`
}


//getTimeSlotID returns the timeslot ID from the url
//panics if it cannot convert the ID into an integer
func getTimeSlotID(r *http.Request) int {
    // parse the timeslot id from the url
    vars := mux.Vars(r)

    // convert the id into an integer and return
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        panic(err)
    
    }
    return id
}



func getAssetOwnerID(r *http.Request) id {
    // Parse the owner id from the url
    vars := mux.Vars(r)

    //convert the id  into an integer and return
    id, err := strcon.Atoi(vars["id"])
    if err != nil {
        panic(err)
    }
    return id 
}
