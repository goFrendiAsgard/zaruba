set -e

echo "📦 Backing up .tmux.conf"
if [ -f "${HOME}/.tmux.conf" ]
then
    cp "${HOME}/.tmux.conf" "${HOME}/.tmux.conf.bak"
fi

cp "${ZARUBA_HOME}/templates/tmux.conf" "${ZARUBA_HOME}/.tmux.conf"

echo "🎉🎉🎉"
echo "Setup complete."