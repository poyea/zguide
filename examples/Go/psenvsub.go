//
//  Pubsub envelope subscriber
//

package main

import (
	zmq "github.com/alecthomas/gozmq"
)

func main() {
	context, _ := zmq.NewContext()
	defer context.Close()

	subscriber, _ := context.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5563")
	subscriber.SetSockOptString(zmq.SUBSCRIBE, "B")

	for {
	  subscriber.Recv(0) // discart the envelope
	  content, _ := subscriber.Recv(0)
	  print(string(content) + "\n")
	}
}
