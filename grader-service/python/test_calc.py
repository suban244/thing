import importlib.util
import sys
import shutil


class GradingResult:
    def __init__(self, obtainedScore: int, maxScore: int, feedback="") -> None:
        self.feedback = feedback
        self.obtainedScore = obtainedScore
        self.maxScore = maxScore

    def __str__(self):
        return (
            f"[SCORE]: {self.obtainedScore}/{self.maxScore}. "
            f"[FEEDBACK]: {self.feedback}"
        )


def call_function_with_args(func, args_list):
    try:
        result = func(*args_list)
        return result
    except Exception as e:
        return f"Error: {str(e)}"


def test_function(func, args_list, expected_output, message: str = "") -> bool:
    try:
        result = call_function_with_args(func, args_list) == expected_output
        return result
    except Exception as e:
        print(f"Error: {str(e)}")
        if message != "":
            print(message)
        result = False
        return result


def test_add(calc) -> bool:
    t1 = test_function(calc.add, [2, 3], 5, "oof ")
    passed = t1
    if passed:
        print("Yay passed add")
    else:
        print("YOu failed subtract")
    return passed


def test_sub(calc) -> bool:
    t1 = test_function(calc.subtract, [4, 2], 2, "oof++")
    passed = t1
    if passed:
        print("Yay passed Subtract")
    else:
        print("You failed subtract")
    return passed


def move_file(filepath: str):
    lastpart = filepath.split("/")[-1]
    target = f"solution/{lastpart}.py"
    shutil.move(filepath, target)
    return target


def load_module(filepath: str):
    moved_file = move_file(filepath)
    try:
        spec = importlib.util.spec_from_file_location("module.name", moved_file)
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


def run_tests(filepath: str) -> GradingResult:
    tests = {test_add: 5, test_sub: 4}
    max_score = sum([x for x in tests.values()])
    obtained_score = 0

    loadedModule = load_module(filepath)
    if loadedModule:
        for k, v in tests.items():
            passed = k(loadedModule)
            obtained_score += v if passed else 0

        print(f"{obtained_score}/{max_score}")
        return GradingResult(obtained_score, max_score, f"{obtained_score}/{max_score}")

    else:
        return GradingResult(
            obtained_score, max_score, "Failed to load module, check file"
        )


if __name__ == "__main__":
    print(run_tests("../../uploaded-files/19"))
