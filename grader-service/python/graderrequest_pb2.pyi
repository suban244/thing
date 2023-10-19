from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class File(_message.Message):
    __slots__ = ["fileid", "filename"]
    FILEID_FIELD_NUMBER: _ClassVar[int]
    FILENAME_FIELD_NUMBER: _ClassVar[int]
    fileid: str
    filename: str
    def __init__(self, fileid: _Optional[str] = ..., filename: _Optional[str] = ...) -> None: ...

class noParams(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class Status(_message.Message):
    __slots__ = ["statusCode"]
    STATUSCODE_FIELD_NUMBER: _ClassVar[int]
    statusCode: int
    def __init__(self, statusCode: _Optional[int] = ...) -> None: ...
