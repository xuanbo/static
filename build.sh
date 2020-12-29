app=static
zip=static.zip

echo "工作路径: $PWD"

# 编译
echo "编译成可执行文件..."
go build -o bin/$app cmd/static/main.go
echo "编译成功 :)"

# dist文件夹不存在，则创建
if [ ! -d "dist" ]; then
    mkdir dist
fi

# 存在打包文件则删除
if [ -d "dist/$zip" ]; then
    rm -rf dist/$zip
fi

# 打包
echo "打包..."
zip -r -q dist/$zip * -x@exclude.lst
echo '打包完成 :)'

# 移除编译后的文件
echo '清理...'
rm -rf bin/$app

echo '完成 :)'
