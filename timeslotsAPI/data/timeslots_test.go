package data

import "testing"

func TestChecksValidation(t *testing.T){
    t := &TimeSlot{}

    err := t.Validate()

    if err != nil {
        t.Fatal(err)
    }

}
