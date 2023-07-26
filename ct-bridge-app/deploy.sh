#!/bin/bash
#rm -rf dist
#npm run build
NOW=`date "+%Y%m%d%H%M%S"`
rsync -avz build/* root@47.98.96.123:/www/wwwroot/bridge.ctblock.cn
