# douyi_download

[![Go](https://github.com/nanlei2000/douyi_download/actions/workflows/go.yml/badge.svg)](https://github.com/nanlei2000/douyi_download/actions/workflows/go.yml)
[![Release cmd/dydl](https://github.com/nanlei2000/douyi_download/actions/workflows/release.yml/badge.svg)](https://github.com/nanlei2000/douyi_download/actions/workflows/release.yml)

通过分享链接，下载抖音无水印视频的命令行工具

## 免责声明

- 本项目纯属个人爱好创作
- 所有视频的版权始终属于「字节跳动」
- 严禁用于任何商业用途，如果构成侵权概不负责

## 获取可执行文件

前往 [release](https://github.com/nanlei2000/douyi_download/releases) 页面下载对应平台压缩包解压即可

## 源码编译

除了下载已经编译好的可执行文件，你也可以自己编译，需要 go 1.18 版本或以上

```
go build -ldflags="-s -w" -o dydl.exe ./cmd/.
```

## 使用示例

### 下载单个视频

1. 拷贝分享链接

2. 执行命令

```
.\dydl.exe "4.10 tRK:/ 怎么泡男孩子啊，多少水温合适啊%%微胖女生 %%rap版呜呼卡点舞  https://v.douyin.com/F4vTT79/ 复制此链接，打开Dou音搜索，直接观看视频！"

2022/05/22 04:06:03 文件名： [filename=D:\go_dev\douyi_download\video\一栗小莎子_1235234024\6947610987423911206.mp4]

```

### 下载主页所有作品

1. 拷贝网页版抖音主页链接，如 https://www.douyin.com/user/MS4wLjABAAAAZimxk0o3KWTEJNNrzwSF3HBjCy4TkS6mpPyHNxEYC2A?relation=1

2. 执行命令

```
.\dydl.exe -up https://www.douyin.com/user/MS4wLjABAAAAZimxk0o3KWTEJNNrzwSF3HBjCy4TkS6mpPyHNxEYC2A?relation=1
```

## 命令行参数

```
   --up, --user_post       下载所有发布的视频 (default: false)
   -p value, --path value  文件下载路径 (default: "./video/")
   -v, --verbose           切换 verbose 模式 (default: false)
   --help, -h              show help (default: false)
```
