# GitHub Actions CI/CD

This directory contains GitHub Actions workflows for automated building and deployment of the RTSP Stream application to Docker Hub.

## Workflows

### 1. `docker-build-deploy.yml`

Main workflow that builds and deploys the core RTSP Stream application.

**Triggers:**

-   Push to `main`/`master` branches
-   Git tags starting with `v*`
-   Pull requests to `main`/`master` branches

**Features:**

-   Multi-platform builds (linux/amd64, linux/arm64)
-   Automatic tagging based on Git references
-   Vulnerability scanning with Trivy
-   Build caching for faster builds
-   Security scanning results uploaded to GitHub Security tab

### 2. `docker-build-deploy-management.yml`

Workflow that builds and deploys the management UI version of the application.

**Triggers:**

-   Push to `main`/`master` branches (when management-related files change)
-   Git tags starting with `v*`
-   Pull requests affecting management files

**Features:**

-   Builds the management image with UI
-   Multi-platform support
-   Path-based triggering (only runs when relevant files change)

## Setup Instructions

### 1. Docker Hub Configuration

You need to set up the following secrets in your GitHub repository:

1. Go to your GitHub repository
2. Click on **Settings** → **Secrets and variables** → **Actions**
3. Add the following repository secrets:

-   `DOCKER_USERNAME`: Your Docker Hub username
-   `DOCKER_PASSWORD`: Your Docker Hub access token (recommended) or password

#### How to create a Docker Hub Access Token:

1. Log in to [Docker Hub](https://hub.docker.com)
2. Go to **Account Settings** → **Security**
3. Click **New Access Token**
4. Give it a name (e.g., "GitHub Actions")
5. Copy the token and use it as `DOCKER_PASSWORD`

### 2. Repository Configuration

Make sure your repository has the following:

-   Public or private repository on GitHub
-   Docker Hub repositories:
    -   `mellomateus/rtsp-stream` (for main application)
    -   `mellomateus/rtsp-stream-management` (for management UI)

### 3. Image Tags

The workflows automatically create the following tags:

#### For regular pushes to main/master:

-   `latest`
-   `main-<sha>` or `master-<sha>`

#### For tagged releases (e.g., `v1.2.3`):

-   `1.2.3`
-   `1.2`
-   `1`
-   `v1.2.3`

#### For pull requests:

-   `pr-<number>`

### 4. Security Scanning

The main workflow includes Trivy vulnerability scanning:

-   Scans the built Docker image for security vulnerabilities
-   Results are uploaded to GitHub Security tab
-   Runs only on successful builds (not on PRs)

## Usage

### Deploying a new version

1. **For development builds:**

    ```bash
    git push origin main
    ```

    This will create a `latest` tag on Docker Hub.

2. **For release builds:**
    ```bash
    git tag v1.0.0
    git push origin v1.0.0
    ```
    This will create version-specific tags: `1.0.0`, `1.0`, `1`, and `v1.0.0`.

### Running the deployed images

```bash
# Latest development build
docker run -p 8080:8080 mellomateus/rtsp-stream:latest

# Specific version
docker run -p 8080:8080 mellomateus/rtsp-stream:1.0.0

# Management UI version
docker run -p 8080:8080 mellomateus/rtsp-stream-management:latest
```

## Troubleshooting

### Build Failures

1. **Docker Hub authentication fails:**

    - Check that `DOCKER_USERNAME` and `DOCKER_PASSWORD` secrets are set correctly
    - Verify the Docker Hub access token is still valid

2. **Build context issues:**

    - Ensure all necessary files are committed to the repository
    - Check that Dockerfile paths are correct

3. **Multi-platform build issues:**
    - The workflow uses Docker Buildx for multi-platform builds
    - If issues persist, you can modify the workflow to build for single platform

### Security Scan Failures

If Trivy scanning fails:

-   Check the GitHub Security tab for detailed vulnerability reports
-   The scan failure won't prevent the Docker image from being built and pushed
-   Consider updating base images or dependencies to address vulnerabilities

## Customization

### Changing Docker Hub Repository

Edit the `DOCKER_IMAGE` environment variable in the workflow files:

```yaml
env:
    DOCKER_IMAGE: your-dockerhub-username/your-repository-name
```

### Adding Different Platforms

Modify the `platforms` field in the build step:

```yaml
platforms: linux/amd64,linux/arm64,linux/arm/v7
```

### Modifying Triggers

Adjust the `on` section to change when workflows run:

```yaml
on:
    push:
        branches: [main, develop]
    release:
        types: [published]
```
