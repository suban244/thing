For Creating the python stub, the basic format is
`python -m grpc_tools.protoc -I . --python_out=. --pyi_out=. --grpc_python_out=. ./request.proto`

For creating the golang stub

- The `--go_out`, and `--go-grpc_out` refer to where the output should be
- The `-I ..` represents the directory of the proto file
- The request.proto is the actual proto file
- This the the acutal one i used here

`protoc --go_out=./graderrequest/ --go_opt=paths=source_relative --go-grpc_out=./graderrequest/ --go-grpc_opt=paths=source_relative ./graderrequest.proto -I .`

- Here is an old one i used
  `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ../request.proto -I ..`
