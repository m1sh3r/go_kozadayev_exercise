package service

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	canbusv1 "go-kozadayev-exercise/task9/api/proto"
)

type CanBusService struct {
	canbusv1.UnimplementedCanBusServiceServer
	mu     sync.RWMutex
	alerts []*canbusv1.Alert
}

func NewCanBusService() *CanBusService {
	return &CanBusService{}
}

func (s *CanBusService) TelemetryStream(stream canbusv1.CanBusService_TelemetryStreamServer) error {
	for {
		signal, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		var alertMsg string
		if signal.Name == "speed" && signal.Value > 120 {
			alertMsg = fmt.Sprintf("High speed detected: %.2f", signal.Value)
		} else if signal.Name == "rpm" && signal.Value > 6000 {
			alertMsg = fmt.Sprintf("High RPM detected: %.2f", signal.Value)
		}

		if alertMsg != "" {
			alert := &canbusv1.Alert{
				DeviceId:  signal.DeviceId,
				Message:   alertMsg,
				Timestamp: time.Now().Unix(),
			}
			s.mu.Lock()
			s.alerts = append(s.alerts, alert)
			s.mu.Unlock()

			if err := stream.Send(alert); err != nil {
				return err
			}
		}
	}
}

func (s *CanBusService) GetAlerts(ctx context.Context, req *canbusv1.GetAlertsRequest) (*canbusv1.GetAlertsResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var res []*canbusv1.Alert
	for _, a := range s.alerts {
		if a.Timestamp >= req.StartTime && a.Timestamp <= req.EndTime {
			res = append(res, a)
		}
	}
	return &canbusv1.GetAlertsResponse{Alerts: res}, nil
}
