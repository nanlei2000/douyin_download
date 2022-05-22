# douyi_download
通过分享链接，下载抖音无水印视频的命令行工具

# 免责声明
- 本项目纯属个人爱好创作
- 所有视频的版权始终属于「字节跳动」
- 严禁用于任何商业用途，如果构成侵权概不负责
- 部分源码来自 [DouYinBot](https://github.com/lifei6671/DouYinBot),  MIT license

# 编译
```
go build -o dydl.exe ./
```

# 使用示例
1. 拷贝分享链接

2. 执行命令
```
.\dydl.exe "4.10 tRK:/ 怎么泡男孩子啊，多少水温合适啊%%微胖女生 %%rap版呜呼卡点舞  https://v.douyin.com/F4vTT79/ 复制此链接，打开Dou音搜索，直接观看视频！"

2022/05/22 04:06:03 文件名： [filename=D:\go_dev\douyi_download\video\一栗小莎子_1235234024\6947610987423911206.mp4]
```

# 命令行参数
```
-d, --debug             切换 debug 模式 (default: false)
-p value, --path value  文件下载路径 (default: "./video/")
```