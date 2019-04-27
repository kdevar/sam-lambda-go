#!/usr/bin/env python

import os

module_name = '{{cookiecutter.repo_location}}/{{cookiecutter.project_name}}'

os.system("go mod init " + module_name)
