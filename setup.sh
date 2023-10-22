# fedora 


if [ "$1" = "go" ]; then
  sudo dnf install go nginx
  go get
  go build

fi

if [ "$1" = "python" ]; then
  sudo dnf install python
  cd grader-service/python
  curl  https://bootstrap.pypa.io/get-pip.py > get-pip.py
  python get-pip.py
  pip install -r requirements.txt

fi

