#
# Apache License 2.0
# Copyright (c) 2026 OTMC Softwares.
# Contributors: Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
#

$global:TOP = Get-Location


function e   { Set-Location $TOP; & "$TOP\env.ps1" @args }
function p   { Set-Location $TOP; & "$TOP\scripts\push.ps1" @args }
function tag { Set-Location $TOP; & "$TOP\scripts\tag.ps1" @args }


function Show-Help {
    Write-Host ""
    Write-Host "  ╔═══════════════════════════════════════════════════════════════╗" -ForegroundColor Cyan
    Write-Host "  ║             OTMC Logger - Environment Help                    ║" -ForegroundColor Cyan
    Write-Host "  ╚═══════════════════════════════════════════════════════════════╝" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "  COMMANDS:" -ForegroundColor Yellow
    Write-Host "    e                     Load environment"
    Write-Host "    p                     Push changes to remote repository"
    Write-Host "    tag b <tag>           Create tag at current HEAD"
    Write-Host "    tag r <tag>           Restore current branch to tag (force push)"
    Write-Host ""
    Write-Host "  EXAMPLES:" -ForegroundColor Yellow
    Write-Host "    tag b v0.1.1          Create tag v0.1.1 at HEAD"
    Write-Host "    tag r v0.1.1          Restore branch to tag v0.1.1"
    Write-Host ""
}

Show-Help
Write-Host ""
Write-Host "   >>> Environment Loaded on Windows!" -ForegroundColor Blue
Write-Host "   >>> Source directory: '$TOP'" -ForegroundColor Blue
Write-Host ""

