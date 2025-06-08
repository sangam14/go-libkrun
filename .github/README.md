# GitHub Actions Configuration

This directory contains the GitHub Actions workflows and configuration for the go-libkrun project.

## Workflows

### CI Workflow (`.github/workflows/ci.yml`)

The main continuous integration workflow that runs on every push and pull request. It includes:

#### Jobs:
1. **Test Job**: 
   - Runs on Ubuntu and macOS
   - Tests with Go 1.21 and 1.22
   - Tests both regular and EFI builds
   - Builds all example applications

2. **Lint Job**:
   - Runs golangci-lint with comprehensive linting rules
   - Checks both regular and EFI builds
   - Uses custom configuration from `.golangci.yml`

3. **Security Job**:
   - Runs Gosec security scanner
   - Uploads results to GitHub Security tab

4. **Build Matrix Job**:
   - Tests Makefile builds on multiple platforms
   - Uploads build artifacts

#### Build Tags Support:
The workflow properly handles the EFI build tag issue you encountered:
- Regular builds: `go build ./cmd/chroot_vm`
- EFI builds: `go build -tags efi ./cmd/boot_efi`

### Release Workflow (`.github/workflows/release.yml`)

Automated release workflow that triggers on version tags or manual dispatch:

#### Features:
- Cross-platform builds (Linux, macOS)
- Multi-architecture support (amd64, arm64)
- Automatic code signing on macOS
- Creates release archives
- Generates changelog
- Uploads release assets

#### Usage:
```bash
# Create and push a tag to trigger release
git tag v1.0.0
git push origin v1.0.0

# Or trigger manually from GitHub Actions tab
```

## Configuration Files

### `.golangci.yml`
Comprehensive linting configuration that:
- Enables most important linters
- Handles CGO-specific exclusions
- Supports build tags
- Excludes test-specific issues

### `.github/dependabot.yml`
Automated dependency updates for:
- Go modules (weekly)
- GitHub Actions (weekly)

## Issue Templates

### Bug Report (`.github/ISSUE_TEMPLATE/bug_report.md`)
Template for reporting bugs with:
- Environment information
- Build command details
- Error output section

### Feature Request (`.github/ISSUE_TEMPLATE/feature_request.md`)
Template for feature requests with:
- Problem description
- Solution proposal
- Implementation considerations

## Pull Request Template

Comprehensive PR template that ensures:
- Clear description of changes
- Build type impact assessment
- Testing checklist
- Code quality checklist

## Key Features

### Build Tag Handling
The workflows properly handle the `//go:build efi` tag in `pkg/krun/krun-efi.go`:
- Regular tests: `go test ./...`
- EFI tests: `go test -tags efi ./...`
- Regular builds: `go build ./cmd/chroot_vm`
- EFI builds: `go build -tags efi ./cmd/boot_efi`

### Dependency Management
- Automatic caching of Go modules
- Installation of libkrun dependencies
- Cross-platform dependency handling

### Security
- Gosec security scanning
- SARIF upload to GitHub Security
- Dependabot for dependency updates

### Artifacts
- Build artifacts uploaded for each platform
- Release binaries for multiple architectures
- Retention policies configured

## Troubleshooting

### Common Issues

1. **libkrun dependencies not found**:
   - The workflows include fallback handling with `|| echo "continuing..."`
   - This allows builds to continue even if libkrun-dev is not available

2. **Build tag issues**:
   - Make sure to use `-tags efi` for EFI builds
   - The workflows test both build configurations

3. **CGO issues**:
   - The workflows set `CGO_ENABLED=1` for release builds
   - Platform-specific dependencies are handled

### Local Testing

To test the workflows locally:

```bash
# Install act (GitHub Actions local runner)
brew install act  # macOS
# or
curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash

# Run the CI workflow
act -j test

# Run specific job
act -j lint
```

## Maintenance

### Updating Workflows
1. Test changes in a fork first
2. Update version numbers in workflows when needed
3. Keep GitHub Actions versions up to date
4. Monitor workflow runs for failures

### Adding New Build Targets
1. Update the build matrix in `ci.yml`
2. Add new targets to `release.yml`
3. Update documentation
