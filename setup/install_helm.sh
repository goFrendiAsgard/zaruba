set -e

echo "🔽 Downloading helm installer."
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3

echo "🔽 Installing."
chmod 700 get_helm.sh
./get_helm.sh

echo "🎉🎉🎉"
echo "Helm installed."