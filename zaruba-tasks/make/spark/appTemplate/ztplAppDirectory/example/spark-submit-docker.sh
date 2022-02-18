echo "Copy example.py to spark master container"
docker cp example.py ztplAppContainerNameMaster:/opt/bitnami/example.py

echo "Perform spark submit"
docker exec -it ztplAppContainerNameMaster \
    spark-submit \
    --master spark://ztplAppContainerNameMaster:7077 \
    /opt/bitnami/example.py

echo "Remove example.py from the container"
docker exec -it ztplAppContainerNameMaster rm -Rf /opt/bitnami/example.py