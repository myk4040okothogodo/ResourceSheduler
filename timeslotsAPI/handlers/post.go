package handlers

import (
    "net/http"
    "github.com/myk4040okothogodo/ResourceSheduler/timeslotsAPI/data"
)

func ( t *TimeSlots) Create(rw http.ResponseWriter, r *http.Request) {
    // fetch the TimeSlot from the context
    timslt = r.Context().Value(KeyTimeSlot{}).(data.TimeSlot)
    t.l.Printf("[DEBUG] Inserting asset: %#v\n", timslt)
    t.co.AddTimeSlot(timslt)
}
