# Contributing to Our Project

## Welcome! 👋

We're thrilled that you're interested in contributing to our project. This document provides guidelines to help you get started and make the contribution process smooth and enjoyable.

## Code of Conduct

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project, you agree to abide by its terms.

## How Can You Contribute?

There are many ways to contribute:

1. **Reporting Bugs** 🐞
   - Use the GitHub Issues section
   - Check existing issues before creating a new one
   - Provide a clear and descriptive title
   - Include detailed steps to reproduce the issue
   - Share your environment details (OS, version, etc.)

2. **Suggesting Enhancements** 💡
   - Open a GitHub Issue
   - Clearly describe the proposed enhancement
   - Provide context and rationale
   - Include potential implementation details if possible

3. **Pull Requests** 🚀
   - Fork the repository
   - Create a new branch for your contribution
   - Follow our coding standards
   - Write clear, concise commit messages
   - Include tests for new features or bug fixes

## Development Setup

### Prerequisites

- Go (version 1.23.2 or later)
- Git
- Your favorite code editor

### Local Development

1. Fork the repository
2. Clone your fork
```bash
git clone https://github.com/your-username/repository-name.git
cd repository-name
```

3. Create a new branch
```bash
git checkout -b feature/your-feature-name
```

4. Install dependencies
```bash
go mod download
```

5. Run tests
```bash
go test ./...
```

## Commit Message Convention

We follow the Conventional Commits specification:

```
<type>(<optional scope>): <description>

<optional body>

<optional footer>
```

### Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code formatting
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

### Example Commit Messages:
```
feat(auth): add login functionality

fix(database): resolve connection leak issue

docs: update README with new installation steps
```

## Pull Request Process

1. Ensure your code passes all tests
2. Update documentation if needed
3. Add yourself to the CONTRIBUTORS.md file (optional)
4. Submit the pull request with a clear title and description

### Pull Request Checklist
- [ ] I have tested my changes
- [ ] I have updated the documentation
- [ ] My code follows the project's style guidelines
- [ ] I have added tests for my changes

## Code Review Process

- All submissions require review
- We may ask for modifications to meet our quality standards
- Be patient and responsive to feedback

## Reporting Security Issues

If you discover a security vulnerability, please send an email to [jhasundarm@gmail.com] instead of creating a public issue.

## Questions?

If you have any questions, feel free to:
- Open an issue with the "question" label
- Ask on our community discussion forum
- Reach out to the maintainers directly

## Thank You! 🎉

Your contributions make open-source amazing. We appreciate your help in making this project better.

---

**Last Updated:** December 2024