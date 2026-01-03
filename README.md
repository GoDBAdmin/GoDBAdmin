# GoDBAdmin

<div align="center">

![GoDBAdmin](https://img.shields.io/badge/GoDBAdmin-MySQL%20Admin-blue)
![License](https://img.shields.io/badge/license-Apache%202.0-blue)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8)
![React](https://img.shields.io/badge/React-18+-61DAFB)

**A modern, fast, and beautiful web-based MySQL database administration tool built with Go and React.**

[Features](#features) • [Installation](#installation) • [Documentation](#documentation) • [Support](#support)

</div>

---

## Overview

GoDBAdmin is a powerful, open-source MySQL database administration tool that provides a modern web interface for managing your databases. Built with Go (Fiber) for high performance and React for a beautiful user experience, it offers an intuitive way to browse databases, execute queries, and manage your MySQL infrastructure.

## Features

- 🚀 **High Performance**: Built with Go (Fiber) for fast query execution and low resource usage
- 🎨 **Beautiful UI**: Modern React interface with Tailwind CSS and dark mode support
- 📊 **Database Browser**: Intuitive tree view of databases and tables with real-time navigation
- ✏️ **SQL Editor**: Monaco editor with syntax highlighting, autocomplete, and query history
- 📋 **Table Management**: View table structure, data with pagination, and export capabilities
- 🔐 **Secure**: JWT-based authentication with encrypted connections
- 📦 **Easy Installation**: Debian/Ubuntu package support with APT repository
- 🔄 **Real-time Updates**: Live database structure updates and query results
- 📱 **Responsive Design**: Works seamlessly on desktop, tablet, and mobile devices

## Screenshots

*Screenshots coming soon*

## Requirements

### Runtime Requirements
- MySQL 5.7+ or MariaDB 10.3+
- Debian/Ubuntu (for package installation) or any Linux distribution (for manual installation)

### Development Requirements
- Go 1.21 or higher
- Node.js 18 or higher
- npm or pnpm

## Quick Start

### Installation from GitHub APT Repository (Recommended)

```bash
# One-line installation
curl -sSL https://raw.githubusercontent.com/GoDBAdmin/GoDBAdmin/master/scripts/setup-apt-repo.sh | sudo bash
sudo apt-get update
sudo apt-get install go-dbadmin
```

The service will start automatically. Access the web interface at `http://localhost:8090`

### Manual Installation

See [Installation Guide](docs/APT-INSTALL.md) for detailed instructions.

## Development

### Quick Start (Recommended)

```bash
# Run both backend and frontend together
./run.sh
```

This will start:
- Backend on `http://localhost:8090`
- Frontend on `http://localhost:3000`

To stop:
```bash
./stop.sh
```

### Manual Development Setup

#### Backend

```bash
cd backend
go mod download
FRONTEND_PATH=../frontend/dist PORT=8090 go run ./cmd/server/main.go
```

#### Frontend

```bash
cd frontend
npm install
npm run dev
```

### Environment Variables

Backend environment variables:

- `DB_HOST`: MySQL host (default: localhost)
- `DB_PORT`: MySQL port (default: 3306)
- `DB_USER`: MySQL username (default: root)
- `DB_PASSWORD`: MySQL password (default: empty)
- `DB_NAME`: MySQL database (default: mysql)
- `PORT`: Server port (default: 8090)
- `JWT_SECRET`: JWT secret key (required for production)
- `FRONTEND_PATH`: Path to frontend assets (default: ./frontend/dist)

## Building

### Build Backend

```bash
make build-backend
```

### Build Frontend

```bash
make build-frontend
```

### Build Everything

```bash
make all
```

### Creating Debian Package

```bash
# Install build dependencies
sudo apt-get install build-essential debhelper golang-go nodejs npm

# Build package
./build-deb.sh 1.0.0

# The .deb file will be created in the parent directory
# Then create APT repository:
./simple-repo.sh
```

See [Building Debian Package](docs/BUILD-DEB.md) for detailed instructions.

## Installation Methods

### 1. From GitHub APT Repository (Recommended)

See [GitHub APT Repository Guide](docs/GITHUB-APT-REPO.md) for detailed instructions.

### 2. From Local APT Repository

See [APT Installation Guide](docs/APT-INSTALL.md) for detailed instructions.

### 3. From .deb Package

```bash
sudo dpkg -i go-dbadmin_*.deb
sudo apt-get install -f  # Install dependencies if needed
```

### Configuration

Edit `/etc/go-dbadmin/config.yaml` (optional):

```yaml
server:
  port: 8090
frontend:
  path: /usr/share/go-dbadmin/frontend
logging:
  level: info
  file: /var/log/go-dbadmin/app.log
```

### Start Service

```bash
sudo systemctl start go-dbadmin
sudo systemctl enable go-dbadmin
```

Access the web interface at `http://localhost:8090`

## Usage

1. **Login**: Access the web interface and login with your MySQL credentials
2. **Browse**: Navigate databases and tables from the sidebar
3. **Explore**: Click on a table to view its structure and data
4. **Query**: Use the Query Editor to execute custom SQL queries
5. **Manage**: View results in a formatted table with export options

For more details, see [Running Guide](docs/RUN.md).

## API Endpoints

### Authentication
- `POST /api/auth/login` - Login with MySQL credentials

### Databases
- `GET /api/databases` - List all databases
- `GET /api/databases/:db/tables` - List tables in a database
- `GET /api/databases/:db/tables/:table` - Get table structure
- `GET /api/databases/:db/tables/:table/data` - Get table data (paginated)
- `POST /api/databases/:db/tables` - Create new table
- `DELETE /api/databases/:db/tables/:table` - Drop table

### Queries
- `POST /api/query` - Execute custom SQL query

## Documentation

Comprehensive documentation is available in the `docs/` directory:

- [APT Installation Guide](docs/APT-INSTALL.md) - Installing from APT repository
- [APT Repository Setup](docs/APT-REPO.md) - Setting up APT repositories
- [Building Debian Package](docs/BUILD-DEB.md) - Building .deb packages
- [Running Guide](docs/RUN.md) - Development and running instructions
- [GitHub APT Repository](docs/GITHUB-APT-REPO.md) - Using GitHub-hosted repository

## Contributing

We welcome contributions! Please feel free to submit a Pull Request. Here's how you can help:

1. **Fork the repository**
2. **Create a feature branch** (`git checkout -b feature/amazing-feature`)
3. **Commit your changes** (`git commit -m 'Add some amazing feature'`)
4. **Push to the branch** (`git push origin feature/amazing-feature`)
5. **Open a Pull Request**

### Development Guidelines

- Follow Go and React best practices
- Write clear commit messages
- Add tests for new features
- Update documentation as needed

## Support

### Financial Support

GoDBAdmin is an open-source project maintained by the community. Your support helps us continue development and improve the tool.

#### Support Plans

We accept cryptocurrency donations through multiple networks:

##### 🥉 Bronze Supporter - $10/month
- Your name in the README contributors section
- Priority issue responses
- Early access to new features

**Payment Options:**
- **Bitcoin (BTC)**: `bc1qxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
- **Ethereum (ETH)**: `0x1234567890123456789012345678901234567890`
- **USDT (TRC20)**: `Txxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
- **USDT (ERC20)**: `0x1234567890123456789012345678901234567890`
- **USDC (ERC20)**: `0x1234567890123456789012345678901234567890`
- **Litecoin (LTC)**: `Lxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
- **Dogecoin (DOGE)**: `Dxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
- **Binance Coin (BNB)**: `0x1234567890123456789012345678901234567890`
- **Cardano (ADA)**: `addr1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
- **Polkadot (DOT)**: `1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
- **Solana (SOL)**: `xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`

##### 🥈 Silver Supporter - $25/month
- All Bronze benefits
- Your logo/brand in the README
- Direct communication channel with maintainers
- Feature request priority

**Payment Options:** Same as Bronze (see above)

##### 🥇 Gold Supporter - $50/month
- All Silver benefits
- Custom feature development (within scope)
- Dedicated support channel
- Quarterly roadmap consultation

**Payment Options:** Same as Bronze (see above)

##### 💎 Platinum Supporter - $100+/month
- All Gold benefits
- Co-maintainer status (if desired)
- Custom branding options
- Direct influence on project direction

**Payment Options:** Same as Bronze (see above)

#### One-Time Donations

We also accept one-time donations of any amount. All supporters will be recognized in our contributors section.

**Payment Addresses:**
- **Bitcoin (BTC)**: `bc1qxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
- **Ethereum (ETH)**: `0x1234567890123456789012345678901234567890`
- **USDT (TRC20)**: `Txxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
- **USDT (ERC20)**: `0x1234567890123456789012345678901234567890`
- **USDC (ERC20)**: `0x1234567890123456789012345678901234567890`
- **Litecoin (LTC)**: `Lxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
- **Dogecoin (DOGE)**: `Dxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
- **Binance Coin (BNB)**: `0x1234567890123456789012345678901234567890`
- **Cardano (ADA)**: `addr1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
- **Polkadot (DOT)**: `1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
- **Solana (SOL)**: `xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`

*Note: Please contact us at support@godbadmin.com after making a payment to ensure proper recognition.*

### Community Support

- **GitHub Issues**: [Report bugs or request features](https://github.com/GoDBAdmin/GoDBAdmin/issues)
- **GitHub Discussions**: [Ask questions and share ideas](https://github.com/GoDBAdmin/GoDBAdmin/discussions)
- **Email**: support@godbadmin.com

## Roadmap

- [ ] Database export/import functionality
- [ ] Multi-database connection support
- [ ] Query history and favorites
- [ ] User management and permissions
- [ ] Database backup/restore tools
- [ ] Performance monitoring dashboard
- [ ] Custom themes and UI customization
- [ ] Mobile app (iOS/Android)
- [ ] Plugin system for extensions

## Security

Security is a top priority. If you discover a security vulnerability, please email security@godbadmin.com instead of using the issue tracker.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Go Fiber](https://gofiber.io/) - Express-inspired web framework
- UI powered by [React](https://react.dev/) and [Tailwind CSS](https://tailwindcss.com/)
- SQL Editor uses [Monaco Editor](https://microsoft.github.io/monaco-editor/)
- Thanks to all [contributors](https://github.com/GoDBAdmin/GoDBAdmin/graphs/contributors) who help make this project better

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=GoDBAdmin/GoDBAdmin&type=Date)](https://star-history.com/#GoDBAdmin/GoDBAdmin&Date)

---

<div align="center">

**Made with ❤️ by the GoDBAdmin community**

[⭐ Star us on GitHub](https://github.com/GoDBAdmin/GoDBAdmin) • [🐛 Report Bug](https://github.com/GoDBAdmin/GoDBAdmin/issues) • [💡 Request Feature](https://github.com/GoDBAdmin/GoDBAdmin/issues)

</div>
