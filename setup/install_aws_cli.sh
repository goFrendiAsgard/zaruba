set -e

echo "🔽 Downloading and installing AWS CLI."
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

echo "🎉🎉🎉"
echo "AWS CLI installed, make sure to update your shell configuration"
