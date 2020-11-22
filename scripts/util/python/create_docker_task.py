import sys
import os

# USAGE
# python create_docker_task.py <image> <container> <task>

def create_docker_task(template_path: str, image: str, container: str, task: str) -> str:
    image = image if image != "" else "nginx"
    container = container if container != "" else image
    task = task if task != "" else "run{}".format(container)
    if not os.path.exists("docker"):
        os.makedirs("docker")
    destination_filename = os.path.join("docker", "{}.zaruba.yaml".format(task))
    if os.path.isfile(destination_filename):
        raise Exception("{} already exists".format(destination_filename))
    print(template_path)


if __name__ == "__main__":
    template_path = os.path.join(
        os.path.dirname(os.path.dirname(os.path.dirname(sys.argv[0]))),
        'docker_task_template'
    )
    image = sys.argv[1] if len(sys.argv) > 1 else ""
    container = sys.argv[2] if len(sys.argv) > 2 else ""
    task = sys.argv[3] if len(sys.argv) > 3 else ""
    create_docker_task(template_path, image, container, task)
