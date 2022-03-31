package main

import (
	"reflect"

	"github.com/mehanizm/airtable"

	"fmt"

	fb "github.com/huandu/facebook/v2"
)

type FacebookFeed struct {
	Id string `facebook:",required"` // this field must exist in response.
	// mind the "," before "required".
	FeedFrom       *FacebookFeedFrom   `facebook:"from"` // use customized field name "from".
	FeedFromShares *FacebookFeedShares `facebook:"shares"`
	CreatedTime    string              `facebook:"created_time,required"` // both customized field name and "required" flag.
	Omitted        string              `facebook:"-"`                     // this field is omitted when decoding.
	Message        string              `facebook:"message"`
}

type FacebookFeedFrom struct {
	Name string `json:"name"`                   // the "json" key also works as expected.
	Id   string `facebook:"id" json:"shadowed"` // if both "facebook" and "json" key are set, the "facebook" key is used.
}

type FacebookFeedShares struct {
	Count int `facebook:"count"` // the "json" key also works as expected.
}

func main() {
	client := airtable.NewClient("keyOmJMHGYoQpMxYw")
	table := client.GetTable("appQntnFzrheCxlir", "Testing")
	var feed FacebookFeed

	{

		res1, _ := fb.Get("107704657675864/feed?fields=shares,message,likes,reactions,created_time,from&limit=1", fb.Params{
			"access_token": "EAAHPmz80hhABAHcBQZA00wokbvHFI5RuqlEaMZCZCPzPAWaUPESEKfnd94IAW3ZABRnq20HMyIoZAr2gsWfEDOcmsEjbmZAkV1BBNU6C0a4yPGopMtVvhZCIChrrZBTR2GDKu0Mn5oMSh87IRqSRT1v3sdVsnm4KDGjDCws0muanb3fTuE0QwZBw955IFS3JMdvONpF3egaVZBLFiohiWqOIiqUzk766irj4yZBdQFxBBueryqeShM2fceq",
		})
		fmt.Println("result ID is ", res1.GetField("data"))
		fmt.Println(reflect.TypeOf(res1))
		res1.DecodeField("data.0", &feed) // read latest feed
		fmt.Println("latest post ID is", feed.Id)
		fmt.Println("latest post shares is", feed.FeedFromShares)
		fmt.Println("latest post message is", feed.Message)
		// fmt.Println("latest post story is", feed.Story)

		recordsToSend := &airtable.Records{
			Records: []*airtable.Record{
				{
					// Fields: map[string]interface{}{"Field1": ("New Record1"), "Field2": ("New Record2")},
				},
			},
		}
		receivedRecords, err := table.AddRecords(recordsToSend)
		if err != nil {
			fmt.Println("Error")
		}
		fmt.Println(receivedRecords)
	}
}
