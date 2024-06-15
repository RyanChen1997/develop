#!/bin/bash

# sudo apt update

# ------------------ docker install ------------------------
read -p "Do you want to install Docker? (yes/no): " install_docker
if [[ $install_docker == "yes" ]]; then
    if dpkg -l | grep -q -E "^ii.*?docker-ce"; then
        echo "docker is installed."
    else
        sudo apt install -y ca-certificates curl
        sudo install -m 0755 -d /etc/apt/keyrings
        sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
        sudo chmod a+r /etc/apt/keyrings/docker.asc
        # Add the repository to Apt sources:
        echo \
          "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
          $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
          sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
        sudo apt update
        sudo apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
        sudo docker run hello-world
    fi
fi

# ----------------- git --------------------------------------
read -p "Do you want to install git? (yes/no): " install_git
if [[ $install_git == "yes" ]]; then
    if command -v git >/dev/null 2>&1; then
        echo "git is installed"
    else
        echo "install git"
        sudo apt install git
    fi

    # git push需要申请git token
    git config --global credential.helper store
fi

# ----------------- singbox -------------------------------------
read -p "Do you want to install singbox? (yes/no): " install_singbox
if [[ $install_singbox == "yes" ]]; then
    sudo curl -fsSL https://sing-box.app/gpg.key -o /etc/apt/keyrings/sagernet.asc
    sudo chmod a+r /etc/apt/keyrings/sagernet.asc
    echo "deb [arch=`dpkg --print-architecture` signed-by=/etc/apt/keyrings/sagernet.asc] https://deb.sagernet.org/ * *" | \
      sudo tee /etc/apt/sources.list.d/sagernet.list > /dev/null
    sudo apt update
    sudo apt install sing-box # or sing-box-beta
    echo 'sing-box 安装成功!'
    echo '1. 访问 部署订阅转换工具，根据官方文档操作，得到config.json'
    echo '2. 保持config.json到/etc/sing-box/config.json'
    echo '3. 启动 sudo sing-box run -c /etc/sing-box/config.json'
    echo '4. http/https代理设置：export https_proxy=http://127.0.0.1:2080 & export http_proxy=http://127.0.0.1:2080'
    echo '5. 访问测试 curl www.google.com'
fi


