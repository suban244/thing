import grpc
from concurrent import futures
from grader_modules.models import GradingReport
import graderrequest_pb2_grpc, graderrequest_pb2

import os

from dotenv import load_dotenv

import psycopg
import grader


def UploadScore(dbParams, fileid: str, result: GradingReport):
    with psycopg.connect(**dbParams, sslmode="require") as conn:
        with conn.cursor() as cur:
            obtainedScore, totalScore = result.getFinalScore()
            cur.execute(
                """
                UPDATE submissions
                SET isgraded = %s,
                    obtainedscore = %s,
                    maxscore = %s,
                    feedback = %s
                WHERE fileid=%s
                returning *;
            """,
                (
                    True,
                    obtainedScore,
                    totalScore,
                    "yay it works",
                    int(fileid),
                ),
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
        result = grader.run_tests(req.fileid)
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

    serve(params)
