#!/bin/sh

ENV_TESTING=.env.testing
ENV_INTEGRAION_TEST=./tests/integration/.env

if [ ! -f ${ENV_INTEGRAION_TEST} ]; then
    cp ${ENV_TESTING} ${ENV_INTEGRAION_TEST}
fi

go test ./tests/integration/