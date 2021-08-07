set -e

echo "🔽 Downloading latest release."
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

echo "🔽 Installing."
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

echo "🎉🎉🎉"
echo "Kubectl installed."