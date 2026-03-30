#!/bin/sh

echo "$CODE" > main.ts

timeout 2s tsx main.ts