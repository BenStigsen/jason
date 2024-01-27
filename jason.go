// Package jason provides functionality for working with JSON in a simpler way.
package jason

import (
	"encoding/json"
)

type Jason struct {
	Data any
}

// Unmarshals a JSON byte array into *Jason.
// This returns *Jason and any error which can occur during the unmarshalling process.
func Unmarshal(data []byte) (*Jason, error) {
	value := map[string]any{}
	if err := json.Unmarshal(data, &value); err != nil {
		return nil, err
	}
	return &Jason{value}, nil
}

// Unmarshals current Jason into destination (like a struct).
// This returns an error which can occur during the unmarshalling process.
func (j *Jason) Unmarshal(destination any) error {
	data, err := json.Marshal(j.Data.(map[string]any))
	if err != nil {
		return err
	}
	return json.Unmarshal(data, destination)
}

// Marshals current Jason data into JSON.
// Returns the JSON content as a byte array and error.
func (j *Jason) Marshal() ([]byte, error) {
	return json.Marshal(j.Data)
}

// Returns whether the key is valid and non-null.
func (j *Jason) IsValid(keys ...string) bool {
	value, ok := j.Get(keys...)
	return ok && value != nil
}

// Gets the value using key accessors.
// Returns the key value and if it was found.
func (j *Jason) Get(keys ...string) (any, bool) {
	data := j.Data
	for _, key := range keys {
		value, ok := data.(map[string]any)[key]
		if !ok {
			return nil, false
		}
		data = value
	}
	return data, true
}

// Gets the key value as *Jason.
func (j *Jason) GetObject(keys ...string) *Jason {
	data, ok := j.Get(keys...)
	if !ok {
		return nil
	}
	return &Jason{data.(map[string]any)}
}

/*
func (j *Jason) GetByteString(keys ...string) []byte {
	data, ok := j.Get(keys...)
	if !ok {
		return nil
	}
	return []byte(data.(string))
}
*/

// Gets the key value as a string.
func (j *Jason) GetString(keys ...string) string {
	data, ok := j.Get(keys...)
	if !ok {
		return ""
	}
	return data.(string)
}

// Gets the key value as a number.
func (j *Jason) GetNumber(keys ...string) float64 {
	data, ok := j.Get(keys...)
	if !ok {
		return 0
	}
	return data.(float64)
}

// Gets the key value as a boolean.
func (j *Jason) GetBool(keys ...string) bool {
	data, ok := j.Get(keys...)
	if !ok {
		return false
	}
	return data.(bool)
}

// Gets the key value as an array of any.
func (j *Jason) GetArray(keys ...string) []any {
	data, ok := j.Get(keys...)
	if !ok {
		return nil
	}
	return data.([]any)
}

// Gets the key value as an array of *Jason.
func (j *Jason) GetObjectArray(keys ...string) []*Jason {
	elements := j.GetArray(keys...)
	output := make([]*Jason, 0, len(elements))
	for _, element := range elements {
		output = append(output, &Jason{element})
	}
	return output
}

// Gets the key value as an array of strings.
func (j *Jason) GetStringArray(keys ...string) []string {
	elements := j.GetArray(keys...)
	output := make([]string, 0, len(elements))
	for _, element := range elements {
		output = append(output, element.(string))
	}
	return output
}

// Gets the key value as an array of numbers.
func (j *Jason) GetNumberArray(keys ...string) []float64 {
	elements := j.GetArray(keys...)
	output := make([]float64, 0, len(elements))
	for _, element := range elements {
		output = append(output, element.(float64))
	}
	return output
}

// Gets the key value as an array of booleans.
func (j *Jason) GetBoolArray(keys ...string) []bool {
	elements := j.GetArray(keys...)
	output := make([]bool, 0, len(elements))
	for _, element := range elements {
		output = append(output, element.(bool))
	}
	return output
}
