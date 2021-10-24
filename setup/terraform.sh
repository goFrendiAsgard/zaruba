set -e

echo "🔽 Downloading and installing Terraform."
curl "https://releases.hashicorp.com/terraform/1.0.9/terraform_1.0.9_linux_amd64.zip" -o "terraform.zip"
unzip terraform.zip

mv terraform "${HOME}/.terraform"
rm terraform.zip

SCRIPT="$(cat "${ZARUBA_HOME}/setup/templates/bash/terraform.sh")"
echo "${SCRIPT}" >> "${ZARUBA_HOME}/init.sh"
. "${ZARUBA_HOME}/init.sh"

echo "🎉🎉🎉"
echo "Terraform installed, make sure to update your shell configuration"