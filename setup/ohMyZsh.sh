set -e

echo "🔽 Downloading and installing oh-my-zsh"
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
echo "🔽 Downloading and installing zinit"
sh -c "$(curl -fsSL https://git.io/zinit-install)"

if [ -f "${HOME}/.zshrc"]
then
    mv "${HOME}/.zshrc" "${HOME}/.zshrc.backup"
fi

echo "🔽 Update .zshrc"
cp "${ZARUBA_HOME}/setup/templates/zshrc" "${HOME}/.zshrc"

echo "🎉🎉🎉"
echo "Oh-my-zsh installed"