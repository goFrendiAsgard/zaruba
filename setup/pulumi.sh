set -e

echo "🔽 Downloading and installing Pulumi."
curl -fsSL https://get.pulumi.com | sh

SCRIPT="$(cat "${ZARUBA_HOME}/setup/templates/bash/pulumi.sh")"
echo "${SCRIPT}" >> "${ZARUBA_HOME}/init.sh"
. "${ZARUBA_HOME}/init.sh"

echo "🎉🎉🎉"
echo "Pulumi installed, make sure to update your shell configuration"