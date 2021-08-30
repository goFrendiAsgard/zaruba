set -e
SETUP_PATH="$(dirname $(realpath "${0}"))"

echo "🔽 Downloading spark."
wget https://downloads.apache.org/spark/spark-3.1.2/spark-3.1.2-bin-hadoop3.2.tgz

echo "🔽 Extracting spark."
tar xvf ./spark-3.1.2-bin-hadoop3.2.tgz -C "${HOME}"
mv "${HOME}/spark-3.1.2-bin-hadoop3.2" "${HOME}/spark"

echo "🔽 Update init script."
SCRIPT="$(cat "${SETUP_PATH}/../templates/bash/spark.sh")"
echo "${SCRIPT}" >> "${SETUP_PATH}/../init.sh"
. "${SETUP_PATH}/../init.sh"

echo "🎉🎉🎉"
echo "Spark installed"
