package main

import (
	"github.com/mehanizm/airtable"
)

func main() {

	client := airtable.NewClient("keyOmJMHGYoQpMxYw")

	table := client.GetTable("appQntnFzrheCxlir", "Testing")

	records, err := table.GetRecords().
		FromView("view_1").
		WithFilterFormula("AND({Field1}='value_1',NOT({Field2}='value_2'))").
		WithSort(sortQuery1, sortQuery2).
		ReturnFields("Field1", "Field2").
		InStringFormat("Europe/Moscow", "ru").
		Do()
	if err != nil {
		// Handle error
	}
}
