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


def call_function_with_args(func, args_list):
    try:
        result = func(*args_list)
        return result
    except Exception as e:
        return f"Error: {str(e)}"
