from typing import List, Tuple, Callable
from functools import reduce


class ExerciseScore:
    def __init__(
        self,
        id: int,
        score: int,
        passed: bool = False,
    ) -> None:
        self.id = id
        self.score = score
        self.passed = passed

    def __str__(self) -> str:
        passed = "Passed" if self.passed else "Failed"

        return f"Ex-{self.id}\t{passed}\t{self.score}"


class GradingReport:
    def __init__(self, alwaysTrue=False) -> None:
        self.alwaysTrue = alwaysTrue  # A special flag that doesn't run the tests
        self.exerciseScores: List[ExerciseScore] = []
        self.count = 0

    def testandScore(self, func: Callable[[], bool], score: int):
        """
        Calls the function and adds the score to the report
        func: A testing function that checks the module and returns true if test passed
        score: how much the testing function is worth
        """
        passed = True
        if self.alwaysTrue:
            passed = True
        else:
            passed = func()

        self.addScore(ExerciseScore(self.count, score, passed))
        self.count += 1
        return passed

    def addScore(self, score: ExerciseScore):
        self.exerciseScores.append(score)

    def getFinalScore(self) -> Tuple[int, int]:
        obtainedScore = reduce(
            lambda val, ele: val + ele.score if ele.passed else 0,
            self.exerciseScores,
            0,
        )
        totalScore = reduce(lambda val, ele: val + ele.score, self.exerciseScores, 0)

        return obtainedScore, totalScore

    def __str__(self):
        return "Exercise\tResult\tScore\n" + "\n".join(
            e.__str__() for e in self.exerciseScores
        )
