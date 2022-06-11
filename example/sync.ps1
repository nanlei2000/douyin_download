$Env:WB_COOKIE = "your_weibo_cookie"

function Run {
  # 苏苏
  .\dydl.exe -wb -up https://weibo.com/u/2286073303
  .\dydl.exe -up https://www.douyin.com/user/MS4wLjABAAAAGTcyBs_MF1p2vK0QB2pUK_N3-huudY2UtA9Shw0o5N8
  .\dydl.exe -up https://www.douyin.com/user/MS4wLjABAAAAfTv1MxjSta_MOctJPPr-7gA6Yg2_1AG6cXKAnScqVVQ_pX5GiA9sMm9vjfrXF9ec
}

# 每小时下载一次
while (1) {
  Run
  Start-Sleep 3600
}