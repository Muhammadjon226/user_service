package user

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Muhammadjon226/user_service/models"
	"github.com/Muhammadjon226/user_service/pkg/logger"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func (s *Service) Created(ctx context.Context, event cloudevents.Event) error {

	var (
		user models.User
	)

	s.logger.Debug("User create", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &user)
	if err != nil {
		log.Println(err)
	}
	_, err = s.storage.User().CreateUser(&user)

	return nil
}
