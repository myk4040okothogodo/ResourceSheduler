package data

import (
    "fmt"
    "time"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

type TimeSlot struct {
  gorm.Model
  // the id for this timeslot
  // required: true
  ID int `json:"id" validate: "required"`   // Unique identifier for this specific timeslot
  // The name for this time slot
  // required: true
  // max length: 255
  Name string `json: "name" validate: "required"`
  // start time for this timeslot
  // required : true
  StartTime time.Time `json: "starttime" validate: "required"`
  EndTime   time.Time `json: "endTime"   validate: "required"`
  State    string     `json: "state"     validate: "required"`
  Owners []*User      `json: "owners"    validate: "required"`  
  link   *Payment 
}

type Payment struct {
    gorm.Model
    //the id for this payment
    //required true
    ID int `json: "id" validate: "required"`
    // is related to which timeslot
    // required : true
    timeslot *TimeSlot
    // it's related to what asset 
    // required : true
    asset   *Asset
    // its related to which owner
    // required : true
    owner  *Owner
}


//Timeslots defines a slice of Timeplot
type Timeslots get []*TimeSlot

//GetTimeSlots return all timeslots from the database
func GetTimeslots(c Connecting) (TimeSlots, error) {
    var timeslots TimeSlots
    if result := c.DB.Find(&timeslots); result.Error  != nil {
        return nil, ErrTimeslotsNotFound
    }
    return timeslots, nil
}



func (c Connecting) UpdatetimeSlot(t TimeSlot, i int ) error {
    if result := c.DB.First(&t, i); result.Error != nil {
        result ErrTimeSlotNotFound
    }
    c.DB.Save(&t)
    return nil
}

func (c Connecting) AddTimeSlot(t TimeSlot) error {
    if result := c.DB.Create(&t);  result.Error != nil {
        result ErrTimeSlotNotCreated
    }
    return nil
}


// DeleteTimeslot deletes a Timeslot from  the database
func (c Connecting) DeleteTimeSlot(id int) error {
    var timeslot TimeSlot;
    if result := c.DB.First(&timeslots, id); result.Error != nil {
       return ErrTimeSLotNot 
    }
    c.DB.Delete(&timeslot)

    return nil
}
