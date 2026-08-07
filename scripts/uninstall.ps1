#Requires -Version 5.0

$Binary     = "cmdock"
$InstallDir = "$env:LOCALAPPDATA\cmdock"
$DbPath     = "$env:USERPROFILE\.cmdock.db"
$Profile    = "$env:USERPROFILE\Documents\PowerShell\Microsoft.PowerShell_profile.ps1"

function Write-Success { param($msg) Write-Host "✓ $msg" -ForegroundColor Green }
function Write-Warn    { param($msg) Write-Host "! $msg" -ForegroundColor DarkGray }

Write-Host ""
Write-Host "▸ cmdock uninstaller" -ForegroundColor White
Write-Host ""

$answer = Read-Host "This will remove the cmdock binary, database, and shell hooks. Continue? [y/N]"
if ($answer -notmatch '^[yY]') {
    Write-Host "Cancelled."
    exit 0
}
Write-Host ""

# --- 1. remove binary ---
$BinaryPath = Join-Path $InstallDir "$Binary.exe"
if (Test-Path $BinaryPath) {
    Remove-Item -Path $InstallDir -Recurse -Force
    Write-Success "Removed binary: $BinaryPath"
} else {
    Write-Warn "No binary found at $InstallDir"
}

# --- 2. remove database ---
if (Test-Path $DbPath) {
    Remove-Item -Path $DbPath -Force
    Write-Success "Database deleted"
} else {
    Write-Warn "No database found"
}

# --- 3. remove from PATH ---
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -like "*$InstallDir*") {
    $newPath = ($userPath -split ';' | Where-Object { $_ -ne $InstallDir }) -join ';'
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Success "Removed $InstallDir from PATH"
}

# --- 4. strip PowerShell profile hook ---
if (Test-Path $Profile) {
    $content = Get-Content $Profile -Raw
    if ($content -match '(?s)# >>> cmdock start >>>.*?# <<< cmdock end <<<') {
        $cleaned = $content -replace '(?s)\r?\n?# >>> cmdock start >>>.*?# <<< cmdock end <<<\r?\n?', "`n"
        Set-Content -Path $Profile -Value $cleaned
        Write-Success "Shell integration removed"
    }
}

Write-Host ""
Write-Host "✓ cmdock fully uninstalled" -ForegroundColor Green
Write-Host "Restart PowerShell for changes to take effect." -ForegroundColor DarkGray
Write-Host ""