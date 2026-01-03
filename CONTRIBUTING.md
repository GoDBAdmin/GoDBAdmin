# Contributing to GoDBAdmin

Thank you for your interest in contributing to GoDBAdmin! This document provides guidelines and instructions for contributing.

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for everyone.

## How to Contribute

### Reporting Bugs

1. Check if the bug has already been reported in [Issues](https://github.com/GoDBAdmin/GoDBAdmin/issues)
2. If not, create a new issue with:
   - Clear title and description
   - Steps to reproduce
   - Expected vs actual behavior
   - Environment details (OS, Go version, Node version, etc.)
   - Screenshots if applicable

### Suggesting Features

1. Check if the feature has already been suggested
2. Create a new issue with:
   - Clear description of the feature
   - Use case and benefits
   - Possible implementation approach (if you have ideas)

### Pull Requests

1. **Fork the repository**
2. **Create a feature branch**:
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Make your changes**:
   - Follow the coding style guidelines
   - Write clear commit messages
   - Add tests if applicable
   - Update documentation
4. **Test your changes**:
   ```bash
   # Test backend
   cd backend
   go test ./...
   
   # Test frontend
   cd frontend
   npm test
   ```
5. **Commit your changes**:
   ```bash
   git commit -m "Add: Description of your changes"
   ```
6. **Push to your fork**:
   ```bash
   git push origin feature/amazing-feature
   ```
7. **Create a Pull Request**:
   - Provide a clear title and description
   - Reference any related issues
   - Wait for review and feedback

## Development Setup

See the [Running Guide](docs/RUN.md) for detailed development setup instructions.

## Coding Guidelines

### Go (Backend)

- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use `gofmt` to format code
- Write meaningful comments for exported functions
- Keep functions focused and small
- Handle errors explicitly

### TypeScript/React (Frontend)

- Follow React best practices
- Use TypeScript for type safety
- Follow the existing component structure
- Use Tailwind CSS for styling
- Keep components small and reusable

## Commit Message Format

Use clear, descriptive commit messages:

```
Add: Feature description
Fix: Bug description
Update: Change description
Refactor: Refactoring description
Docs: Documentation update
```

## Testing

- Write tests for new features
- Ensure all tests pass before submitting PR
- Add integration tests for API endpoints
- Test UI changes in multiple browsers

## Documentation

- Update relevant documentation files
- Add comments for complex logic
- Update README if needed
- Add examples for new features

## Questions?

Feel free to:
- Open an issue for questions
- Start a discussion in GitHub Discussions
- Contact maintainers

Thank you for contributing to GoDBAdmin! 🎉

