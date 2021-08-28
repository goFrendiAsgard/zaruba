set -e
SETUP_PATH="$(dirname $(realpath "${0}"))"

echo "🔽 Downloading and installing gvm."
curl -o- https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash

echo "🔽 Update init script."
SCRIPT="$(cat "${SETUP_PATH}/../templates/bash/gvm.sh")"
echo "${SCRIPT}" >> "${SETUP_PATH}/../init.sh"
. "${SETUP_PATH}/../init.sh"

echo "🎉🎉🎉"
echo "Gvm installed, make sure to update your shell configuration"