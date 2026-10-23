package bot

import (
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"

	"github.com/eswar-7116/justlive/internal/input"
)

const botTimeout = 1 * time.Second

type Server struct {
	conn   *net.UDPConn
	states chan State

	mu         sync.Mutex
	client     *net.UDPAddr
	lastAddr   string
	lastSeen   time.Time
	latest     input.Input
	dashCount  uint32
	startCount uint32
}

func Listen(addr string) (*Server, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, err
	}
	log.Printf("bot: listening on udp %s", addr)

	s := &Server{conn: conn, states: make(chan State, 1)}
	go s.readLoop()
	go s.writeLoop()
	return s, nil
}

func (s *Server) readLoop() {
	buf := make([]byte, 64*1024)
	for {
		n, addr, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("bot: read: %v", err)
			return
		}

		var a Action
		if err := json.Unmarshal(buf[:n], &a); err != nil {
			log.Printf("bot: bad action: %v", err)
			continue
		}

		s.mu.Lock()
		if s.client == nil || s.client.String() != addr.String() {
			log.Printf("bot: connected %s", addr)
			s.client = addr
			s.latest = input.Input{}
			// New bot starts counting presses from 0
			if addr.String() != s.lastAddr {
				s.dashCount = 0
				s.startCount = 0
			}
			s.lastAddr = addr.String()
		}
		s.lastSeen = time.Now()

		in := a.heldInput()
		// Keep presses until the game reads them
		in.Dash = s.latest.Dash || a.DashCount != s.dashCount
		in.Start = s.latest.Start || a.StartCount != s.startCount
		s.dashCount = a.DashCount
		s.startCount = a.StartCount
		s.latest = in
		s.mu.Unlock()
	}
}

func (s *Server) writeLoop() {
	for st := range s.states {
		s.mu.Lock()
		client := s.client
		s.mu.Unlock()
		if client == nil {
			continue
		}

		data, err := json.Marshal(st)
		if err != nil {
			log.Printf("bot: encode state: %v", err)
			continue
		}
		if _, err := s.conn.WriteToUDP(data, client); err != nil {
			log.Printf("bot: write: %v", err)
		}
	}
}

func (s *Server) Input() input.Input {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Release all controls when the bot goes silent
	if s.client != nil && time.Since(s.lastSeen) > botTimeout {
		log.Printf("bot: timed out %s", s.client)
		s.client = nil
		s.latest = input.Input{}
	}

	in := s.latest
	s.latest.Dash = false
	s.latest.Start = false
	return in
}

func (s *Server) Send(st State) {
	select {
	case <-s.states:
	default:
	}
	select {
	case s.states <- st:
	default:
	}
}
