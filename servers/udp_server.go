package servers

import (
"net"
"fmt"
LOG "github.com/cihub/seelog"
)


type UdpServer struct {
	Port int
	keepRunning bool
	buffer []byte
	processor BytesProcessor
}

func NewUdpServer(port int, bufferSize int, processor BytesProcessor)* UdpServer {
	s := &UdpServer{
		Port:port,
		keepRunning:true,
		buffer: make([]byte, bufferSize),
		processor:processor,
	}
	return s
}
func (s *UdpServer) processCnn(conn *net.UDPConn) error{

	n, err := conn.Read(s.buffer[0:])
	if err != nil {
		return err
	} else {
		s.processor.Run(s.buffer, n)
	}
	return nil
}
func (s *UdpServer)Run()  {
	protocol := "udp"
	p := fmt.Sprintf(":%d", s.Port)
	udpAddr, err := net.ResolveUDPAddr(protocol, p)
	if err != nil {
		LOG.Error("Invalid Address")
		return
	}

	udpConn, err := net.ListenUDP(protocol, udpAddr)
	if err != nil {
		LOG.Error(err.Error())
		return
	}
	//Keep calling this function
	for ;s.keepRunning == true;{
		s.processCnn(udpConn)
	}

}

func (s *UdpServer)stop()  {
	s.keepRunning = false
}