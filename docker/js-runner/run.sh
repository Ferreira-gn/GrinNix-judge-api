#!/bin/sh
set -e

# Escreve o runner (imutável) em tempo de execução
# USER_CODE, FUNCTION_NAME e TEST_CASES chegam via ENV

exec timeout "${TIMEOUT:-5}" node /app/runner.js