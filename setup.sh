# fedora 
if [ "$1" = "setup" ]; then
  sudo dnf install nginx tmux git
fi

if [ "$1" = "go" ]; then
  sudo dnf install go
  go get

  go build
  ./thing
fi

if [ "$1" = "python" ]; then
  sudo dnf install python
  cd grader-service/python
  curl  https://bootstrap.pypa.io/get-pip.py > get-pip.py
  python get-pip.py
  pip install -r requirements.txt
  python grpcServer.py
fi

