package data

import (
    "fmt"
    "github.com/go-playground/validator"
)


//ValidationError wraps the validator FieldError so that we do not expose this to the outside code

type ValidationError struct{
    validator.FieldError
}


func (v ValidationError) Error() string{
    return fmt.Sprintf(
        "Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
        v.Namespace(),
        v.Field(),
        v.Tag(),
    )
}


//ValidationErrors is a collection of ValidationErrors
type ValidationErrors []ValidationError

//Error cocnverts the slice into a string slice
func (v ValidationErrors) Errors() []string {
    errs := []string{}
    for _, err := range v {
        errs = append(errs, err.Error())
    }
    return errs
}
