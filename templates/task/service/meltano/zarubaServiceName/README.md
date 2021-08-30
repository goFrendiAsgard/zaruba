# Meltano

You should edit `start.sh` and edit the command.

# Known Issue

* [Installing meltano using pipenv is problematic](https://gitlab.com/meltano/meltano/-/issues/141). Currently meltano should be installed globally (i.e: `pip install meltano`)
* Meltano require some prerequisites like `psycopg2-binary`. If you are using ubuntu/debian, please consider running `~/.zaruba/setup/init.sh` and choose `setup ubuntu` to make sure that all necessary packages are installed.