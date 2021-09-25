

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
    [string]$version = "DEV",
    [switch]$windows = $false
)

$goarchitecture="amd64"

Write-Host $arm
Write-Host $version
Write-Host $build

# Add goarchitecture if statement

if ($windows -eq $true) {
    Write-Host "=> Compiling Windows " + $goarchitecture + " binary"
    $env:GOOS="windows"; $env:GOARCH=$goarchitecture; go build -ldflags `
    "-X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Version=$version'
      -X 'github.com/SumoLogic-Incubator/sumocli/internal/build.Build=$build'" `
    ./cmd/sumocli
    if ($version -ne "DEV" -and $build -ne "DEV") {
        Write-Host "=> preparing to sign code"
        Write-Host "=> Install azuresigntool"
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
