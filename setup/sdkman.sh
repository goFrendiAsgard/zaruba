set -e

echo "🔽 Downloading and installing sdkman."
curl -s "https://get.sdkman.io" | bash

echo "🔽 Update init script."
SCRIPT="$(cat "${ZARUBA_HOME}/setup/templates/bash/sdkman.sh")"
echo "${SCRIPT}" >> "${ZARUBA_HOME}/init.sh"
. "${ZARUBA_HOME}/init.sh"

echo "🔽 Install java"
sdk install java

echo "🎉🎉🎉"
echo "SDKMan installed, make sure to update your shell configuration"