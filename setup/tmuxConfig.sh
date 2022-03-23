set -e

echo "📦 Backing up .tmux.conf"
if [ -f "${HOME}/.tmux.conf" ]
then
    cp "${HOME}/.tmux.conf" "${HOME}/.tmux.conf.bak"
fi

cp "${ZARUBA_HOME}/setup/templates/tmux.conf" "${HOME}/.tmux.conf"

echo "🎉🎉🎉"
echo "Setup complete."