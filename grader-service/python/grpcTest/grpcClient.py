import grpc
import request_pb2_grpc
import request_pb2


def run():
    with grpc.insecure_channel("localhost:4000") as channel:
        stub = request_pb2_grpc.GradeServiceStub(channel)
        print("unary")
        res = stub.GradeFile(request_pb2.File(fileid="1", filename="2.py"))
        print("Got response")
        print(res)


if __name__ == "__main__":
    run()
