
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
            Write-Host "=> Creating apt repo folder"
            mkdir ~/aptsumocli/
            Write-Host "=> Syncing aptsumocli S3 bucket locally"
            aws s3 sync s3://aptsumocli ~/aptsumocli/
            Write-Host "=> Creating pools directory"
            mkdir -p ~/aptsumocli/pool/main
            Write-Host "=> Moving deb package to local apt repo"
            mv ~/deb/sumocli_$version-1_$goarchitecture.deb ~/aptsumocli/pool/main/sumocli_$version-1_$goarchitecture.deb
            Write-Host "=> Creating packages directory"
            mkdir -p ~/aptsumocli/dists/stable/main/binary-$goarchitecture
            Write-Host "=> Removing old packages file"
            rm dists/stable/main/binary-$goarchitecture/Packages
            rm dists/stable/main/binary-$goarchitecture/Packages.gz
            Write-Host "=> Generating new packages file"
            dpkg-scanpackages --arch $goarchitecture ~/aptsumocli/pool/ > ~/aptsumocli/dists/stable/main/binary-$goarchitecture/Packages
            Write-Host "=> Compressing packages file"
            cat ~/aptsumocli/dists/stable/main/binary-$goarchitecture/Packages | gzip -9 > ~/aptsumocli/dists/stable/main/binary-$goarchitecture/Packages.gz
            Write-Host "=> Creating release file"
            $releaseFile = @"
Origin: apt.sumocli.app
Suite: stable
Codename: stable
Version: $version
Architectures: amd64
Components: main
Description: Sumocli is a CLI application written in Go that allows you to manage your Sumo Logic tenancy from the command line.
Date: $(Get-Date)

MD5Sum:
$(pwsh "$PSScriptRoot/create-debianrelease.ps1" -algorithm MD5)
SHA1:

SHA256:
"@
            Write-Host $releaseFile


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
