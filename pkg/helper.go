package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// Error Handling
func errorHandle(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Reverse the order of events
func reverseSlice(input []*cloudformation.StackEvent) []*cloudformation.StackEvent {
	var output []*cloudformation.StackEvent
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output
}

//Get Json data and return StackID
func GetData(s string) {
	if s != "-" {
		data = s
	} else {
		var v map[string]interface{}
		err := json.NewDecoder(os.Stdin).Decode(&v)
		errorHandle(err)
		if v["StackId"] == nil {
			log.Fatal("StackId not found in data.")
		}
		value := fmt.Sprintf("%v", v["StackId"])

		data = value
	}
}

func GetTime(t string) error {
	layout := "2006-01-02 15:04"
	ttime, err := time.Parse(layout, t)
	if err != nil {
		return err
	}
	StartTime = ttime
	return nil
}

func GetTimeout(t int) {

	timeout = time.Duration(t) * time.Minute
}

func GetRegion(r string) {
	region = r
}
