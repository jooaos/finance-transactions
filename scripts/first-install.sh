#!/bin/sh

ENV=.env
ENV_EXAMPLE=.env.example
ENV_TESTING=.env.testing
ENV_INTEGRAION_TEST=./tests/integration/.env

if [ ! -f ${ENV} ]; then
    cp ${ENV_EXAMPLE} ${ENV}
fi

if [ ! -f ${ENV_INTEGRAION_TEST} ]; then
    cp ${ENV_TESTING} ${ENV_INTEGRAION_TEST}
fi