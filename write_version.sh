# only used while deploy docker

. ./bash/get_version.sh
VERSION="$(get_version)"
echo "${VERSION}" > .version