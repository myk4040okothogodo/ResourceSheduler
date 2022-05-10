package data

import (
    "fmt"
    "time"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "github.com/myk4040okothogodo/ResourceSheduler/assetsAPI/data"
)
// ErrTimeSlot is an error raised when then timeslot has not been found
var ErrTimeSlotNotFound = fmt.Errorf("TimeSlot not found")
// ErrTimeSlotNotCreated is an error raised when a timeslot is not successfully created
var ErrTimeSlotNotCreated = fmt.Errorf("TimeSlot not created")
// ErrTimeSLotNotDeleted is an error raised when a timeslot is not successfully deleted
var ErrTimeSlotNotDeleted = fmt.Errorf("TimeSlot not deleted")
//ErrActiveTimeSlotNotFound is an error when an asset user or owner deosnt have an active timeslot
var ErrActiveTimeSlotNotFound = fmt.Errorf("Active timeslot not Found")

type TimeSlotState string

const (
    past     TimeSlotState  = "past"
    Locked   TimeSlotState  = "locked"
    Unlocked TimeSlotState  = "unlocked"
) 

type TimeSlot struct {
  gorm.Model
  // The name for this time slot
  // required: true
  // max length: 255
  Name string `json: "name" validate: "required"`
  // start time for this timeslot
  // required : true
  StartTime time.Time `json: "starttime" validate: "required"`
  // end time for this timeslot
  // required : true
  EndTime   time.Time `json: "endTime"   validate: "required"`
  // state  of the time slot that had been booked
  // required
  State    TimeSlotState     `json: "state"     validate: "required"`
  // The creator or user of this timeslot, the booking party
  // required : true
  Owners   User      `json: "owners"    validate: "required"`  
  // payments or job orders attached to this timeslot, a timeslot cannot be booked with being paid for
  // required : trure
  Link   *Payment   `json: "payment" validate: "required"`
}

type Payment struct {
    gorm.Model
    Asset Asset
}

type Payment struct {
    gorm.Model
    //the id for this payment
    //required true
    ID int `json: "id" validate: "required"`
    // is related to which timeslot
    // required : true
    timeslot *TimeSlot   `json: "timeslot", validate: "required"`
    // it's related to what asset 
    // required : true
    asset   *Asset    `json: "asset", validate: "required"`
    // its related to which owner
    // required : true
    owner  *User      `json: "owner",  validate: "required"`
}


//Timeslots defines a slice of Timeplot
type Timeslots get []*TimeSlot

//GetTimeSlots return all timeslots from the database
func (c Connecting)GetTimeSlots() (TimeSlots, error) {
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
       return ErrTimeSlotNotFound 
    }
    c.DB.Delete(&timeslot)

    return nil
}


func (c Connecting) getActiveTimeSlots(id int) (TimeSlots, error){
    var timeslots TimeSlots
    if result := c.DB.Where("? BETWEEN ? AND ?  AND owner.id = ?",time.Now().Format("2006-01-02 15:4:5"), starttime, endtime, id).Find(&timeslots) ; result.Error != nil {
        return nil, ErrActiveTimeSlotNotFound
    }
    fmt.Println(timeslots)
    return timeslots , nil

}




