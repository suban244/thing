import grpc
from concurrent import futures

import graderrequest_pb2
import graderrequest_pb2_grpc


class RequestService(graderrequest_pb2_grpc.GraderRequestService):
    def GradeFile(self, req, context):
        print("[SERVER] Got request to grade a file")
        print(req)

        reply = graderrequest_pb2.Status(statusCode=200)
        return reply


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=4))
    graderrequest_pb2_grpc.add_GraderRequestServiceServicer_to_server(
        RequestService(), server
    )
    server.add_insecure_port("localhost:4000")

    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
