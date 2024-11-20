#!pwsh

$CurrWorkDir = Get-Location
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
Set-Location $ScriptDir -ErrorAction Stop
Set-Location ..

Get-ChildItem -Path . -Recurse -Filter "*.go" | ForEach-Object {
    $file = $_
    Write-Output $file.FullName
    Get-Content $file.FullName | Select-String -Pattern $args[0] -CaseSensitive
}

Set-Location $CurrWorkDir
