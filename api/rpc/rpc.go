package rpc

import (
	"context"
	"github.com/james-woo/wakeus/api/models"
	"google.golang.org/grpc"
	"log"
	"os"
)

var PerformBasic = func(ctx context.Context, color models.Color, intensity float64) (bool, error) {
	conn := createServiceConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("%s\n", err)
		}
	}()
	client := NewHardwareCommandClient(conn)
	response, err := client.Basic(
		ctx,
		&BasicRequest{
			Color: &Color{
				R: int32(color.R),
				G: int32(color.G),
				B: int32(color.B),
			},
			Intensity: float32(intensity),
		},
	)
	if err == nil {
		log.Printf("Successfully performed basic\n")
	} else {
		log.Printf("Basic error: %s\n", err)
	}
	return response.Result, err
}

var PerformFade = func(ctx context.Context, startColor models.Color, endColor models.Color, startIntensity float64, endIntensity float64, duration int) (bool, error) {
	conn := createServiceConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("%s\n", err)
		}
	}()
	client := NewHardwareCommandClient(conn)
	response, err := client.Fade(
		ctx,
		&FadeRequest{
			StartColor: &Color{
				R: int32(startColor.R),
				G: int32(startColor.G),
				B: int32(startColor.B),
			},
			EndColor: &Color{
				R: int32(endColor.R),
				G: int32(endColor.G),
				B: int32(endColor.B),
			},
			StartIntensity: float32(startIntensity),
			EndIntensity: float32(endIntensity),
			Duration: int32(duration),
		},
	)
	if err == nil {
		log.Printf("Successfully performed fade\n")
	} else {
		log.Printf("Fade error: %s\n", err)
	}
	return response.Result, err
}

var PerformRainbow = func(ctx context.Context, cycles int) (bool, error) {
	conn := createServiceConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("%s\n", err)
		}
	}()
	client := NewHardwareCommandClient(conn)
	response, err := client.Rainbow(
		ctx,
		&RainbowRequest{
			Cycles: int32(cycles),
		},
	)
	if err == nil {
		log.Printf("Successfully performed rainbow\n")
	} else {
		log.Printf("Rainbow error: %s\n", err)
	}
	return response.Result, err
}

var PerformClear = func(ctx context.Context) (bool, error) {
	conn := createServiceConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("%s\n", err)
		}
	}()
	client := NewHardwareCommandClient(conn)
	response, err := client.Clear(ctx, &ClearRequest{})
	if err == nil {
		log.Printf("Successfully performed clear\n")
	} else {
		log.Printf("Clear error: %s\n", err)
	}
	return response.Result, err
}

var PerformTest = func(ctx context.Context) (bool, error) {
	conn := createServiceConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("%s\n", err)
		}
	}()
	client := NewHardwareCommandClient(conn)
	response, err := client.Test(ctx, &TestRequest{})
	if err == nil {
		log.Printf("Successfully performed test\n")
	} else {
		log.Printf("Test error: %s\n", err)
	}
	return response.Result, err
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