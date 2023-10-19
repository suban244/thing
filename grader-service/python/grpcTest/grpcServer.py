import grpc
from concurrent import futures

import request_pb2
import request_pb2_grpc


class RequestService(request_pb2_grpc.GradeService):
    def GradeFile(self, req, context):
        print("[SERVER] Got request to grade a file")
        print(req)

        reply = request_pb2.Status(statusCode=200)
        return reply


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=4))
    request_pb2_grpc.add_GradeServiceServicer_to_server(RequestService(), server)
    server.add_insecure_port("localhost:4000")

    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
