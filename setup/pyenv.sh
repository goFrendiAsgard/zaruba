set -e

echo "🔽 Downloading and installing pyenv."
curl https://pyenv.run | bash

echo "🔽 Update init script."
SCRIPT="$(cat "${ZARUBA_HOME}/setup/templates/bash/pyenv.sh")"
echo "${SCRIPT}" >> "${ZARUBA_HOME}/init.sh"
. "${ZARUBA_HOME}/init.sh"

echo "🔽 Install python 3.8.0."
pyenv install 3.8.0
pyenv global 3.8.0

echo "🔽 Install pipenv."
pip install pipenv

echo "🎉🎉🎉"
echo "Pyenv installed, make sure to update your shell configuration"