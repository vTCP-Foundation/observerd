package interfaces

import (
	"github.com/vTCP-Foundation/observerd/common/settings"
	"net"
)

type PublicInterface struct {

}

func NewPublicInterface() (i *PublicInterface, err error) {
	i = &PublicInterface{}
	return
}

func (i *PublicInterface) Run() (flow <-chan error) {
	errorsFlow := make(chan error)
	flow = errorsFlow

	listener, err := net.Listen("tcp", settings.Conf.Interfaces.Public.Interface())
	if err != nil {
		errorsFlow <- err
		return
	}

	//noinspection GoUnhandledErrorResult
	defer listener.Close()

	// todo: add logger here

	for {
		conn, err := listener.Accept()
		if err != nil {
			errorsFlow <- err
			return
		}

		go r.handleConnection(conn, globalErrorsFlow)
	}
}

func (i *PublicInterface) handleConnection(conn net.Conn, globalErrorsFlow chan<- error) {
	processError := func(err errors.E) {
		conn.Close()
	}

	message, err := r.receiveData(conn)
	if err != nil {
		processError(err)
		return
	}

	request, e := v0.ParseRequest(message)
	if e != nil {
		processError(e)
		return
	}

	go r.handleRequest(conn, request, globalErrorsFlow)
}
