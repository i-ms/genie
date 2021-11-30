package genie

import (
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Validation struct {
	Data   url.Values
	Errors map[string]string
}

// Validator : initializes the validator struct
func (g *Genie) Validator(data url.Values) *Validation {
	return &Validation{
		Data:   data,
		Errors: make(map[string]string),
	}
}

// Valid : returns true if the data is valid else false
func (v *Validation) Valid() bool {
	return len(v.Errors) == 0
}

// AddError : adds a new error to the Validation struct
func (v *Validation) AddError(key, message string) {
	if _, exists := v.Errors[key]; exists {
		v.Errors[key] = message
	}
}

// Has : returns true if
func (v *Validation) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Required : adds error if field is empty
func (v *Validation) Required(r *http.Request, fields ...string) {
	for _, field := range fields {
		value := r.Form.Get(field)
		if strings.TrimSpace(value) == "" {
			v.AddError(field, "This field is required")
		}
	}
}

// Check : add error if ok is false
func (v *Validation) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// IsEmail : check if provided email is valid or not
// add error if email is invalid
func (v *Validation) IsEmail(field, value string) {
	if !govalidator.IsEmail(value) {
		v.AddError(field, "Invalid email address")
	}
}

// IsInt : check if provided value is an integer
func (v *Validation) IsInt(field, value string) {
	_, err := strconv.Atoi(value)
	if err != nil {
		v.AddError(field, "This field must be an integer")
	}
}

// IsFloat : check if provided float is valid or not
func (v *Validation) IsFloat(field, value string) {
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		v.AddError(field, "This field must be a float point number")
	}
}

// IsDateISO : check for valid date format
func (v *Validation) IsDateISO(field, value string) {
	_, err := time.Parse("2006-01-02", value)
	if err != nil {
		v.AddError(field, "This field must be a valid date in format YYYY-MM-DD")
	}
}

// NoSpaces : check if the value has spaces in it
func (v *Validation) NoSpaces(field, value string) {
	if govalidator.HasWhitespace(value) {
		v.AddError(field, "This field must not contain spaces")
	}
}
