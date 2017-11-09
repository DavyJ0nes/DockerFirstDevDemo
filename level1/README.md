# Level 1
```
docker run hello-world
docker run -it --rm ubuntu cat /etc/lsb-release
docker run -it --rm amazonlinux bash

docker run -d --rm --name firefox -e DISPLAY=192.168.2.66:0 -v /tmp/.X11-unix:/tmp/.X11-unix jess/firefox
docker run -d --rm --name vscode -e DISPLAY=192.168.2.66:0 -v /tmp/.X11-unix:/tmp/.X11-unix jess/vscode
```
