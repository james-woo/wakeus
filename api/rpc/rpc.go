package rpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
)

var PerformBasic = func(ctx context.Context, color Color, intensity int32) {
	conn := createServiceConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("%s\n", err)
		}
	}()
	client := NewHardwareCommandClient(conn)
	_, err := client.Basic(
		ctx,
		&BasicRequest{
			Color: &color,
			Intensity: intensity,
		},
	)
	if err == nil {
		log.Printf("Successfully performed basic\n")
	} else {
		log.Printf("Basic error: %s\n", err)
	}
}

var PerformFade = func(ctx context.Context, startColor Color, endColor Color, startIntensity int32, endIntensity int32, duration int32) {
	conn := createServiceConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("%s\n", err)
		}
	}()
	client := NewHardwareCommandClient(conn)
	_, err := client.Fade(
		ctx,
		&FadeRequest{
			StartColor: &startColor,
			EndColor: &endColor,
			StartIntensity: startIntensity,
			EndIntensity: endIntensity,
			Duration: duration,
		},
	)
	if err == nil {
		log.Printf("Successfully performed fade\n")
	} else {
		log.Printf("Fade error: %s\n", err)
	}
}

var PerformClear = func(ctx context.Context) {
	conn := createServiceConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("%s\n", err)
		}
	}()
	client := NewHardwareCommandClient(conn)
	_, err := client.Clear(ctx, &ClearRequest{})
	if err == nil {
		log.Printf("Successfully performed clear\n")
	} else {
		log.Printf("Clear error: %s\n", err)
	}
}

var PerformTest = func(ctx context.Context) {
	conn := createServiceConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("%s\n", err)
		}
	}()
	client := NewHardwareCommandClient(conn)
	_, err := client.Test(ctx, &TestRequest{})
	if err == nil {
		log.Printf("Successfully performed test\n")
	} else {
		log.Printf("Test error: %s\n", err)
	}
}

func createServiceConnection() *grpc.ClientConn {
	host := os.Getenv("HARDWARE_SERVICE_HOST")
	if len(host) == 0 {
		host = "localhost:50051"
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		panic(err)
	} else {
		log.Printf("Connected to %s\n", host)
	}
	return conn
}