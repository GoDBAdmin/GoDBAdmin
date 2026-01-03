# راهنمای راه‌اندازی GitHub Pages برای مخزن APT

این راهنما نحوه راه‌اندازی `https://GoDBAdmin.github.io/GoDBAdmin` را توضیح می‌دهد.

## پیش‌نیازها

1. ریپو `GoDBAdmin/GoDBAdmin` در GitHub ایجاد شده باشد
2. دسترسی به ریپو برای push کردن
3. پکیج `.deb` ساخته شده باشد

## مراحل راه‌اندازی

### مرحله 1: ساخت پکیج Debian

```bash
cd mysql-admin-tool
./local-scripts/build-deb.sh 1.0.0
```

پکیج `.deb` در دایرکتوری والد ایجاد می‌شود.

### مرحله 2: ایجاد مخزن APT محلی

```bash
./local-scripts/simple-repo.sh
```

این دستور:
- پکیج `.deb` را به `repo/pool/main/` کپی می‌کند
- فایل `Packages` و `Packages.gz` را ایجاد می‌کند

### مرحله 3: ایجاد branch برای GitHub Pages

**نکته مهم:** GitHub Pages می‌تواند از هر branch استفاده کند. شما می‌توانید:
- از branch `gh-pages` استفاده کنید (پیش‌فرض قدیمی)
- از branch `main` یا `master` استفاده کنید
- از هر branch دیگری که می‌خواهید استفاده کنید

**گزینه 1: استفاده از branch جداگانه (پیشنهادی)**
```bash
# اطمینان از اینکه در branch main هستید
git checkout main
git pull origin main

# ایجاد branch جدید (می‌توانید هر نامی بدهید)
git checkout -b gh-pages
# یا
git checkout -b pages
# یا هر نام دیگری که می‌خواهید
```

**گزینه 2: استفاده از branch main**
```bash
# اگر می‌خواهید از branch main استفاده کنید
git checkout main
git pull origin main
```

**گزینه 3: استفاده از branch دیگر**
```bash
# ایجاد branch با نام دلخواه
git checkout -b apt-repo
```

### مرحله 4: آماده‌سازی فایل‌ها برای GitHub Pages

```bash
# پاک کردن همه فایل‌های قبلی در gh-pages (اختیاری)
# git rm -rf .

# کپی کردن فایل‌های مخزن APT
mkdir -p apt-repo
cp -r repo/* apt-repo/

# یا اگر می‌خواهید فقط فایل‌های ضروری را کپی کنید:
mkdir -p apt-repo/pool/main
cp repo/pool/main/*.deb apt-repo/pool/main/
cp repo/Packages apt-repo/
cp repo/Packages.gz apt-repo/
```

### مرحله 5: Commit و Push به GitHub

```bash
# اضافه کردن فایل‌ها
git add apt-repo/

# Commit
git commit -m "Add APT repository for GoDBAdmin"

# Push به GitHub
git push origin gh-pages
```

### مرحله 6: فعال‌سازی GitHub Pages

1. به ریپو `GoDBAdmin/GoDBAdmin` در GitHub بروید
2. به **Settings** → **Pages** بروید
3. در بخش **Source**:
   - **Branch**: branch مورد نظر خود را انتخاب کنید (مثلاً `gh-pages`، `main`، یا هر branch دیگری)
   - **Folder**: `/apt-repo` را انتخاب کنید (یا `/ (root)` اگر فایل‌ها در root هستند)
4. روی **Save** کلیک کنید

**نکته:** اگر از branch `main` استفاده می‌کنید و فایل‌ها در پوشه `apt-repo/` هستند، باید Folder را `/apt-repo` تنظیم کنید.

### مرحله 7: بررسی

پس از چند دقیقه، آدرس زیر باید در دسترس باشد:
```
https://GoDBAdmin.github.io/GoDBAdmin/
```

می‌توانید با دستور زیر بررسی کنید:
```bash
curl -I https://GoDBAdmin.github.io/GoDBAdmin/Packages.gz
```

