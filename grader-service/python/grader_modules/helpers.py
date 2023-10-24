import importlib.util
import sys


def load_module(filepath: str, module_name="module.name"):
    try:
        spec = importlib.util.spec_from_file_location(module_name, filepath)
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
