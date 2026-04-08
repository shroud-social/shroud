package service

import (
	subjectsV1 "services/internal/api/service/v1"
	"services/internal/comm/pubsub"

	"github.com/nats-io/nats.go"
)

var (
	SubscriptionsCore = map[string]nats.MsgHandler{
		subjectsV1.SubjectUploadNew:    subjectsV1.NewUpload,
		subjectsV1.SubjectUploadDelete: subjectsV1.DeleteUpload,
	}
)

func Subscribe(map[string]nats.MsgHandler) error {
	for s, h := range SubscriptionsCore {
		_, err := pubsub.Connection.Subscribe(s, h)
		if err != nil {
			return err
		}
	}
	return nil
}
