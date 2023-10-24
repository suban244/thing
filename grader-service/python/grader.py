import importlib.util
import sys
import boto3
import os

from grader_modules.models import GradingReport
from grader_modules.helpers import load_module


def download_file(fileid: str) -> str:
    s3 = boto3.client("s3")
    saveto = f"solution/calc-{fileid}.py"
    s3.download_file(os.getenv("BACKBLAZE_KEY_NAME"), fileid, saveto)
    return saveto


def load_moduleOld(fileid: str, module_name="module.name"):
    module_file = download_file(fileid)
    try:
        spec = importlib.util.spec_from_file_location(module_name, module_file)
        if spec:
            loadedModule = importlib.util.module_from_spec(spec)
            sys.modules["module.name"] = loadedModule
            if spec.loader:
                spec.loader.exec_module(loadedModule)
                return loadedModule
        return None
    except Exception as e:
        print(e)
        return None


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
    print(run_tests("calc.py").getFinalScore())
