游戏规则详见<https://www.bilibili.com/read/cv11689293?from=search&spm_id_from=333.337.0.0>。目前游戏已实现基本流程，但未实现与帮手相关的效果，因此仅供试玩。

### 运行游戏
首先开启后端服务，进入后端所在的文件夹
```shell
cd Server
```
运行后端服务
```shell
go run main.go
```

前端运行前，需要先下载相应的依赖包：
```shell
cd Client 
npm install
```

待下载完成后，再运行前端服务
```shell
npm run dev
```

默认情况下是两人游戏，在浏览器中打开网页<http://localhost:5173/>两次，作为两位玩家。然后分别点击这两个网页左下角的“连接服务器”按钮，再在任一网页点击“开始游戏”按钮即可开始游戏。
