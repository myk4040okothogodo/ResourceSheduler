package data

import (
    "fmt"
    "github.com/go-playground/validator"
)


//validationError wraps the validators FieldError so we do not expose this to
//the out code

type ValidationError struct {
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

//ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
    errs := []string{}
    for _, err := range v {
        errs = append(errs, err.Error())

    }
    return errs
}

//validation contains
type Validation struct {
    validate *validator.Validate

}


// New Validation creates a new Validation type
func NewValidation() *Validation {
    validate := validator.New()
    validate.RegisterValidation("sku", validateSKU)

    return &Validation{validate}
}

//Validate the item
//for more detail the returned error can  be cast into a validator.ValidationErrors collection
//
//if ve, ok := err.(validator.ValidationErrors); ok {
//                     fmt.Println(ve.Namespace())
//                     fmt.Println(ve.Field())
//                     fmt.Println(ve.StructNamespace())
//                     fmt.Println(ve.Tag())
//                     fmt.Println(ve.ActualTag())
//                     fmt.Println(ve.Kind())
//                     fmt.Println(ve.type())
//                     fmt.Println(ve.Value())
//                     fmt.Println(ve.Param())
//               
//
//}

func (v *Validation) Validate(i interface{}) ValidationErrors {
    errs := v.validate.Struct(i).(validator.ValidationErrors)

    if len(errs) == 0 {
        return nil
    }
    var returnErrs []ValidationError
    for _, err := range errs {
        // cast the FieldError into our ValidationError and append to
        ve := ValidationError{err.(validator.FieldError)}
        returnErrs = append(returnErrs, ve)
    }

    return returnErrs
}



//ValidateSKU
func validateSKU(f1 validator.FieldLevel) bool {
    return true
}