اگر `200 OK` دریافت کردید، همه چیز درست است!

## به‌روزرسانی مخزن

برای به‌روزرسانی مخزن بعد از ساخت نسخه جدید:

```bash
# 1. ساخت پکیج جدید
./local-scripts/build-deb.sh 1.0.1

# 2. ایجاد مخزن جدید
./local-scripts/simple-repo.sh

# 3. رفتن به branch gh-pages
git checkout gh-pages

# 4. کپی کردن فایل‌های جدید
cp repo/pool/main/*.deb apt-repo/pool/main/
cp repo/Packages apt-repo/
cp repo/Packages.gz apt-repo/

# 5. Commit و Push
git add apt-repo/
git commit -m "Update APT repository to version 1.0.1"
git push origin gh-pages
```

## ساختار نهایی

پس از راه‌اندازی، ساختار در branch انتخابی باید به این صورت باشد:

**اگر از پوشه apt-repo استفاده می‌کنید:**
```
branch-name/
└── apt-repo/
    ├── Packages
    ├── Packages.gz
    └── pool/
        └── main/
            └── go-dbadmin_1.0.0-1_amd64.deb
```

**اگر فایل‌ها را مستقیماً در root قرار می‌دهید:**
```
branch-name/
├── Packages
├── Packages.gz
└── pool/
    └── main/
        └── go-dbadmin_1.0.0-1_amd64.deb
```

**نکته:** در تنظیمات GitHub Pages، اگر فایل‌ها در `apt-repo/` هستند، Folder را `/apt-repo` تنظیم کنید. اگر در root هستند، `/ (root)` را انتخاب کنید.

## عیب‌یابی

### خطا: "404 Not Found"

- مطمئن شوید که branch `gh-pages` ایجاد شده است
- بررسی کنید که GitHub Pages فعال است (Settings → Pages)
- چند دقیقه صبر کنید تا GitHub Pages deploy شود

### خطا: "Packages file not found"

- مطمئن شوید که فایل `Packages.gz` در `apt-repo/` وجود دارد
- بررسی کنید که مسیر درست است: `apt-repo/Packages.gz`

### خطا: "Repository structure incorrect"

- ساختار باید دقیقاً به این صورت باشد:
  ```
  apt-repo/
  ├── Packages
  ├── Packages.gz
  └── pool/main/*.deb
  ```

## استفاده از اسکریپت خودکار (پیشنهادی)

می‌توانید یک اسکریپت برای خودکارسازی این فرآیند ایجاد کنید:

```bash
#!/bin/bash
# deploy-apt-repo.sh

VERSION=${1:-1.0.0}

# Build package
./local-scripts/build-deb.sh $VERSION

# Create repository
./local-scripts/simple-repo.sh

# Switch to gh-pages
git checkout gh-pages || git checkout -b gh-pages

# Copy files
mkdir -p apt-repo/pool/main
cp repo/pool/main/*.deb apt-repo/pool/main/
cp repo/Packages apt-repo/
cp repo/Packages.gz apt-repo/

# Commit and push
git add apt-repo/
git commit -m "Update APT repository to version $VERSION"
git push origin gh-pages

# Switch back to main
git checkout main

echo "APT repository deployed successfully!"
echo "URL: https://GoDBAdmin.github.io/GoDBAdmin/"
```

## نکات مهم

1. **همیشه branch `gh-pages` را جدا نگه دارید** - این branch فقط برای GitHub Pages است
2. **فایل‌های غیرضروری را اضافه نکنید** - فقط فایل‌های مخزن APT را push کنید
3. **از `.gitignore` استفاده کنید** - فایل‌های build و temporary را ignore کنید
4. **بررسی کنید که فایل‌ها public هستند** - GitHub Pages فقط فایل‌های public را serve می‌کند

## پشتیبانی

اگر مشکلی داشتید:
- Issues: https://github.com/GoDBAdmin/GoDBAdmin/issues
- Documentation: [GitHub APT Repository Guide](docs/GITHUB-APT-REPO.md)

