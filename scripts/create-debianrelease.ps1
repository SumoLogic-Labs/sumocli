
param (
    [string]$algorithm
)

$releaseFile = @"
"@

Get-ChildItem -Path ~/aptsumocli/dists/stable/main -recurse -File | Foreach-Object {
    $hash = Get-FileHash $_.FullName -Algorithm $algorithm
    $releaseFile = $hash.Hash $_.Length $_.FullName
}

Write-Host $releaseFile
