$algorithms = @(
    "MD5",
    "SHA1",
    "SHA256"
)
$hashContent = @()

foreach ($i in $algorithms) {
    if ($i -eq "MD5") {
        $hashContent = $hashContent + "MD5Sum:"
    } else {
        algorithm = $i + ":"
        $hashContent = $hashContent + $algorithm
    }
}

Write-Host $hashContent

#Get-ChildItem -Path ~/aptsumocli/dists/stable/main -recurse -File | Foreach-Object {
#    $hash = Get-FileHash $_.FullName -Algorithm $algorithm
#    $relativePath = Resolve-Path -Path ~/aptsumocli/dists/stable/ | Select-Object -ExpandProperty Path
#    $fileName = $_.FullName.Replace($relativePath, "")
#    $data = $hash.Hash + $_.Length + $fileName
#    $data
#}
