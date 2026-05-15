package middleware

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

// changelogEntry — запись об изменении сущности.
type changelogEntry struct {
	Timestamp string `json:"timestamp"`
	Method    string `json:"method"`
	Request   any    `json:"request"`
	Success   bool   `json:"success"`
	Error     string `json:"error,omitempty"`
}

// changelogFile — файл журнала событий.
const changelogFile = "changelog.log"

// ChangelogInterceptor логирует все gRPC-вызовы в stdout и в файл changelog.log.
func ChangelogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)

	entry := changelogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Method:    info.FullMethod,
		Request:   req,
		Success:   err == nil,
	}
	if err != nil {
		entry.Error = err.Error()
	}

	// Вывод в stdout
	log.Printf("CHANGELOG: %s | success=%v | req=%+v", info.FullMethod, entry.Success, req)

	// Запись в файл
	writeChangelogEntry(entry)

	return resp, err
}

func writeChangelogEntry(entry changelogEntry) {
	f, ferr := os.OpenFile(changelogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if ferr != nil {
		log.Printf("changelog: cannot open file: %v", ferr)
		return
	}
	defer f.Close()

	data, _ := json.Marshal(entry)
	f.Write(data)
	f.WriteString("\n")
}
