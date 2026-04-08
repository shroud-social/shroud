package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"services/internal/domain/realm/upload"

	"github.com/nats-io/nats.go"
)

const (
	ApiVersion1          = "v1"
	service              = ApiVersion1 + ".service"
	messages             = service + ".message"
	SubjectMessageNew    = messages + ".new"
	SubjectMessageUpdate = messages + ".update"
	SubjectMessageDelete = messages + ".delete"
	uploads              = service + ".uploads"
	SubjectUploadNew     = uploads + ".new"
	SubjectUploadDelete  = uploads + ".delete"
)

func NewUpload(m *nats.Msg) {
	var receipt upload.Receipt
	err := json.Unmarshal(m.Data, &receipt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt)
}

func DeleteUpload(m *nats.Msg) {

}
