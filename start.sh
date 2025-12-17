#!/bin/bash

# 🚀 汉字寻宝引擎启动脚本

echo "🚀 汉字寻宝引擎 · MVP演示版"
echo "=================================="

# 检查Go环境
if ! command -v go &> /dev/null; then
    echo "❌ 错误: 未找到Go，请先安装Go 1.18+"
    exit 1
fi

echo "✅ Go环境检查通过"

# 检查端口是否被占用
if lsof -Pi :8080 -sTCP:LISTEN -t >/dev/null ; then
    echo "⚠️  警告: 端口8080已被占用，尝试终止..."
    lsof -ti:8080 | xargs kill -9 2>/dev/null
    sleep 2
fi

# 启动后端服务
echo "🔄 启动后端API服务..."
cd app/hanbao/api

# 后台启动服务
go run hanbao.go -f etc/hanbao-api.yaml > ../logs/api.log 2>&1 &
API_PID=$!

echo "✅ API服务启动中 (PID: $API_PID)"

# 等待服务启动
sleep 3

# 检查服务是否启动成功
if curl -s http://localhost:8080/api/v1/hanbao/session/start > /dev/null 2>&1; then
    echo "✅ API服务启动成功"
else
    echo "❌ API服务启动失败，请检查日志"
    cat ../logs/api.log
    exit 1
fi

cd ../../..

# 启动前端服务器
echo "🔄 启动前端演示页面..."

# 检查Python可用性
if command -v python3 &> /dev/null; then
    cd web
    python3 -m http.server 3000 > ../logs/web.log 2>&1 &
    WEB_PID=$!
    cd ..
    echo "✅ 前端服务启动成功 (PID: $WEB_PID)"
elif command -v python &> /dev/null; then
    cd web
    python -m http.server 3000 > ../logs/web.log 2>&1 &
    WEB_PID=$!
    cd ..
    echo "✅ 前端服务启动成功 (PID: $WEB_PID)"
else
    echo "⚠️  Python不可用，请手动打开web/index.html"
fi

echo ""
echo "🎉 汉字寻宝引擎启动完成！"
echo ""
echo "📍 后端API: http://localhost:8080"
echo "🌐 前端演示: http://localhost:3000"
echo "📊 API文档: http://localhost:8080/api/v1/hanbao"
echo ""
echo "💡 演示流程:"
echo "   1. 打开前端页面"
echo "   2. 输入中文词语开始解锁"
echo "   3. 体验三种解谜关卡"
echo "   4. 查看个人藏宝图"
echo ""
echo "🛑 按 Ctrl+C 停止所有服务"

# 创建日志目录
mkdir -p app/hanbao/logs

# 保存进程ID
echo $API_PID > app/hanbao/api.pid
if [ ! -z "$WEB_PID" ]; then
    echo $WEB_PID > web.pid
fi

# 等待用户中断
trap 'echo ""; echo "🛑 正在停止服务..."; cleanup' INT

cleanup() {
    if [ -f app/hanbao/api.pid ]; then
        kill $(cat app/hanbao/api.pid) 2>/dev/null
        rm app/hanbao/api.pid
    fi
    if [ -f web.pid ]; then
        kill $(cat web.pid) 2>/dev/null
        rm web.pid
    fi
    echo "✅ 所有服务已停止"
    exit 0
}

wait
