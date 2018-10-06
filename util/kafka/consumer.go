package kafka

import (
	"errors"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	LOG "github.com/cihub/seelog"
)

type Consumer struct {
	consumer *cluster.Consumer
	mutex    *sync.Mutex
}

const OffsetNewest = sarama.OffsetNewest
const OffsetOldest = sarama.OffsetOldest

func NewConsumer(brokersList []string, topic string, consumerGroup string, offset int64) (*Consumer, error) {
	var err error
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = false
	config.Consumer.Offsets.Initial = offset
	// init consumer
	topics := []string{topic}
	consumer, err := cluster.NewConsumer(brokersList,
		consumerGroup, topics, config)
	if err != nil {
		return nil, err
	}
	c := &Consumer{
		consumer: consumer,
		mutex:    &sync.Mutex{},
	}
	return c, err
}

func (c *Consumer) GetMessage(delay int, withLock bool) ([]byte, error) {
	if withLock == true {
		c.mutex.Lock()
		defer c.mutex.Unlock()
	}
	var msg *sarama.ConsumerMessage
	select {
	case message, more := <-c.consumer.Messages():
		if more {
			c.consumer.MarkOffset(message, "") // mark message as processed
			msg = message
		}
	case err, more := <-c.consumer.Errors():
		if more {
			LOG.Errorf("Read kafka message failed, msg=%s", err.Error())
			return nil, err
		}
	case <-time.After(time.Second * time.Duration(delay)):
		{
			return nil, errors.New("TimeOut")
		}
	}
	return msg.Value, nil
}
