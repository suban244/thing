import calc


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


def test_add() -> bool:
    t1 = test_function(calc.add, [2, 3], 5, "oof ")
    passed = t1
    if passed:
        print("Yay passed add")
    else:
        print("YOu failed subtract")
    return passed


def test_sub() -> bool:
    t1 = test_function(calc.subtract, [4, 2], 2, "oof++")
    passed = t1
    if passed:
        print("Yay passed Subtract")
    else:
        print("You failed subtract")
    return passed


if __name__ == "__main__":
    tests = {test_add: 5, test_sub: 4}

    max_score = 0
    obtained_score = 0

    for k, v in tests.items():
        passed = k()
        max_score += v
        obtained_score += v if passed else 0

    print(f"{obtained_score}/{max_score}")
