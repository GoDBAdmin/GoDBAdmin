# Complete Production Setup Guide for GoDBAdmin

This guide explains how to configure GoDBAdmin in a production environment with Nginx/Apache and SSL.

## Table of Contents

1. [Installing GoDBAdmin](#installing-godbadmin)
2. [Configuring Nginx](#configuring-nginx)
3. [Configuring Apache](#configuring-apache)
4. [Configuring SSL with Certbot](#configuring-ssl-with-certbot)
5. [Configuring Frontend for Domain](#configuring-frontend-for-domain)
6. [Troubleshooting](#troubleshooting)

---

## Installing GoDBAdmin

```bash
# Install from APT repository
curl -sSL https://raw.githubusercontent.com/GoDBAdmin/GoDBAdmin/master/scripts/setup-apt-repo.sh | sudo bash
sudo apt-get update
sudo apt-get install go-dbadmin

# Or install from .deb file
sudo dpkg -i go-dbadmin_*.deb
sudo apt-get install -f
```

---

## Configuring Nginx

### Method 1: Using Automated Script (Recommended)

```bash
# With domain and SSL
sudo ./local-scripts/setup-nginx.sh your-domain.com yes

# With IP only (without SSL)
sudo ./local-scripts/setup-nginx.sh YOUR_IP no
```

### Method 2: Manual Configuration

Complete guide in [NGINX-SETUP.md](NGINX-SETUP.md)

---

## Configuring Apache

### Method 1: Using Automated Script (Recommended)

```bash
# With domain and SSL
sudo ./local-scripts/setup-apache.sh your-domain.com yes

# With IP only (without SSL)
sudo ./local-scripts/setup-apache.sh YOUR_IP no
```

### Method 2: Manual Configuration

Complete guide in [APACHE-SETUP.md](APACHE-SETUP.md)

---

## Configuring SSL with Certbot

### Installing Certbot

**For Nginx:**
```bash
sudo apt-get update
sudo apt-get install certbot python3-certbot-nginx
```

**For Apache:**
```bash
sudo apt-get update
sudo apt-get install certbot python3-certbot-apache
```

### Obtaining SSL Certificate

**For Nginx:**
```bash
sudo certbot --nginx -d your-domain.com
```

**For Apache:**
```bash
sudo certbot --apache -d your-domain.com
```

### Automatic Renewal

Certbot automatically renews certificates. To test:

```bash
sudo certbot renew --dry-run
```

---

## Configuring Frontend for Domain

For the frontend to use the domain or server IP instead of `localhost`, you need to set the `VITE_API_URL` variable at build time.

### Method 1: Using Environment Variable

```bash
cd frontend
VITE_API_URL=https://your-domain.com/api npm run build
# or for IP
VITE_API_URL=http://YOUR_IP/api npm run build
```

### Method 2: Using .env File

In `frontend/.env.production`:

```env
VITE_API_URL=https://your-domain.com/api
```

Then build:

```bash
cd frontend
npm run build
```

### Method 3: Using .env.local File (for production)

```bash
cd frontend
echo "VITE_API_URL=https://your-domain.com/api" > .env.production
npm run build
```

**Important Note:** After rebuilding the frontend, you must rebuild and reinstall the `.deb` package.

---

## Complete Setup Steps

### 1. Install GoDBAdmin

```bash
curl -sSL https://raw.githubusercontent.com/GoDBAdmin/GoDBAdmin/master/scripts/setup-apt-repo.sh | sudo bash
sudo apt-get update
sudo apt-get install go-dbadmin
```

### 2. Configure Reverse Proxy

**For Nginx:**
```bash
sudo ./local-scripts/setup-nginx.sh your-domain.com yes
```

**For Apache:**
```bash
sudo ./local-scripts/setup-apache.sh your-domain.com yes
```

### 3. Rebuild Frontend with Correct API URL (if needed)

```bash
# Clone repository
git clone https://github.com/GoDBAdmin/GoDBAdmin.git
cd GoDBAdmin/frontend

# Build with API URL
VITE_API_URL=https://your-domain.com/api npm run build

# Build new package
cd ..
./local-scripts/build-deb.sh 1.0.1
sudo dpkg -i ../go-dbadmin_1.0.1-1_amd64.deb
sudo systemctl restart go-dbadmin
```

---

## Verification

### Check Service Status

```bash
# Check GoDBAdmin
sudo systemctl status go-dbadmin

# Check Nginx
sudo systemctl status nginx

# Check Apache
sudo systemctl status apache2
```

### Test Access

```bash
# Test HTTP
curl http://your-domain.com

# Test HTTPS
curl https://your-domain.com

# Test API
curl https://your-domain.com/api/databases
```

---

## Troubleshooting

### Issue: Frontend Makes Requests to localhost

**Solution:**
1. Rebuild frontend with correct `VITE_API_URL`
2. Rebuild and reinstall the package
3. Restart the service

### Issue: 502 Bad Gateway

**Solution:**
1. Check that GoDBAdmin is running: `sudo systemctl status go-dbadmin`
2. Check that it's listening on port 8090: `sudo netstat -tlnp | grep 8090`
3. Check that the config has `host: "0.0.0.0"` (not `localhost`)

### Issue: SSL Not Working

**Solution:**
1. Check that the certificate is valid: `sudo certbot certificates`
2. Check that port 443 is open: `sudo ufw allow 443`
3. Check that DNS is configured correctly

### Issue: CORS Error

**Solution:**
- This should not occur as the backend has CORS enabled for all origins
- If it occurs, check that proxy headers are configured correctly

---

## Security

### Firewall

```bash
# Open required ports
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw allow 22/tcp  # SSH
sudo ufw enable
```

### Updates

```bash
# Update system
sudo apt-get update
sudo apt-get upgrade

# Update GoDBAdmin
sudo apt-get update
sudo apt-get install --only-upgrade go-dbadmin
```

---

## Quick Commands

```bash
# Complete setup with Nginx and SSL
curl -sSL https://raw.githubusercontent.com/GoDBAdmin/GoDBAdmin/master/scripts/setup-apt-repo.sh | sudo bash
sudo apt-get update && sudo apt-get install go-dbadmin
sudo ./local-scripts/setup-nginx.sh your-domain.com yes

# Restart services
sudo systemctl restart go-dbadmin
sudo systemctl restart nginx

# Check logs
sudo journalctl -u go-dbadmin -f
sudo tail -f /var/log/nginx/error.log
```

---

## Support

- Issues: https://github.com/GoDBAdmin/GoDBAdmin/issues
- Documentation: [NGINX-SETUP.md](NGINX-SETUP.md), [APACHE-SETUP.md](APACHE-SETUP.md)
