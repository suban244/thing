import grpc
from concurrent import futures
import graderrequest_pb2
import graderrequest_pb2_grpc
import os

from dotenv import load_dotenv

import test_calc
import psycopg


def UploadScore(dbParams, fileid: str, result: test_calc.GradingResult):
    with psycopg.connect(**dbParams, sslmode="require") as conn:
        with conn.cursor() as cur:
            cur.execute(
                """
                UPDATE submissions
                SET isgraded = %s,
                    feedback = %s
                WHERE fileid=%s
                returning *;
            """,
                (True, result.feedback, int(fileid)),
            )
            print(cur.fetchone())

            conn.commit()


def viewAll(dbParams):
    with psycopg.connect(**dbParams, sslmode="require") as conn:
        with conn.cursor() as cur:
            cur.execute(
                """
                select * from submissions
                """,
            )
            for r in cur:
                print(r)

            conn.commit()


class RequestService(graderrequest_pb2_grpc.GraderRequestService):
    def __init__(self, dbParams) -> None:
        super().__init__()
        self.dbParams = dbParams

    def GradeFile(self, req, context):
        print("[SERVER] Got request to grade a file")
        print(req)

        # TODO: Launch a thread for this
        # TODO: Download file
        result = test_calc.run_tests("../../uploaded-files/" + req.fileid)
        print(result)
        UploadScore(self.dbParams, req.fileid, result)

        reply = graderrequest_pb2.Status(statusCode=200)
        return reply


def serve(dbParams):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=4))
    graderrequest_pb2_grpc.add_GraderRequestServiceServicer_to_server(
        RequestService(dbParams), server
    )
    server.add_insecure_port("localhost:4000")

    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    load_dotenv("../../.env")
    params = {
        "dbname": os.getenv("DB_NAME"),
        "user": os.getenv("DB_USER"),
        "password": os.getenv("PASSWORD"),
        "host": os.getenv("HOST"),
        "port": os.getenv("DB_PORT"),
    }
    # UploadScore(params, "1", test_calc.GradingResult(0, 1, "test"))
    # viewAll(params)

    serve(params)
