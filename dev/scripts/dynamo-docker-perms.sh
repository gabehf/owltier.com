#!/bin/bash
# In first time setup, the data folder that dynamo's docker image binds
# to had permissions set incorrectly so dynamo didnt work. This fixes that.
sudo chown $USER ./dev/db/data -R
chmod 775 -R ./dev/db/data