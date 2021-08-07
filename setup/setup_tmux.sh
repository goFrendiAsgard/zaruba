set -e

SETUP_PATH="$(dirname $(realpath "${0}"))"

echo "📦 Backing up .tmux.conf"
if [ -f "${HOME}/.tmux.conf" ]
then
    cp "${HOME}/.tmux.conf" "${HOME}/.tmux.conf.bak"
fi

cp "${SETUP_PATH}/../templates/tmux.conf" "${HOME}/.tmux.conf"

echo "🎉🎉🎉"
echo "Setup complete."