set -e

echo "🔽 Downloading and installing nvm."
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash

echo "🔽 Update init script."
SCRIPT="$(cat "${ZARUBA_HOME}/setup/templates/bash/nvm.sh")"
echo "${SCRIPT}" >> "${ZARUBA_HOME}/init.sh"
. "${ZARUBA_HOME}/init.sh"

echo "🔽 Install node lts and latest."
nvm install --lts
nvm install node

echo "🔽 Set node latest as default."
nvm use node
nvm alias default node

echo "🎉🎉🎉"
echo "Nvm installed, make sure to update your shell configuration"