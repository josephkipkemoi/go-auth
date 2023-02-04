package controllers

import (

)

// CustomEncoder takes the struct input & Request from Endpoint
// validates & returns error if any
func CustomEncoder(i *JackpotMarketInput) error {
	// d := json.NewDecoder(*http.Request.Body)
	// err := d.Decode(i)
	// if err != nil {
	// 	log.Fatalf("JSON DECODE ERROR: %v\n", err)
	// }

	// validate = validator.New()
	// e := validate.Struct(i)
	return nil
}