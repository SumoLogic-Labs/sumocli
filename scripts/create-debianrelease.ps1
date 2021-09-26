
param (
    [string]$algorithm
)

Get-ChildItem -Path ~/aptsumocli/dists/stable/main -recurse -File | Foreach-Object {
    $hash = Get-FileHash $_.FullName -Algorithm $algorithm
    $relativePath = Resolve-Path -Path ~/aptsumocli/dists/stable | Select-Object -ExpandProperty Path
    $fileName = $_.FullName.TrimStart(" ", $relativePath + "/")
    Write-Host $hash.Hash $_.Length $fileName
}
