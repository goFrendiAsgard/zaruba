
export SPARK_HOME="${HOME}/.spark/spark-{{ .GetValue "setup.sparkVersion" }}-bin-hadoop{{ .GetValue "setup.hadoopVersion" }}"
export PATH="$PATH:$SPARK_HOME/bin"
