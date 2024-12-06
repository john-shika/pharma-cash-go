#!pwsh

$CurrWorkDir = Get-Location
$ScriptRootDir = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $ScriptRootDir
Set-Location ..

$ZIPFILE = "app.zip"

if (Test-Path $ZIPFILE) {
    Remove-Item $ZIPFILE
}

$FILES = @(
    "app",
    "pkg",
    "main.go",
    "go.mod",
    "go.work",
    "nokowebapi.yaml.example"
)

Compress-Archive -Path $FILES -DestinationPath $ZIPFILE -Force

Set-Location $CurrWorkDir
