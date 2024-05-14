package testhelper

import (
	"context"
	"encoding/json"
	"log"
	h "net/http"
	"reflect"

	"github.com/onshape-public/go-client/onshape"
)

var DocumentName = "Test document"
var DocumentDescription = "This is a test document"

// SetupDocument creates an Onshape document
func SetupDocument(ctx context.Context, client *onshape.APIClient, name string) (string, string,
	func() (respStatus int, err error)) {
	docParams := onshape.NewBTDocumentParams()
	docParams.Name = &DocumentName
	docParams.Description = &DocumentDescription
	docParams.SetName(name)
	//docParams.SetIsPublic(true)
	//create document
	var rawResp *h.Response
	var err error
	docInfo, rawResp, err := client.DocumentAPI.CreateDocument(ctx).BTDocumentParams(*docParams).Execute()
	if err != nil || (rawResp != nil && rawResp.StatusCode >= 300) {
		log.Fatal("err: ", err, " -- Response status: ", rawResp)
	}
	return docInfo.GetId(), *docInfo.GetDefaultWorkspace().Id, func() (int, error) {
		_, rawResp, err := client.DocumentAPI.DeleteDocument(ctx, docInfo.GetId()).Execute()
		return rawResp.StatusCode, err
	}
}

// JSONBytesEqual compared two Json byte arrays
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}
