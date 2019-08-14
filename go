#! /bin/bash
hexo clean
hexo generate
hexo deploy
git add .
git commit -m "Automatic commit by ./go"
git push