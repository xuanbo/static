app=static
addr=:12345
log="logs/startup.log"

lastDir=$(basename "$PWD")
if [ $lastDir = "bin" ]; then
    cd ..
fi
echo "工作路径: $PWD"

# logs文件夹不存在，则创建
if [ ! -d "logs" ]; then
    mkdir logs
fi

# project文件夹不存在，则创建
if [ ! -d "project" ]; then
    mkdir project
fi

# 后台运行
echo "后台启动服务..."
nohup ./bin/$app --addr $addr >> $log &

echo "启动成功 :)"