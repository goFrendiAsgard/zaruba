set -e

echo "🔽 Downloading and installing gvm."
curl -o- https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash

echo "🔽 Update init script."
SCRIPT="$(cat "${ZARUBA_HOME}/templates/bash/gvm.sh")"
echo "${SCRIPT}" >> "${ZARUBA_HOME}/init.sh"
. "${ZARUBA_HOME}/init.sh"

echo "🎉🎉🎉"
echo "Gvm installed, make sure to update your shell configuration"