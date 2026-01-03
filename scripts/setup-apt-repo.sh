#!/bin/bash

# GoDBAdmin - Setup APT Repository from GitHub
# This script adds the GitHub Pages APT repository to apt sources

set -e

GITHUB_REPO="GoDBAdmin/GoDBAdmin"
GITHUB_PAGES_URL="https://GoDBAdmin.github.io/GoDBAdmin"

echo "========================================="
echo "GoDBAdmin - Setup APT Repository"
echo "========================================="
echo ""
echo "Repository: https://github.com/$GITHUB_REPO"
echo "APT Repository: $GITHUB_PAGES_URL"
echo ""

# Check if running as root or with sudo
if [ "$EUID" -ne 0 ]; then 
    echo "This script requires sudo privileges"
    echo "Please run: sudo bash -c \"\$(curl -sSL https://raw.githubusercontent.com/$GITHUB_REPO/main/scripts/setup-apt-repo.sh)\""
    exit 1
fi

# Add repository to sources.list.d
echo "Adding repository to apt sources..."
cat > /etc/apt/sources.list.d/go-dbadmin.list <<EOF
deb [trusted=yes] $GITHUB_PAGES_URL /
EOF

# Import GPG key if needed (for signed repositories)
# Uncomment if you add GPG signing later
# curl -sSL https://raw.githubusercontent.com/$GITHUB_REPO/main/apt-repo/KEY.gpg | apt-key add -

# Update apt
echo "Updating apt cache..."
apt-get update

echo ""
echo "========================================="
echo "Repository setup complete!"
echo "========================================="
echo ""
echo "You can now install with:"
echo "  sudo apt-get install go-dbadmin"
echo ""
echo "To remove repository:"
echo "  sudo rm /etc/apt/sources.list.d/go-dbadmin.list"
echo "  sudo apt-get update"
echo ""

