set -e

echo "🔽 Downloading spark."
wget https://downloads.apache.org/spark/spark-3.1.2/spark-3.1.2-bin-hadoop3.2.tgz

echo "🔽 Extracting spark."
tar xvf ./spark-3.1.2-bin-hadoop3.2.tgz -C "${HOME}"
mv "${HOME}/spark-3.1.2-bin-hadoop3.2" "${HOME}/spark"

echo "🔽 Update init script."
SCRIPT="$(cat "${ZARUBA_HOME}/setup/templates/bash/spark.sh")"
echo "${SCRIPT}" >> "${ZARUBA_HOME}/init.sh"
. "${ZARUBA_HOME}/init.sh"

echo "🎉🎉🎉"
echo "Spark installed"
