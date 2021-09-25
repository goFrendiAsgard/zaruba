set -e

echo "🔽 Downloading and installing nvm."
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash

echo "🔽 Update init script."
SCRIPT="$(cat "${ZARUBA_HOME}/templates/bash/nvm.sh")"
echo "${SCRIPT}" >> "${ZARUBA_HOME}/init.sh"
. "${ZARUBA_HOME}/init.sh"

echo "🎉🎉🎉"
echo "Nvm installed, make sure to update your shell configuration"