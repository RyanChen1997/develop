# 安装vpn
1. 参考：https://v2raya.org/docs/prologue/quick-start/
2. 参考：https://www.pengtech.net/network/v2rayA_install.html#4-7-%E6%A3%80%E6%9F%A5%E8%AE%BE%E7%BD%AE

## 安装v2ray-core和v2rayA
### 安装v2ray-core
v2ray-core是v2rayA的内核，需要先安装v2ray-core。
1. 首先参考这个文章：https://www.pengtech.net/network/v2rayA_install.html#4-7-%E6%A3%80%E6%9F%A5%E8%AE%BE%E7%BD%AE
2. 安装里面的教程安装v2ray-core
```sh
wget -O /tmp/v2ray-linux-64.zip https://ghproxy.net/https://github.com/v2fly/v2ray-core/releases/download/v5.13.0/v2ray-linux-64.zip

sudo unzip /tmp/v2ray-linux-64.zip -d /usr/local/v2ray-core

sudo mkdir -p /usr/local/share/v2ray/

sudo cp /usr/local/v2ray-core/*dat /usr/local/share/v2ray/
```
### 安装v2ray
参考：https://v2raya.org/docs/prologue/quick-start/
1. 现在自己的电脑上下载v2rayA安装包 https://github.com/v2rayA/v2rayA/releases
2. 使用`rz`上传到服务器
3. `apt install ./xxxx.deb` 安装
4. 修改配置`vim /etc/default/v2raya`
```
# 将V2rayA和v2ray-core关联起来
# 添加配置两行配置
V2RAYA_V2RAY_BIN=/usr/local/v2ray-core/v2ray
V2RAYA_V2RAY_CONFDIR=/usr/local/v2ray-core
```
5. 启动：`systemctl start v2raya`
6. 关闭：`systemctl stop v2raya`
**默认端口号：2017，记得配置安全组，在本机上可以访问。**

## 配置v2rayA
1. 访问`http://localhost:2017`
2. 输入账号密码
3. 订阅url
4. 透明代理/系统代理：启动：不进行分流
5. 透明代理/系统代理实现方式：redirect
6. 测试：`curl www.google.com`