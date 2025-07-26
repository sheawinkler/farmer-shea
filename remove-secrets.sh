#!/bin/sh

git filter-branch --force --index-filter \
  'git rm --cached --ignore-unmatch packages/core/src/code_assist/oauth2.ts' \
  --prune-empty --tag-name-filter cat -- --all

