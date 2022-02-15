echo "Copy example.py to spark master container"
docker cp example.py ztplAppContainerNameSparkMaster:/opt/bitnami/example.py

echo "Perform spark submit"
docker exec -it ztplAppContainerNameSparkMaster spark-submit /opt/bitnami/example.py
