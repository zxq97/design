package kit

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"github.com/zxq97/design/pkg/mq/kafka"
	"golang.org/x/sync/errgroup"
)

type kafkaServer struct {
	ctx    context.Context
	cancel context.CancelFunc

	topics []string

	group   sarama.ConsumerGroup
	handler sarama.ConsumerGroupHandler
}

func NewKafkaServer(addr, topics []string, name string, handler sarama.ConsumerGroupHandler) Server {
	group, err := kafka.NewConsumerGroup(addr, name)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &kafkaServer{
		ctx:     ctx,
		cancel:  cancel,
		topics:  topics,
		group:   group,
		handler: handler,
	}
}

func (s *kafkaServer) Run() error {
	eg, ctx := errgroup.WithContext(s.ctx)
	eg.Go(func() error {
		var err error
		for {
			select {
			case err = <-s.group.Errors():
				// log
			case err = <-s.ctx.Done():
				return err
			}
		}
	})
	eg.Go(func() error {
		for {
			if err := s.group.Consume(ctx, s.topics, s.handler); err != nil {
				// log

				if errors.Is(err, context.Canceled) {
					return nil
				}
			}
		}
	})
	return eg.Wait()
}

func (s *kafkaServer) Stop() error {
	s.cancel()
	return s.group.Close()
}
