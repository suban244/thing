import grpcServer
import os
from dotenv import load_dotenv


if __name__ == "__main__":
    load_dotenv("../../.env")
    params = {
        "dbname": os.getenv("DB_NAME"),
        "user": os.getenv("DB_USER"),
        "password": os.getenv("PASSWORD"),
        "host": os.getenv("HOST"),
        "port": os.getenv("DB_PORT"),
    }
    grpcServer.viewAll(params)
