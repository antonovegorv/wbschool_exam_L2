package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

// TelnetClient provides main telnet operations
type TelnetClient interface {
	Connect() error
	Send() error
	Receive() error
	Close() error
	Done() <-chan struct{}
}

// Instance of a telnet.
type basicTelnetClient struct {
	address    string
	connection net.Conn
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	done       chan struct{}
}

// NewTelnetClient returns a telnet client instance
func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan struct{})
	go func() {
		<-signals
		close(done)
	}()
	return &basicTelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
		done:    done,
	}
}

// Connect — connects to the host.
func (btc *basicTelnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", btc.address, btc.timeout)
	if err != nil {
		return err
	}
	btc.connection = conn
	return nil
}

// Send — sends message to the host.
func (btc *basicTelnetClient) Send() error {
	buffer, err := bufio.NewReader(btc.in).ReadBytes('\n')
	switch {
	case endOfTransmissionCheck(err):
		tryToCloseChannel(btc.done)
		return nil
	case err != nil:
		return err
	default:
	}
	_, err = btc.connection.Write(buffer)
	return err
}

// Receive — receives messages from the host.
func (btc *basicTelnetClient) Receive() error {
	buffer, err := bufio.NewReader(btc.connection).ReadBytes('\n')
	switch {
	case endOfTransmissionCheck(err):
		tryToCloseChannel(btc.done)
		return nil
	case err != nil:
		return err
	default:
	}
	_, err = btc.out.Write(buffer)
	return err
}

// Close — closes the connection.
func (btc *basicTelnetClient) Close() error {
	return btc.connection.Close()
}

// Done — signals that we are done.
func (btc *basicTelnetClient) Done() <-chan struct{} {
	return btc.done
}

// Check whether we end transmission or not.
func endOfTransmissionCheck(err error) bool {
	return err == io.EOF
}

func tryToCloseChannel(c chan struct{}) {
	select {
	case <-c:
	default:
		close(c)
	}
}

// india.colorado.edu 13
func main() {
	timeout := flag.String("timeout", "10s", "timeout for a connection")
	flag.Parse()
	if len(flag.Args()) == 0 {
		usage()
		return
	}
	hostPort := flag.Arg(0)

	timeoutDuration, err := time.ParseDuration(*timeout)
	if err != nil {
		usage()
		panic(err.Error())
	}

	telnetClient := NewTelnetClient(hostPort, timeoutDuration, os.Stdin, os.Stdout)
	err = telnetClientConnectAndAttachStandartStreams(telnetClient)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	<-telnetClient.Done()
	telnetClient.Close()
}

func usage() {
	fmt.Printf("Usage: %s [--timeout 10s] hostname:port\n", os.Args[0])
}

func telnetClientConnectAndAttachStandartStreams(tc TelnetClient) error {
	if err := tc.Connect(); err != nil {
		return err
	}

	// Receive messages in goroutine.
	go func() {
		for {
			select {
			case <-tc.Done():
				return
			default:
			}
			tc.Receive()
		}
	}()

	// Send messages in goroutine.
	go func() {
		for {
			select {
			case <-tc.Done():
				return
			default:
			}
			tc.Send()
		}
	}()

	return nil
}
