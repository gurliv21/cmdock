#Requires -Version 5.0

$ErrorActionPreference = "Stop"

$Repo      = "gurliv21/cmdock"
$Binary    = "cmdock"
$InstallDir = "$env:LOCALAPPDATA\cmdock"

function Write-Info    { param($msg) Write-Host $msg -ForegroundColor Cyan }
function Write-Success { param($msg) Write-Host "✓ $msg" -ForegroundColor Green }
function Write-ErrorMsg{ param($msg) Write-Host "✗ $msg" -ForegroundColor Red }

Write-Host ""
Write-Host @"
  ██████╗███╗   ███╗██████╗  ██████╗  ██████╗██╗  ██╗
 ██╔════╝████╗ ████║██╔══██╗██╔═══██╗██╔════╝██║ ██╔╝
 ██║     ██╔████╔██║██║  ██║██║   ██║██║     █████╔╝
 ██║     ██║╚██╔╝██║██║  ██║██║   ██║██║     ██╔═██╗
 ╚██████╗██║ ╚═╝ ██║██████╔╝╚██████╔╝╚██████╗██║  ██╗
  ╚═════╝╚═╝     ╚═╝╚═════╝  ╚═════╝  ╚═════╝╚═╝  ╚═╝
"@ -ForegroundColor Cyan
Write-Host ""

# --- detect architecture ---
$arch = $env:PROCESSOR_ARCHITECTURE
switch ($arch) {
    "AMD64" { $Arch = "amd64" }
    "ARM64" { $Arch = "arm64" }
    default {
        Write-ErrorMsg "Unsupported architecture: $arch"
        exit 1
    }
}

Write-Info "Detected: windows/$Arch"

# --- get latest version ---
Write-Host "Fetching latest release..."
try {
    $release = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest"
    $Version = $release.tag_name
} catch {
    Write-ErrorMsg "Failed to fetch latest version"
    exit 1
}

if (-not $Version) {
    Write-ErrorMsg "Failed to fetch latest version"
    exit 1
}

Write-Success "Latest version: $Version"

$Archive = "${Binary}_windows_${Arch}.zip"
$Url = "https://github.com/$Repo/releases/download/$Version/$Archive"

$TmpDir = Join-Path $env:TEMP "cmdock-install-$(Get-Random)"
New-Item -ItemType Directory -Path $TmpDir | Out-Null

try {
    Write-Host "Downloading..."
    $ArchivePath = Join-Path $TmpDir $Archive
    Invoke-WebRequest -Uri $Url -OutFile $ArchivePath -UseBasicParsing
    Write-Success "Downloaded"

    Expand-Archive -Path $ArchivePath -DestinationPath $TmpDir -Force

    $BinaryPath = Join-Path $TmpDir "$Binary.exe"
    if (-not (Test-Path $BinaryPath)) {
        Write-ErrorMsg "Binary not found in archive"
        exit 1
    }

    if (-not (Test-Path $InstallDir)) {
        New-Item -ItemType Directory -Path $InstallDir | Out-Null
    }

    $DestPath = Join-Path $InstallDir "$Binary.exe"
    Move-Item -Path $BinaryPath -Destination $DestPath -Force

    # --- add to PATH if missing ---
    $userPath = [Environment]::GetEnvironmentVariable("Path", "User")
    if ($userPath -notlike "*$InstallDir*") {
        [Environment]::SetEnvironmentVariable("Path", "$userPath;$InstallDir", "User")
        $env:Path += ";$InstallDir"
        Write-Success "Added $InstallDir to PATH"
    }

    Write-Host ""
    Write-Host "✓ cmdock installed successfully" -ForegroundColor Green
    Write-Host ""

    # --- shell init ---
    $answer = Read-Host "Run shell init now to enable command tracking? [Y/n]"
    if ($answer -match '^[nN]') {
        Write-Host ""
        Write-Host "Skipped. Run this later:" -ForegroundColor DarkGray
        Write-Host "  cmdock init"
        Write-Host ""
    } else {
        Write-Host ""
        & $DestPath init
        Write-Success "Shell initialized"
        Write-Host ""
        Write-Host "Restart PowerShell to start using cmdock." -ForegroundColor DarkGray
        Write-Host ""
    }
}
finally {
    Remove-Item -Path $TmpDir -Recurse -Force -ErrorAction SilentlyContinue
}