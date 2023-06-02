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

	collection, err := client.Collections.Get(context.TODO(), "17453f78-cb13-471f-a7ee-2599e1425141")

	if err != nil {
		panic(err)
	}

	fmt.Println("Collection: ", collection.Description.Id)
}

func TestExample_Get_CollectionApis(t *testing.T) {
	client, _ := crunchclient.NewClient("https://vz-main.42crunch.com", crunchclient.WithAPIKey("..."))

	collection, err := client.API.GetByCollection(context.TODO(), "17453f78-cb13-471f-a7ee-2599e1425141")

	if err != nil {
		panic(err)
	}

	fmt.Println("Api Count: ", collection.Count)
}

func TestExample_Get_Api(t *testing.T) {
	client, _ := crunchclient.NewClient("https://vz-main.42crunch.com", crunchclient.WithAPIKey("..."))

	api, err := client.API.Get(context.TODO(), "09d09b75-6415-43f0-a3ba-ead0de498fa2")

	if err != nil {
		panic(err)
	}

	fmt.Println("Api Id: ", api.Description.Id)
}
