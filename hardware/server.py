#!/usr/bin/env python3
import argparse
import time
from concurrent import futures

import grpc

import service_pb2
import service_pb2_grpc
from strip import Strip

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class HardwareServicer(service_pb2_grpc.HardwareCommandServicer):
    def Basic(self, request, context):
        print("Perform basic", request, context, flush=True)
        print("Color", request.color.r, request.color.g, request.color.b, flush=True)
        print("Intensity", request.intensity, flush=True)
        strip.clear()
        strip.solid_color(
            color={"r": int(request.color.r), "g": int(request.color.g), "b": int(request.color.b)},
            intensity=request.intensity
        )
        return service_pb2.BasicResponse(result=True)

    def Fade(self, request, context):
        print("Perform Fade", request, context, flush=True)
        print("Start color", request.startColor.r, request.startColor.g, request.startColor.b, flush=True)
        print("End color", request.endColor.r, request.endColor.g, request.endColor.b, flush=True)
        print("Start intensity", request.startIntensity, flush=True)
        print("End intensity", request.endIntensity, flush=True)
        print("Duration", request.duration, flush=True)
        strip.clear()
        result = strip.fade(
            color1={"r": int(request.startColor.r), "g": int(request.startColor.g), "b": int(request.startColor.b)},
            color2={"r": int(request.endColor.r), "g": int(request.endColor.g), "b": int(request.endColor.b)},
            intensity1=request.startIntensity,
            intensity2=request.endIntensity,
            duration_ms=int(request.duration)
        )
        return service_pb2.FadeResponse(result=result)

    def Rainbow(self, request, context):
        print("Perform Rainbow", request, context, flush=True)
        print("Cycles", request.cycles, flush=True)
        strip.clear()
        strip.rainbow(cycles=request.cycles)
        return service_pb2.RainbowResponse(result=True)

    def Clear(self, request, context):
        print("Perform Clear", request, context, flush=True)
        strip.clear()
        return service_pb2.ClearResponse(result=True)

    def Test(self, request, context):
        print("Perform Test", request, context, flush=True)
        strip.test()
        return service_pb2.TestResponse(result=True)


if __name__ == '__main__':
    # Process arguments
    parser = argparse.ArgumentParser()
    parser.add_argument('-t', '--test', action='store_true', help='runs test')
    parser.add_argument('-p', '--port', action='store_const', const=50051)
    args = parser.parse_args()

    # Create server
    host = '[::]:%s' % parser.parse_args(['--port']).port
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=2))
    server.add_insecure_port(host)
    service_pb2_grpc.add_HardwareCommandServicer_to_server(HardwareServicer(), server)

    # Create NeoPixel object with appropriate configuration
    strip = Strip()

    if args.test:
        print('Running test', flush=True)
        strip.test()
        exit(0)

    try:
        server.start()
        print('Running hardware server on %s' % host, flush=True)
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        strip.clear()
        server.stop(0)
