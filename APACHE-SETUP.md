# Apache Setup Guide for GoDBAdmin

This guide explains how to configure Apache as a reverse proxy for GoDBAdmin.

## Prerequisites

- GoDBAdmin must be installed
- Apache must be installed
- Domain or server IP must be available

## Installing Apache

```bash
sudo apt-get update
sudo apt-get install apache2
sudo systemctl start apache2
sudo systemctl enable apache2
```

## Enabling Required Modules

```bash
sudo a2enmod proxy
sudo a2enmod proxy_http
sudo a2enmod headers
sudo a2enmod rewrite
sudo a2enmod ssl
sudo systemctl restart apache2
```

## Configuring Apache

### 1. Create Configuration File

```bash
sudo nano /etc/apache2/sites-available/godbadmin.conf
```

### 2. Configuration Content (with IP or without SSL)

```apache
<VirtualHost *:80>
    ServerName your-domain.com  # or your server IP
    
    # Frontend and API
    ProxyPreserveHost On
    ProxyRequests Off
    
    # Increase timeout for long queries
    ProxyTimeout 300
    
    # Frontend
    ProxyPass / http://localhost:8090/
    ProxyPassReverse / http://localhost:8090/
    
    # Headers
    RequestHeader set X-Forwarded-Proto "http"
    RequestHeader set X-Forwarded-Port "80"
</VirtualHost>
```

### 3. Enable Site

```bash
sudo a2ensite godbadmin.conf
sudo a2dissite 000-default.conf  # Disable default site (optional)
sudo apache2ctl configtest  # Test configuration
sudo systemctl reload apache2
```

## SSL Configuration with Certbot

### 1. Install Certbot

```bash
sudo apt-get update
sudo apt-get install certbot python3-certbot-apache
```

### 2. Obtain SSL Certificate

```bash
sudo certbot --apache -d your-domain.com
```

Or for multiple domains:

```bash
sudo certbot --apache -d your-domain.com -d www.your-domain.com
```

### 3. Automatic Configuration

Certbot automatically updates the Apache configuration and adds HTTP to HTTPS redirect.

### 4. Automatic Renewal

Certbot automatically renews certificates. To test:

```bash
sudo certbot renew --dry-run
```

## Final Configuration with SSL

After running certbot, the configuration will look like this:

```apache
<VirtualHost *:80>
    ServerName your-domain.com
    Redirect permanent / https://your-domain.com/
</VirtualHost>

<VirtualHost *:443>
    ServerName your-domain.com
    
    # SSL Configuration
    SSLEngine on
    SSLCertificateFile /etc/letsencrypt/live/your-domain.com/fullchain.pem
    SSLCertificateKeyFile /etc/letsencrypt/live/your-domain.com/privkey.pem
    Include /etc/letsencrypt/options-ssl-apache.conf
    
    # Frontend and API
    ProxyPreserveHost On
    ProxyRequests Off
    ProxyTimeout 300
    
    # Frontend
    ProxyPass / http://localhost:8090/
    ProxyPassReverse / http://localhost:8090/
    
    # Headers
    RequestHeader set X-Forwarded-Proto "https"
    RequestHeader set X-Forwarded-Port "443"
</VirtualHost>
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

```apache
<VirtualHost *:80>
    ServerName YOUR_SERVER_IP  # e.g., 192.168.1.100
    # ... rest of configuration
</VirtualHost>
```

And in build:

```bash
VITE_API_URL=http://YOUR_SERVER_IP/api npm run build
```

**Note:** For SSL with IP, you cannot use Let's Encrypt. You must use a self-signed certificate or purchase a domain.

## Verification

After configuration:

1. Check that Apache is running:
   ```bash
   sudo systemctl status apache2
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
# Install and configure Apache
sudo apt-get install apache2
sudo a2enmod proxy proxy_http headers rewrite ssl
sudo nano /etc/apache2/sites-available/godbadmin.conf
# (copy configuration above)
sudo a2ensite godbadmin.conf
sudo apache2ctl configtest
sudo systemctl reload apache2

# Install SSL
sudo apt-get install certbot python3-certbot-apache
sudo certbot --apache -d your-domain.com

# Check status
sudo systemctl status apache2
sudo systemctl status go-dbadmin
```
