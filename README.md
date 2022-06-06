# douyi_download

[![Go](https://github.com/nanlei2000/douyi_download/actions/workflows/go.yml/badge.svg)](https://github.com/nanlei2000/douyi_download/actions/workflows/go.yml)
[![Release cmd/dydl](https://github.com/nanlei2000/douyi_download/actions/workflows/release.yml/badge.svg)](https://github.com/nanlei2000/douyi_download/actions/workflows/release.yml)

通过分享链接，下载抖音无水印视频的命令行工具

## 免责声明

- 本项目纯属个人爱好创作
- 所有视频的版权始终属于「字节跳动」
- 严禁用于任何商业用途，如果构成侵权概不负责

## 功能

- 下载抖音无水印视频、图集
- 下载微博博文、主页相册原图（有水印）

## 安装

### 1. 获取可执行文件

前往 [release](https://github.com/nanlei2000/douyi_download/releases) 页面下载对应平台压缩包解压即可

#### 源码编译

除了下载已经编译好的可执行文件，你也可以自己编译，需要 go 1.18 版本或以上，如 windows 下：

```
go build -ldflags="-s -w" -o dydl.exe ./cmd/.
```

### 2. 添加执行权限

windows 可忽略

```
chmod +x ./dydl
```

## 使用示例

| 功能                     | 链接类型 | 命令示例                                                                                |
| ------------------------ | ----------- | ------------------------------------------------------------------ |
| 下载单个抖音作品         | 分享链接 | .\dydl.exe "0.79 cNj:/ %%这座城市 https://v.douyin.com/FTdTfDw/ 复制此链接，打开 Dou 音搜索，直接观看视频！"  |
| 下载主页所有抖音作品     | 主页链接 | .\dydl.exe -up https://www.douyin.com/user/MS4wLjABAAAAZimxk0o3KWTEJNNrzwSF3HBjCy4TkS6mpPyHNxEYC2A?relation=1 |
| 下载微博博文原图         | 博文链接 | .\dydl.exe -wb https://weibo.com/2286073303/LvhJiA6Fh                                                         |
| 下载微博主页相册所有原图 | 主页链接 | .\dydl.exe -wb -up https://weibo.com/u/2286073303                                                             |

### 微博 cookie

“下载微博主页相册所有原图”功能需要微博 cookie，在运行命令之前执行

```
# windows powershell
$Env:WB_COOKIE = "your_weibo_cookie"

# mac, linux
export WB_COOKIE = "your_weibo_cookie"

```

## 命令行参数

```
   --up, --user_post       下载所有发布的内容 (default: false)
   --wb, --weibo           下载微博图片 (default: false)
   -p value, --path value  文件下载路径 (default: "./video/")
   -v, --verbose           切换 verbose 模式 (default: false)
   --help, -h              show help (default: false)
```
