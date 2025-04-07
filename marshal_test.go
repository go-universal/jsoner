package jsoner_test

import (
	"testing"

	"github.com/go-universal/jsoner"
	"github.com/stretchr/testify/assert"
)

func TestJsoner(t *testing.T) {
	type Book struct {
		Title string
		ISBN  string `json:"isbn"`
	}

	type Author struct {
		Name      string `json:"name"`
		Family    string `json:"family"`
		Age       int    `json:"age,string,omitempty"`
		IsMariage bool   `json:",string"`
		Books     []Book `json:"author_books"`
		Skills    []string
		Address   map[string]any
		// ignored fields
		PrivateField string `json:"-"`
		ignored      string
	}

	john := Author{
		Name:      "John",
		Family:    "Doe",
		Age:       0,
		IsMariage: false,
		Books: []Book{
			{Title: "Basics Of C", ISBN: "12345"},
			{Title: "Golang", ISBN: "88888"},
		},
		Skills: []string{"Web dev", "System programming", "IOT"},
		Address: map[string]any{
			"state": map[string]string{
				"country": "USA",
				"county":  "NY",
			},
			"city":   "NY city",
			"street": "ST. 23",
			"no":     13,
		},
		PrivateField: "Some private information",
		ignored:      "i'm ignored",
	}

	encoded, err := jsoner.Marshal(
		john,
		"family",
		"Address.state.country",
		"author_books.isbn",
	)
	assert.NoError(t, err)

	expected := `{"Address":{"city":"NY city","no":13,"state":{"county":"NY"},"street":"ST. 23"},"IsMariage":"false","Skills":["Web dev","System programming","IOT"],"author_books":[{"Title":"Basics Of C"},{"Title":"Golang"}],"name":"John"}`
	assert.JSONEq(t, expected, string(encoded))
}
