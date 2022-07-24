package control

import "github.com/filipgorny/workshop/internal/control/command"

type BusCommands struct {
	chCommands  chan *command.Command
	Subscribers []*BusCommandsSubscriber
}

type BusCommandsSubscriber struct {
	Consume func(c *command.Command)
}

func NewBusCommands() *BusCommands {
	bc := &BusCommands{}
	bc.chCommands = make(chan *command.Command)
	return bc
}

func (bc *BusCommands) SendCommand(cmd *command.Command) {
  func send(subscriber *BBusCommandsSubscriber, cmd *command.Command) {
    subscriber.Consume(cmd)
  }

	for _, subscriber := range bc.Subscribers {
		go send(subscriber, cmd)
	}
}

func (bc *BusCommands) CreateSubcriber() *BusCommandsSubscriber {
	subscriber := BusCommandsSubscriber{}
	subscriber.Consume = func(c *command.Command) {

	}

	bc.Subscribers = append(bc.Subscribers, &subscriber)

	return &subscriber
}
