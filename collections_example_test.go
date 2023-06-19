package crunchclient_test

import (
	"context"
	"fmt"
	"testing"

	crunchclient "github.com/verizonconnect/42crunch-client-go"
)

func TestExample_Get_Collections(t *testing.T) {
	client, _ := crunchclient.NewClient("https://vz-main.42crunch.com", crunchclient.WithAPIKey("..."))

	collections, err := client.Collections.GetAll(context.TODO())

	if err != nil {
		panic(err)
	}

	fmt.Println("Collections retrieved ", collections.Count)
}

func TestExample_Get_Collection(t *testing.T) {
	client, _ := crunchclient.NewClient("https://vz-main.42crunch.com", crunchclient.WithAPIKey("..."))

	collection, err := client.Collections.ReadCollection(context.TODO(), "17453f78-cb13-471f-a7ee-2599e1425141")

	if err != nil {
		panic(err)
	}

	fmt.Println("Collection: ", collection.Description.Id)
}

func TestExample_Get_CollectionApis(t *testing.T) {
	client, _ := crunchclient.NewClient("https://vz-main.42crunch.com", crunchclient.WithAPIKey("..."))

	collection, err := client.API.ListApis(context.TODO(), "17453f78-cb13-471f-a7ee-2599e1425141")

	if err != nil {
		panic(err)
	}

	fmt.Println("Api Count: ", collection.Count)
}

func TestExample_Get_Api(t *testing.T) {
	client, _ := crunchclient.NewClient("https://vz-main.42crunch.com", crunchclient.WithAPIKey("aa9a4781-fa22-4cce-8248-ad11fc8eedbe"))

	api, err := client.API.ReadApi(context.TODO(), "e35da92e-1a2d-41c3-a083-4eb4f6487d3c")

	if err != nil {
		panic(err)
	}

	fmt.Println("Api Id: ", api.Description.Id)
}

func TestExample_Get_ApiAssessmentReport(t *testing.T) {
	// var opt ClientOption
	// opt.debug = true
	client, _ := crunchclient.NewClient("https://vz-main.42crunch.com", crunchclient.WithAPIKey("aa9a4781-fa22-4cce-8248-ad11fc8eedbe"))

	apiReport, err := client.API.ReadAssessmentReport(context.TODO(), "7f5160de-634d-4f3e-befd-1f630b101dea") //"e35da92e-1a2d-41c3-a083-4eb4f6487d3c")

	if err != nil {
		panic(err)
	}

	fmt.Println("Api Id: ", apiReport.APIID)
}
