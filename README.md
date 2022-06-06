# douyi_download

[![Go](https://github.com/nanlei2000/douyi_download/actions/workflows/go.yml/badge.svg)](https://github.com/nanlei2000/douyi_download/actions/workflows/go.yml)
[![Release cmd/dydl](https://github.com/nanlei2000/douyi_download/actions/workflows/release.yml/badge.svg)](https://github.com/nanlei2000/douyi_download/actions/workflows/release.yml)

下载抖音无水印视频、图片，微博原图的命令行工具

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

## 使用说明
注意，非 windows 下，下述命令中，`.\dydl.exe` 需要替换为 `./dydl`

### 下载单个抖音作品

- 链接类型：分享链接
- 命令示例：`.\dydl.exe "4.10 tRK:/ 怎么泡男孩子啊，多少水温合适啊%%微胖女生 %%rap版呜呼卡点舞 https://v.douyin.com/F4vTT79/ 复制此链接，打开Dou音搜索，直接观看视频！"`

### 下载主页所有抖音作品

- 链接类型：主页链接
- 命令示例： `.\dydl.exe -up https://www.douyin.com/user/MS4wLjABAAAAZimxk0o3KWTEJNNrzwSF3HBjCy4TkS6mpPyHNxEYC2A?relation=1`

### 下载微博博文原图

- 链接类型：博文链接
- 命令示例：`.\dydl.exe -wb https://weibo.com/2286073303/LvhJiA6Fh`

### 下载微博个人主页相册原图

- 链接类型：个人主页链接
- 命令示例：`.\dydl.exe -wb -up https://weibo.com/u/2286073303`
- 注意：需要 `WB_COOKIE` 环境变量

## `WB_COOKIE` 环境变量

微博相关功能可能需要 `WB_COOKIE` 环境变量。登录微博网页版，按 f12 打开调试工具，切换到 network 模块，点击任意 https://weibo.com/ajax/ 路径下请求， 拿到请求头 cookie 字段，在上述命令执行之前，执行以下命令

```
# 替换引号内的 your_weibo_cookie 为真实的 cookie

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
