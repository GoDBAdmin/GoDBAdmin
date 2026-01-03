# راهنمای کامل Deploy مخزن APT به GitHub Pages (Branch Master)

این راهنما نحوه deploy کردن مخزن APT به GitHub Pages با استفاده از branch `master` را توضیح می‌دهد.

## دستور کامل Deploy

```bash
cd mysql-admin-tool
./local-scripts/deploy-apt-repo-to-github.sh 1.0.0 master
```

یا به صورت ساده (چون `master` پیش‌فرض است):

```bash
cd mysql-admin-tool
./local-scripts/deploy-apt-repo-to-github.sh 1.0.0
```

## مراحل دستی (اگر اسکریپت کار نکرد)

### 1. ساخت پکیج Debian

```bash
cd mysql-admin-tool
./local-scripts/build-deb.sh 1.0.0
```

### 2. ایجاد مخزن APT

```bash
./local-scripts/simple-repo.sh
```

### 3. آماده‌سازی فایل‌ها

```bash
# ایجاد پوشه apt-repo
mkdir -p apt-repo/pool/main

# کپی فایل .deb
cp ../go-dbadmin_*.deb apt-repo/pool/main/

# کپی فایل‌های Packages
cp repo/Packages apt-repo/
cp repo/Packages.gz apt-repo/
```

### 4. Commit و Push به branch master

```bash
# اطمینان از اینکه در branch master هستید
git checkout master
git pull origin master

# اضافه کردن فایل‌ها
git add apt-repo/

# Commit
git commit -m "Add APT repository for GoDBAdmin v1.0.0"

# Push به GitHub
git push origin master
```

## تنظیمات GitHub Pages

بعد از push، در GitHub:

1. به ریپو `GoDBAdmin/GoDBAdmin` بروید
2. **Settings** → **Pages** را باز کنید
3. در بخش **Source**:
   - **Branch**: `master` را انتخاب کنید
   - **Folder**: `/apt-repo` را انتخاب کنید
4. روی **Save** کلیک کنید

## استفاده در سرور اوبونتو

بعد از فعال‌سازی GitHub Pages، در سرور اوبونتو:

### روش 1: استفاده از اسکریپت (پیشنهادی)

```bash
curl -sSL https://raw.githubusercontent.com/GoDBAdmin/GoDBAdmin/master/scripts/setup-apt-repo.sh | sudo bash
sudo apt-get update
sudo apt-get install go-dbadmin
```

### روش 2: دستی

```bash
# اضافه کردن مخزن
echo "deb [trusted=yes] https://GoDBAdmin.github.io/GoDBAdmin/apt-repo /" | sudo tee /etc/apt/sources.list.d/go-dbadmin.list

# به‌روزرسانی apt
sudo apt-get update

# نصب
sudo apt-get install go-dbadmin
```

## بررسی

برای بررسی اینکه همه چیز درست کار می‌کند:

```bash
# بررسی دسترسی به فایل Packages.gz
curl -I https://GoDBAdmin.github.io/GoDBAdmin/apt-repo/Packages.gz

# باید پاسخ 200 OK دریافت کنید
```

## ساختار نهایی در branch master

```
master/
├── (سایر فایل‌های پروژه)
└── apt-repo/
    ├── Packages
    ├── Packages.gz
    └── pool/
        └── main/
            └── go-dbadmin_1.0.0-1_amd64.deb
```

## به‌روزرسانی مخزن

برای به‌روزرسانی بعد از ساخت نسخه جدید:

```bash
# ساخت پکیج جدید
./local-scripts/build-deb.sh 1.0.1

# ایجاد مخزن جدید
./local-scripts/simple-repo.sh

# کپی فایل‌های جدید
cp ../go-dbadmin_*.deb apt-repo/pool/main/
cp repo/Packages apt-repo/
cp repo/Packages.gz apt-repo/

# Commit و Push
git add apt-repo/
git commit -m "Update APT repository to version 1.0.1"
git push origin master
```

## عیب‌یابی

### خطا: "404 Not Found"

- مطمئن شوید که branch `master` push شده است
- بررسی کنید که GitHub Pages فعال است (Settings → Pages)
- مطمئن شوید که Folder را `/apt-repo` تنظیم کرده‌اید
- چند دقیقه صبر کنید تا GitHub Pages deploy شود

### خطا: "Packages file not found"

- بررسی کنید که فایل `Packages.gz` در `apt-repo/` وجود دارد
- مسیر درست است: `apt-repo/Packages.gz`
- در تنظیمات GitHub Pages، Folder را `/apt-repo` تنظیم کنید

### خطا در apt-get update

```bash
# بررسی محتویات فایل sources.list
cat /etc/apt/sources.list.d/go-dbadmin.list

# باید این باشد:
# deb [trusted=yes] https://GoDBAdmin.github.io/GoDBAdmin/apt-repo /

# بررسی دسترسی
curl -I https://GoDBAdmin.github.io/GoDBAdmin/apt-repo/Packages.gz
```

## نکات مهم

1. **همیشه branch `master` را pull کنید** قبل از deploy
2. **فقط فایل‌های `apt-repo/` را commit کنید** - سایر فایل‌های پروژه را تغییر ندهید
3. **Folder را `/apt-repo` تنظیم کنید** در تنظیمات GitHub Pages
4. **چند دقیقه صبر کنید** بعد از push تا GitHub Pages به‌روزرسانی شود

## دستورات سریع

```bash
# Deploy کامل (یک خط)
./local-scripts/deploy-apt-repo-to-github.sh 1.0.0 master

# نصب در سرور اوبونتو (یک خط)
curl -sSL https://raw.githubusercontent.com/GoDBAdmin/GoDBAdmin/master/scripts/setup-apt-repo.sh | sudo bash && sudo apt-get update && sudo apt-get install go-dbadmin
```

