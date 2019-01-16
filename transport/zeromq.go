package transport

import (
	log "github.com/Sirupsen/logrus"
	flowmessage "github.com/cloudflare/goflow/pb"
	proto "github.com/golang/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
)

type ZeroMQState struct {
	publisher *zmq.Socket
	topic     string
}

func StartZeroMQTransport(uri string, topic string) *ZeroMQState {
	publisher, _ := zmq.NewSocket(zmq.PUB)
	// defer publisher.Close()

	log.Infof("Listening on '%s' topic='%s' Type=ZeroMQ", uri, topic)
	publisher.Bind(uri)
	state := ZeroMQState{
		publisher: publisher,
		topic:     topic,
	}
	return &state
}
func (s ZeroMQState) SendZeroMQFlowMessage(flowMessage *flowmessage.FlowMessage) {
	b, _ := proto.Marshal(flowMessage)
	s.publisher.SendMessage(s.topic, b)
}
