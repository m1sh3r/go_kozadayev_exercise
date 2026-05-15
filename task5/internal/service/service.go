package service

import (
	"context"
	"sync"
	"time"

	smarthomev1 "go-kozadayev-exercise/task5/api/proto"
)

type SmartHomeService struct {
	smarthomev1.UnimplementedSmartHomeServiceServer
	mu          sync.RWMutex
	subscribers map[string][]chan *smarthomev1.Reading
}

func NewSmartHomeService() *SmartHomeService {
	return &SmartHomeService{
		subscribers: make(map[string][]chan *smarthomev1.Reading),
	}
}

func (s *SmartHomeService) AddReading(ctx context.Context, req *smarthomev1.AddReadingRequest) (*smarthomev1.AddReadingResponse, error) {
	reading := &smarthomev1.Reading{
		DeviceId:  req.DeviceId,
		Value:     req.Value,
		Timestamp: time.Now().Unix(),
	}

	s.mu.RLock()
	subs := s.subscribers[req.DeviceId]
	s.mu.RUnlock()

	for _, ch := range subs {
		select {
		case ch <- reading:
		default:
		}
	}

	return &smarthomev1.AddReadingResponse{Success: true}, nil
}

func (s *SmartHomeService) MonitorReadings(req *smarthomev1.MonitorReadingsRequest, stream smarthomev1.SmartHomeService_MonitorReadingsServer) error {
	ch := make(chan *smarthomev1.Reading, 10)

	s.mu.Lock()
	s.subscribers[req.DeviceId] = append(s.subscribers[req.DeviceId], ch)
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		subs := s.subscribers[req.DeviceId]
		for i, c := range subs {
			if c == ch {
				s.subscribers[req.DeviceId] = append(subs[:i], subs[i+1:]...)
				break
			}
		}
		s.mu.Unlock()
	}()

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case reading := <-ch:
			if err := stream.Send(reading); err != nil {
				return err
			}
		}
	}
}
