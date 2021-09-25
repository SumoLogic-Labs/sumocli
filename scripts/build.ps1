

param (
    [switch]$arm = $false,
    [string]$build = "DEV",
    [string]$keyVaultCertificate = "",
    [string]$keyVaultClientId = "",
    [string]$keyVaultClientSecret = "",
    [string]$keyvaulttenantId = "",
    [string]$keyVaultUrl = "",
    [switch]$linux = $false,
    [switch]$macos = $false,
    [string]$maintainer = "kyle@thepublicclouds.com",
    [switch]$release = $false,
    [string]$version = "DEV",
    [switch]$windows = $false
)

$goarchitecture="amd64"

Write-Host $arm
Write-Host $version
Write-Host $build

# Add goarchitecture if statement

if ($linux -eq $true) {
    Write-Host "Compiling Linux $goarchitecture binary"
    $env:GOOS="linux"; $env:GOARCH=$goarchitecture; go build -ldflags `
    "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
      -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" `
    ./cmd/sumocli
    if ($version -ne "DEV" -and $build -ne "DEV") {
        Write-Host "=> Creating deb package"
        mkdir -p ~/deb/sumocli_$version-1_amd64/usr/bin
        cp sumocli ~/deb/sumocli_$version-1_amd64/usr/bin
        Write-Host "=> Creating DEBIAN control file"
        mkdir -p ~/deb/sumocli_$version-1_amd64/DEBIAN
        $controlFile = @"
Package: sumocli
Version: $version
Maintainer: $maintainer
Architecture: amd64
Homepage: https://github.com/SumoLogic-Incubator/sumocli
Description: Sumocli is a CLI application written in Go that allows you to manage your Sumo Logic tenancy from the command line.
"@
        Set-Content -Path ~/deb/sumocli_$version-1_amd64/DEBIAN/control -Value $controlFile
        Write-Host "=> Building deb package"
        dpkg --build ~/deb/sumocli_$version-1_$goarchitecture
        if ($release -eq $true) {
            aws s3 sync s3://aptsumocli ~/aptsumocli/
            # Sync the contents of the apt s3 bucket locally
            # Copy the deb file into /pool/main
            # Generate a new packages file
            # Generate a new releases file
            # Sync contents of repo back to the S3 bucket
        }
    }
}

if ($windows -eq $true) {
    Write-Host "=> Compiling Windows $goarchitecture binary"
    $env:GOOS="windows"; $env:GOARCH=$goarchitecture; go build -ldflags `
    "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
      -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" `
    ./cmd/sumocli
    if ($version -ne "DEV" -and $build -ne "DEV") {
        Write-Host "=> preparing to sign code"
        Write-Host "=> Installing azuresigntool"
        dotnet tool install --global AzureSignTool --version 3.0.0
        Write-Host "=> Signing Windows binary with Azure Key Vault"
        azuresigntool sign --description-url "https://github.com/SumoLogic-Incubator/sumocli" --file-digest sha256 `
        --azure-key-vault-url $keyVaultUrl `
        --azure-key-vault-client-id $keyVaultClientId `
        --azure-key-vault-client-secret $keyVaultClientSecret `
        --azure-key-vault-certificate $keyVaultCertificate `
        --azure-key-vault-tenant-id $keyvaulttenantId `
        --timestamp-rfc3161 http://timestamp.sectigo.com `
        --timestamp-digest sha256 `
        sumocli.exe
    }
}
