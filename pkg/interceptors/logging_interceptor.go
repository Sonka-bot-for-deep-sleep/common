package interceptors

import (
	"context"
	"fmt"
	"time"

	"github.com/Sonka-bot-for-deep-sleep/common/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log, _ := logger.New()
	requestID := uuid.New().String()

	log.WithOptions(zap.Fields(zap.String("request_id", requestID)))

	log.Info("request", zap.Any("request_data", req),
		zap.String("method", info.FullMethod), zap.Time("time", time.Now()))

	resTimeStart := time.Now()
	resp, err = handler(ctx, req)
	if err != nil {
		log.Error("Failed create response")
		return nil, fmt.Errorf("Failed create response")
	}

	log.Info("response",
		zap.Duration("response_time", time.Since(resTimeStart)), zap.Any("response_data", resp))

	return resp, nil
}
