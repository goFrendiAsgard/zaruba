# only used while deploye docker

. ./bash/get_version.sh
VERSION="$(get_version)"
echo "${VERSION}"
echo "${VERSION}" > .version