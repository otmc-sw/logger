#
# Apache License 2.0
# Copyright (c) 2026 OTMC Softwares.
# Contributors: Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
#

Set-Location -Path $PSScriptRoot/..

$LicenseHeader = @'
/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
'@

$IgnoredDirs = @(
    "/sqlc/",
    "/node_modules/",
    "/dist/",
    "/data/"
)

$KeepPatterns = @(
    "OTMC License",
    "Copyright",
    "Trung Ng",
    "TODO: ",
    "go:embed",
    "Contributors:"
)

function Is-IgnoredPath($path) {
    foreach ($dir in $IgnoredDirs) {
        if ($path -like "*$dir*") { return $true }
    }
    return $false
}

function Has-Header($content) {
    return $content -match '@License Apache License 2.0'
}

function Remove-Comments($content) {
    $lines = $content -split "`n"
    $result = @()
    
    foreach ($line in $lines) {
        # Skip full-line //
        if ($line -match '^\s*//') {
            # Keep if whitelisted
            $keep = $false
            foreach ($pattern in $KeepPatterns) {
                if ($line -match [regex]::Escape($pattern)) {
                    $keep = $true
                    break
                }
            }
            if ($keep) { $result += $line }
            continue
        }
        
        # Remove block comments
        if ($line -match '^\s*/\*') {
            $keep = $false
            foreach ($pattern in $KeepPatterns) {
                if ($line -match [regex]::Escape($pattern)) {
                    $keep = $true
                    break
                }
            }
            if (!$keep) { continue }
        }
        
        # Remove inline //
        if ($line -match '//') {
            $cleaned = $line -replace '\s*//.*$', ''
            $result += $cleaned.TrimEnd()
            continue
        }
        
        $result += $line
    }
    
    return ($result -join "`n")
}

# Get all .go files
$files = Get-ChildItem -Path "." -Recurse -Filter "*.go" |
    Where-Object { -not (Is-IgnoredPath $_.FullName) }

Write-Host "Found $($files.Count) Go files" -ForegroundColor Cyan

$index = 0
foreach ($file in $files) {
    $index++
    Write-Host "[$index/$($files.Count)] Processing: $($file.FullName)"
    
    $content = Get-Content $file.FullName -Raw
    
    # Add header if missing
    if (-not (Has-Header $content)) {
        $content = "$LicenseHeader`n$content"
        Write-Host "  → Added license header" -ForegroundColor Green
    }
    
    # Remove comments
    $newContent = Remove-Comments $content
    
    # Write back if changed
    if ($newContent -ne $content) {
        [System.IO.File]::WriteAllText($file.FullName, $newContent, [System.Text.Encoding]::UTF8)
        Write-Host "  → Removed comments" -ForegroundColor Yellow
    }
}

Write-Host "`nDone!" -ForegroundColor Cyan