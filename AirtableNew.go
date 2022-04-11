package main

import (
	"fmt"

	"github.com/mehanizm/airtable"
)

func main() {

	client := airtable.NewClient("keyOmJMHGYoQpMxYw")

	table := client.GetTable("appQntnFzrheCxlir", "Testing")

	records, err := table.GetRecords().
		FromView("Test").
		ReturnFields("FacebookID", "Message").
		Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	for i := 0; i < len(records.Records); i++ {
		fmt.Print(records.Records[i].ID, ", ", records.Records[i].Fields["Message"], "\n")
	}

	toUpdateRecords := &airtable.Records{
		Records: []*airtable.Record{

			{
				ID:          "reca0PKxNyFnqKN72",
				Fields:      map[string]interface{}{"FacebookID": "Testing", "Message": "Add"},
				CreatedTime: "",
				Deleted:     false,
				Typecast:    false,
			},
		},
	}
	updatedRecords, err := table.UpdateRecords(toUpdateRecords)
	if err != nil {
		// Handle error
		panic(err)

	}
	fmt.Println(updatedRecords)
}
