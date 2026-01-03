# راهنمای فعال‌سازی GitHub Pages

برای فعال‌سازی آدرس `https://GoDBAdmin.github.io/GoDBAdmin` باید مراحل زیر را انجام دهید:

## مرحله 1: Deploy فایل‌ها به GitHub

ابتدا باید فایل‌های مخزن APT را به branch `master` push کنید:

```bash
cd mysql-admin-tool
./local-scripts/deploy-apt-repo-to-github.sh 1.0.0 master
```

یا به صورت دستی:

```bash
# 1. ساخت پکیج
./local-scripts/build-deb.sh 1.0.0

# 2. ایجاد مخزن APT
./local-scripts/simple-repo.sh

# 3. آماده‌سازی فایل‌ها
mkdir -p apt-repo/pool/main
cp ../go-dbadmin_*.deb apt-repo/pool/master/
cp repo/Packages apt-repo/
cp repo/Packages.gz apt-repo/

# 4. Commit و Push
git checkout master
git add apt-repo/
git commit -m "Add APT repository"
git push origin master
```

## مرحله 2: فعال‌سازی GitHub Pages در GitHub

### روش 1: از طریق وب‌سایت GitHub

1. به ریپو `GoDBAdmin/GoDBAdmin` در GitHub بروید
2. روی **Settings** کلیک کنید (در منوی بالای ریپو)
3. در منوی سمت چپ، **Pages** را انتخاب کنید
4. در بخش **Source**:
   - **Branch**: `master` را انتخاب کنید
   - **Folder**: `/apt-repo` را انتخاب کنید (یا `/ (root)` اگر فایل‌ها در root هستند)
5. روی **Save** کلیک کنید

### روش 2: از طریق GitHub CLI (اختیاری)

```bash
gh repo edit GoDBAdmin/GoDBAdmin --enable-pages --pages-branch master --pages-source /apt-repo
```

## مرحله 3: بررسی

بعد از فعال‌سازی:

1. **چند دقیقه صبر کنید** (معمولاً 1-5 دقیقه) تا GitHub Pages deploy شود
2. **بررسی کنید** که deploy موفق بوده:
   - به Settings → Pages بروید
   - در بخش "Build and deployment" باید پیام "Your site is live at..." را ببینید
3. **تست کنید**:
   ```bash
   curl -I https://GoDBAdmin.github.io/GoDBAdmin/apt-repo/Packages.gz
   ```
   باید پاسخ `200 OK` دریافت کنید

## ساختار مورد نیاز

برای اینکه GitHub Pages کار کند، ساختار فایل‌ها باید به این صورت باشد:

```
master/
└── apt-repo/
    ├── Packages
    ├── Packages.gz
    └── pool/
        └── main/
            └── go-dbadmin_1.0.0-1_amd64.deb
```

**نکته مهم:** در تنظیمات GitHub Pages، اگر فایل‌ها در `apt-repo/` هستند، Folder را `/apt-repo` تنظیم کنید.

## عیب‌یابی

### مشکل: "404 Not Found"

**راه‌حل:**
1. مطمئن شوید که branch `master` push شده است
2. بررسی کنید که GitHub Pages فعال است (Settings → Pages)
3. مطمئن شوید که Folder را درست تنظیم کرده‌اید:
   - اگر فایل‌ها در `apt-repo/` هستند → `/apt-repo`
   - اگر فایل‌ها در root هستند → `/ (root)`
4. چند دقیقه صبر کنید تا deploy شود

### مشکل: "Your site is ready to be published"

**راه‌حل:**
- این پیام یعنی GitHub Pages هنوز deploy نشده است
- چند دقیقه صبر کنید
- یا یک commit جدید push کنید تا deploy دوباره شروع شود

### مشکل: "Build failed"

**راه‌حل:**
- بررسی کنید که فایل `index.html` یا فایل‌های static در مسیر درست هستند
- برای مخزن APT، نیازی به `index.html` نیست
- فقط فایل‌های `Packages` و `Packages.gz` و `.deb` کافی است

## بررسی وضعیت Deploy

برای بررسی وضعیت deploy:

1. به Settings → Pages بروید
2. در بخش "Build and deployment" می‌توانید:
   - آخرین deploy را ببینید
   - تاریخ و زمان deploy را ببینید
   - در صورت خطا، پیام خطا را ببینید

## نکات مهم

1. **اولین deploy ممکن است 5-10 دقیقه طول بکشد**
2. **بعد از هر push جدید، deploy دوباره انجام می‌شود**
3. **فقط branch های public قابل استفاده برای GitHub Pages هستند**
4. **اگر ریپو private است، باید GitHub Pro داشته باشید**

## دستورات سریع

```bash
# Deploy کامل
./local-scripts/deploy-apt-repo-to-github.sh 1.0.0 master

# بررسی دسترسی
curl -I https://GoDBAdmin.github.io/GoDBAdmin/apt-repo/Packages.gz

# بررسی محتویات
curl https://GoDBAdmin.github.io/GoDBAdmin/apt-repo/Packages.gz | gunzip | head -20
```

## پشتیبانی

اگر مشکل داشتید:
- Issues: https://github.com/GoDBAdmin/GoDBAdmin/issues
- GitHub Pages Docs: https://docs.github.com/en/pages

