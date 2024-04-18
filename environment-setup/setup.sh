#!/bin/bash

sudo apt update

# ------------------ docker install ------------------------
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

# ----------------- git --------------------------------------
if command -v git >/dev/null 2>&1; then
	echo "git is installed"
else
	echo "install git"
	sudo apt install git
fi

# git push需要申请git token
git config --global credential.helper store

