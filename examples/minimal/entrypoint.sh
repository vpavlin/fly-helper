#!/bin/bash

flyhelper --config-env FLY_HELPER_CONFG_ENV secrets pull

cd data/
python3 -m http.server