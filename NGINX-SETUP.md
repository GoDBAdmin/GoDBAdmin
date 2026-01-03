# Nginx Setup Guide for GoDBAdmin

This guide explains how to configure Nginx as a reverse proxy for GoDBAdmin.

## Prerequisites

- GoDBAdmin must be installed
- Nginx must be installed
- Domain or server IP must be available

## Installing Nginx

```bash
sudo apt-get update
sudo apt-get install nginx
sudo systemctl start nginx
sudo systemctl enable nginx
```

## Configuring Nginx

### 1. Create Configuration File

```bash
sudo nano /etc/nginx/sites-available/godbadmin
```

### 2. Configuration Content (with IP or without SSL)

```nginx
server {
    listen 80;
    server_name your-domain.com;  # or your server IP

    # Increase timeout for long queries
    proxy_read_timeout 300s;
    proxy_connect_timeout 75s;

    # Frontend
    location / {
        proxy_pass http://localhost:8090;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
    }

    # API
    location /api {
        proxy_pass http://localhost:8090/api;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

### 3. Enable Site

```bash
sudo ln -s /etc/nginx/sites-available/godbadmin /etc/nginx/sites-enabled/
sudo nginx -t  # Test configuration
sudo systemctl reload nginx
```

## SSL Configuration with Certbot

### 1. Install Certbot

```bash
sudo apt-get update
sudo apt-get install certbot python3-certbot-nginx
```

### 2. Obtain SSL Certificate

```bash
sudo certbot --nginx -d your-domain.com
```

Or for multiple domains:

```bash
sudo certbot --nginx -d your-domain.com -d www.your-domain.com
```

### 3. Automatic Configuration

Certbot automatically updates the Nginx configuration and adds HTTP to HTTPS redirect.

### 4. Automatic Renewal

Certbot automatically renews certificates. To test:

```bash
sudo certbot renew --dry-run
```

## Final Configuration with SSL

After running certbot, the configuration will look like this:

```nginx
server {
    listen 80;
    server_name your-domain.com;
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name your-domain.com;

    ssl_certificate /etc/letsencrypt/live/your-domain.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/your-domain.com/privkey.pem;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers HIGH:!aNULL:!MD5;

    # Increase timeout
    proxy_read_timeout 300s;
    proxy_connect_timeout 75s;

    # Frontend
    location / {
        proxy_pass http://localhost:8090;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
    }

    # API
    location /api {
        proxy_pass http://localhost:8090/api;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

## Configuring Frontend to Use Domain

For the frontend to use the domain instead of localhost, you need to set the `VITE_API_URL` variable at build time.

### Method 1: Using Environment Variable

```bash
# Build frontend with API URL
cd frontend
VITE_API_URL=https://your-domain.com/api npm run build
```

### Method 2: Using .env File

In `frontend/.env.production`:

```
VITE_API_URL=https://your-domain.com/api
```

Then build:

```bash
cd frontend
npm run build
```

## Using IP Instead of Domain

If you don't have a domain and want to use an IP address:

```nginx
server {
    listen 80;
    server_name YOUR_SERVER_IP;  # e.g., 192.168.1.100
    # ... rest of configuration
}
```

And in build:

```bash
VITE_API_URL=http://YOUR_SERVER_IP/api npm run build
```

**Note:** For SSL with IP, you cannot use Let's Encrypt. You must use a self-signed certificate or purchase a domain.

## Verification

After configuration:

1. Check that Nginx is running:
   ```bash
   sudo systemctl status nginx
   ```

2. Check that GoDBAdmin is running:
   ```bash
   sudo systemctl status go-dbadmin
   ```

3. Test:
   ```bash
   curl http://your-domain.com
   # or
   curl https://your-domain.com
   ```

## Troubleshooting

### Error: "502 Bad Gateway"

- Check that GoDBAdmin is running: `sudo systemctl status go-dbadmin`
- Check that it's running on port 8090: `sudo netstat -tlnp | grep 8090`

### Error: "Connection refused"

- Check that GoDBAdmin is listening on `0.0.0.0:8090` (not just localhost)
- Check the configuration file `/etc/go-dbadmin/config.yaml`

### Error: SSL

- Check that the SSL certificate is valid: `sudo certbot certificates`
- Check that port 443 is open: `sudo ufw allow 443`

## Quick Commands

```bash
# Install and configure Nginx
sudo apt-get install nginx
sudo nano /etc/nginx/sites-available/godbadmin
# (copy configuration above)
sudo ln -s /etc/nginx/sites-available/godbadmin /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx

# Install SSL
sudo apt-get install certbot python3-certbot-nginx
sudo certbot --nginx -d your-domain.com

# Check status
sudo systemctl status nginx
sudo systemctl status go-dbadmin
```
