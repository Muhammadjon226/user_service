package event

import (
	"context"
	"fmt"
	"log"

	// "go_boilerplate/pkg/logger"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/Muhammadjon226/user_service/config"
	"github.com/Muhammadjon226/user_service/pkg/logger"
)
//Kafka ...
type Kafka struct {
	log          logger.Logger
	cfg          config.Config
	consumers    map[string]*Consumer
	publishers   map[string]*Publisher
	saramaConfig *sarama.Config
}
//NewKafka ...
func NewKafka(cfg config.Config, log logger.Logger) (*Kafka, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0

	kafka := &Kafka{
		log:          log,
		cfg:          cfg,
		consumers:    make(map[string]*Consumer),
		publishers:   make(map[string]*Publisher),
		saramaConfig: saramaConfig,
	}

	return kafka, nil
}

// RunConsumers ...
func (r *Kafka) RunConsumers(ctx context.Context) {
	var wg sync.WaitGroup

	for _, consumer := range r.consumers {
		wg.Add(1)
		go func(wg *sync.WaitGroup, c *Consumer) {
			defer wg.Done()

			err := c.cloudEventClient.StartReceiver(context.Background(), c.handler)

			log.Panic("Failed to start consumer", err)
		}(&wg, consumer)
		fmt.Println("Key:", consumer.topic, "=>", "consumer:", consumer)
	}

	wg.Wait()
}
