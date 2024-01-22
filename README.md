mkdir -p ~/.ssh
ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519 -N ''
cat ~/.ssh/id_ed25519.pub


git config --global user.email "farabitolybai@gmail.com"
git config --global user.name "FarabiTol"
