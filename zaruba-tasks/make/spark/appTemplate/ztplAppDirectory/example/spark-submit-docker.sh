echo "Copy example.py to spark master container"
docker cp example.py ztplAppContainerNameSparkMaster:/opt/bitnami/example.py

echo "Perform spark submit"
docker exec -it ztplAppContainerNameSparkMaster \
    spark-submit \
    --master spark://ztplAppContainerNameSparkMaster:7077 \
    /opt/bitnami/example.py

echo "Remove example.py from the container"
docker exec -it ztplAppContainerNameSparkMaster rm -Rf /opt/bitnami/example.py