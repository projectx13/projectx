#/bin/sh

function test {
  echo "+ $@"
  "$@"
  local status=$?
  if [ $status -ne 0 ]; then
    exit $status
  fi
  return $status
}

GIT_VERSION=`cd ${GOPATH}/src/github.com/projectx13/projectx; git describe --tags`

if [ "$1" == "local" ]
then
  # This will run with local go
  cd $GOPATH
  set -e
  test go build -ldflags="-w -X github.com/projectx13/projectx/util.Version=${GIT_VERSION}" -o /var/tmp/projectx github.com/projectx13/projectx
  test chmod +x /var/tmp/projectx
  test cp -rf /var/tmp/projectx $HOME/.kodi/addons/plugin.video.projectx/resources/bin/linux_x64/
  test cp -rf /var/tmp/projectx $HOME/.kodi/userdata/addon_data/plugin.video.projectx/bin/linux_x64/
elif [ "$1" == "docker" ]
then
  # This will run with docker libtorrent:linux-x64 image
  cd $GOPATH/src/github.com/projectx13/projectx
  test make linux-x64
  test cp -rf build/linux_x64/projectx $HOME/.kodi/addons/plugin.video.projectx/resources/bin/linux_x64/
  test cp -rf build/linux_x64/projectx $HOME/.kodi/userdata/addon_data/plugin.video.projectx/bin/linux_x64/
fi
