package data

import (
    "encoding/json"
    "io"
)

//TOJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encoder(i)
}

//FromJSON deserializes the object from JSON string
//in an io.Reader to the given interface
func FromJSON (i interface{}, r io.reader) error{
    d := json.NewDecoder(r)
    return d.Decode(i)
}