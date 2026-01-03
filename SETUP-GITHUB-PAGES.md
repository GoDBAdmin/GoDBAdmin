# GitHub Pages Setup Guide for APT Repository

This guide explains how to set up `https://GoDBAdmin.github.io/GoDBAdmin`.

## Prerequisites

1. Repository `GoDBAdmin/GoDBAdmin` must be created on GitHub
2. Access to the repository for pushing
3. `.deb` package must be built

## Setup Steps

### Step 1: Build Debian Package

```bash
cd mysql-admin-tool
./local-scripts/build-deb.sh 1.0.0
```

The `.deb` package will be created in the parent directory.

### Step 2: Create Local APT Repository

```bash
./local-scripts/simple-repo.sh
```

This command:
- Copies the `.deb` package to `repo/pool/master/`
- Creates `Packages` and `Packages.gz` files

### Step 3: Create Branch for GitHub Pages

**Important Note:** GitHub Pages can use any branch. You can:
- Use the `gh-pages` branch (old default)
- Use the `main` or `master` branch
- Use any other branch you want

**Option 1: Using Separate Branch (Recommended)**
```bash
# Make sure you're on the main branch
git checkout main
git pull origin main

# Create new branch (you can use any name)
git checkout -b gh-pages
# or
git checkout -b pages
# or any other name you want
```

**Option 2: Using main Branch**
```bash
# If you want to use the main branch
git checkout main
git pull origin main
```

**Option 3: Using Another Branch**
```bash
# Create branch with desired name
git checkout -b apt-repo
```

### Step 4: Prepare Files for GitHub Pages

```bash
# Remove all previous files in gh-pages (optional)
# git rm -rf .

# Copy APT repository files
mkdir -p apt-repo
cp -r repo/* apt-repo/

# Or if you want to copy only essential files:
mkdir -p apt-repo/pool/main
cp repo/pool/master/*.deb apt-repo/pool/master/
cp repo/Packages apt-repo/
cp repo/Packages.gz apt-repo/
```

### Step 5: Commit and Push to GitHub

```bash
# Add files
git add apt-repo/

# Commit
git commit -m "Add APT repository for GoDBAdmin"

# Push to GitHub
git push origin gh-pages
```

### Step 6: Enable GitHub Pages

1. Go to the `GoDBAdmin/GoDBAdmin` repository on GitHub
2. Go to **Settings** → **Pages**
3. In the **Source** section:
   - **Branch**: Select your desired branch (e.g., `gh-pages`, `main`, or any other branch)
   - **Folder**: Select `/apt-repo` (or `/ (root)` if files are in root)
4. Click **Save**

**Note:** If you're using the `main` branch and files are in the `apt-repo/` folder, you must set the Folder to `/apt-repo`.

### Step 7: Verification

After a few minutes, the following address should be accessible:
```
https://GoDBAdmin.github.io/GoDBAdmin/
```

You can verify with:
```bash
curl -I https://GoDBAdmin.github.io/GoDBAdmin/Packages.gz
```

If you get `200 OK`, everything is working correctly!

## Updating Repository

To update the repository after building a new version:

```bash
# 1. Build new package
./local-scripts/build-deb.sh 1.0.1

# 2. Create new repository
./local-scripts/simple-repo.sh

# 3. Switch to gh-pages branch
git checkout gh-pages

# 4. Copy new files
cp repo/pool/master/*.deb apt-repo/pool/master/
cp repo/Packages apt-repo/
cp repo/Packages.gz apt-repo/

# 5. Commit and Push
git add apt-repo/
git commit -m "Update APT repository to version 1.0.1"
git push origin gh-pages
```

## Final Structure

After setup, the structure in the selected branch should look like this:

**If using apt-repo folder:**
```
branch-name/
└── apt-repo/
    ├── Packages
    ├── Packages.gz
    └── pool/
        └── main/
            └── go-dbadmin_1.0.0-1_amd64.deb
```

**If files are placed directly in root:**
```
branch-name/
├── Packages
├── Packages.gz
└── pool/
    └── main/
        └── go-dbadmin_1.0.0-1_amd64.deb
```

**Note:** In GitHub Pages settings, if files are in `apt-repo/`, set Folder to `/apt-repo`. If they're in root, select `/ (root)`.

## Troubleshooting

### Error: "404 Not Found"

- Make sure the `gh-pages` branch has been created
- Check that GitHub Pages is enabled (Settings → Pages)
- Wait a few minutes for GitHub Pages to deploy

### Error: "Packages file not found"

- Make sure the `Packages.gz` file exists in `apt-repo/`
- Check that the path is correct: `apt-repo/Packages.gz`

### Error: "Repository structure incorrect"

- The structure must be exactly like this:
  ```
  apt-repo/
  ├── Packages
  ├── Packages.gz
  └── pool/master/*.deb
  ```

## Using Automated Script (Recommended)

You can create a script to automate this process:

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
cp repo/pool/master/*.deb apt-repo/pool/master/
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

## Important Notes

1. **Always keep the `gh-pages` branch separate** - This branch is only for GitHub Pages
2. **Don't add unnecessary files** - Only push APT repository files
3. **Use `.gitignore`** - Ignore build and temporary files
4. **Check that files are public** - GitHub Pages only serves public files

## Support

If you encounter any issues:
- Issues: https://github.com/GoDBAdmin/GoDBAdmin/issues
- Documentation: [GitHub APT Repository Guide](docs/GITHUB-APT-REPO.md)
