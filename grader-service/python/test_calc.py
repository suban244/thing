from grader_modules.models import GradingReport
from grader_modules.testers import test_function


class Tester:
    def __init__(self, module) -> None:
        # No Tests here
        self.module = module
        self.report = GradingReport()

    # this design terrible so what if we just have a variable in grading report that
    # then true returns the full score without any fuss
    def run(self) -> None:
        self.report.testandScore(self.test_add, 5)
        self.report.testandScore(self.test_sub, 4)

    def test_add(self) -> bool:
        t1 = test_function(self.module.add, [2, 3], 5, "oof ")
        passed = t1
        return passed

    def test_sub(self) -> bool:
        t1 = test_function(self.module.subtract, [4, 2], 2, "oof++")
        passed = t1
        return passed
