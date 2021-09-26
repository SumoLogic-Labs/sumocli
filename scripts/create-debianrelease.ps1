
param (
    [string]$algorithm
)

Get-ChildItem -Path ~/aptsumocli/dists/stable/main -recurse -File | Foreach-Object {
    $hash = Get-FileHash $_.FullName -Algorithm $algorithm
    echo $hash.Hash $_.Length $_.FullName
}
