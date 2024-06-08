$ErrorActionPreference = 'Stop'; # stop on all errors
$toolsDir   = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"
$url64      = "https://github.com/intelops/qualitytrace/releases/download/v${env:chocolateyPackageVersion}/qualitytrace_${env:chocolateyPackageVersion}_windows_amd64.tar.gz" # 64bit URL here (HTTPS preferred) or remove - if installer contains both (very rare), use $url

$packageArgs = @{
  packageName   = $env:ChocolateyPackageName
  unzipLocation = $toolsDir

  url64bit      = $url64
  checksum64    = '%checksum%'
  checksumType64= 'sha256'
}

Install-ChocolateyZipPackage @packageArgs
Get-ChocolateyUnzip -FileFullPath "$toolsDir/qualitytrace_${env:chocolateyPackageVersion}_windows_amd64.tar" -Destination $toolsDir
