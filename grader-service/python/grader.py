import importlib.util
import sys
import boto3
import os

from dotenv import load_dotenv

from grader_modules.models import GradingReport
from grader_modules.helpers import load_module


def download_file(fileid: str) -> str:
    s3 = boto3.client(
        "s3",
        aws_access_key_id=os.getenv("BACKBLAZE_KEY_ID"),
        aws_secret_access_key=os.getenv("BACKBLAZE_APPLICATION_KEY"),
        endpoint_url=os.getenv("AWS_ENDPOINT_URL"),
    )
    saveto = f"solution/calc-{fileid}.py"
    s3.download_file("auto-grader", fileid, saveto)
    return saveto


def run_tests(fileid: str, assignmentid: str = "test_calc.py") -> GradingReport:
    testing_module = load_module(filepath=assignmentid)
    filename = download_file(fileid)
    module_to_test = load_module(filename)

    try:
        if testing_module is not None:
            tester = testing_module.Tester(module_to_test)
            tester.run()
            return tester.report
        else:
            return GradingReport()

    except Exception as e:
        print(e)
        return GradingReport()


if __name__ == "__main__":
    load_dotenv("../../.env")
    print(run_tests("32").getFinalScore())
