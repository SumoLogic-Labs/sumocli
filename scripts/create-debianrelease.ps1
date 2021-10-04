$algorithms = @(
    "MD5",
    "SHA1",
    "SHA256",
    "SHA512"
)
$hashContent = @()

foreach ($i in $algorithms) {
    if ($i -eq "MD5") {
        $hashContent = $hashContent + "MD5Sum:"
    } else {
        $algorithm = $i + ":"
        $hashContent = $hashContent + $algorithm
    }
    Get-ChildItem -Path ~/aptsumocli/dists/stable/main -recurse -File | ForEach-Object {
        $hash = Get-FileHash $_.FullName -Algorithm $i
        $relativePath = Resolve-Path -Path ~/aptsumocli/dists/stable/ | Select-Object -ExpandProperty Path
        $fileName = $_.FullName.Replace($relativePath, "")
        $data = " " + $hash.Hash.ToLower() + " " + $_.Length + " " + $fileName
        $hashContent = $hashContent + $data
    }
}

foreach ($i in $hashContent) {
    Write-Host $i
}
