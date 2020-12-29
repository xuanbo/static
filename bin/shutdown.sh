app=static

lastDir=$(basename "$PWD")
if [ $lastDir = "bin" ]; then
    cd ..
fi
echo "工作路径: $PWD"

# 找到进程，并杀死
pid=`ps -ef | grep $app | grep -v grep | awk '{print $2}'`
if [ -n "$pid" ]; then
   echo "Kill pid:" $pid
   kill $pid
   echo "关闭服务成功 :)"
else
    echo "服务未运行"
fi