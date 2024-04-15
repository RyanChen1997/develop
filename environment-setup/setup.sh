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

if [ "$(git config --get credential.helper)" = "manager-core" ]; then
	echo "Git Credential Manager is configured"
else
	sudo apt install libsecret-1-0 libsecret-1-dev
	wget https://github.com/microsoft/Git-Credential-Manager-Core/releases/download/v2.0.498/gcmcore-linux_amd64.2.0.498.54650.tar.gz
	mkdir ~/git-credential-manager
	tar -xzf gcmcore-linux_amd64.2.0.498.54650.tar.gz -C ~/git-credential-manager
	echo '\n' >> ~/.bashrc
	echo '# git' >> ~/.bashrc
	echo 'export PATH="$PATH:$HOME/git-credential-manager"' >> ~/.bashrc
	source ~/.bashrc
	rm gcmcore-linux_amd64.2.0.498.54650.tar.gz
	git config --global credential.helper manager-core
	git config --global credential.credentialStore plaintext

	bash <(wget -qO- https://aka.ms/gcm/linux-install-source.sh)
fi
