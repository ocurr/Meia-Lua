#!/bin/bash

antlr -Dlanguage=Go -visitor -no-listener -o parser Lua.g4
