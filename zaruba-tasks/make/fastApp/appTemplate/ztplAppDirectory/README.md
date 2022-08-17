# ZtplAppDirectory

ZtplAppDirectory is a microservice-ready monolith application.


# How to run

ZtplAppDirectory is created using [Zaruba](https://github.com/state-alchemists/zaruba).

You can run it by invoking:

```bash
zaruba please startZtplAppDirectory
```

Alternatively, you can also invoke the following script:

```bash
# create virtual environment if not exist
if [ ! -d ./venv ]; then python -m venv ./venv; fi

# activate virtual environment
source venv/bin/activate

# install pip packages
pip install -r requirements.txt

# load environments
source template.env

# run the application
./start.sh
```

# Documentation

Please visit ZtplAppDirectory documentation [here](_docs/README.md).