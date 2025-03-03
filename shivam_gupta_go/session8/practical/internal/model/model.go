package model



// 1. Define Structs with Nested Structures:

// • Create a struct Employee with the following fields:

// • ID (int)

// • Name (string)

// • Age (int)

// • Department (string)

// • Salary (float64)

// • Address (nested struct) with:

// • Street (string)

// • City (string)

// • ZipCode (string)

// • Skills (slice of strings)

// • Use struct tags to ensure correct JSON marshalling/unmarshalling.

type Address struct{
	ID int `json:"id"`
	Street string `json:"street"`
	City string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type Person struct{
  ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Department string `json:"department"`
	Salary float64 `json:"salary"`
	Skills []string `json:"skills"`
	Address Address `json:"address"`
}


