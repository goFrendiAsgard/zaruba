set -e

echo "🔽 Downloading and installing pyenv."
curl https://pyenv.run | bash

echo "🔽 Update init script."
SCRIPT="$(cat "${ZARUBA_HOME}/setup/templates/bash/pyenv.sh")"
echo "${SCRIPT}" >> "${ZARUBA_HOME}/init.sh"
. "${ZARUBA_HOME}/init.sh"

echo "🔽 Install python 3.9.0."
pyenv install 3.9.0

echo "🔽 Set pyenv 3.9.0 as default."
pyenv global 3.9.0

echo "🔽 Install pipenv."
pip install pipenv

echo "🎉🎉🎉"
echo "Pyenv installed, make sure to update your shell configuration"