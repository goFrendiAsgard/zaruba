set -e
SETUP_PATH="$(dirname $(realpath "${0}"))"

echo "🔽 Downloading and installing sdkman."
curl -s "https://get.sdkman.io" | bash

echo "🔽 Update init script."
SCRIPT="$(cat "${SETUP_PATH}/../templates/bash/sdkman.sh")"
echo "${SCRIPT}" >> "${SETUP_PATH}/../init.sh"
. "${SETUP_PATH}/../init.sh"

echo "🔽 Install java"
sdk install java

echo "🎉🎉🎉"
echo "SDKMan installed, make sure to update your shell configuration"