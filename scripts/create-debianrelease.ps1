
param (
    [string]$algorithm,
    [string]$releaseFileHashBlock
)

Write-Host $releaseFileHashBlock

Get-ChildItem -Path ~/aptsumocli/dists/stable/main -recurse -File | Foreach-Object {
    $hash = Get-FileHash $_.FullName -Algorithm $algorithm
    Write-Host $hash.Hash $_.Length $_.FullName
}
