$url = "https://example.com"

try {
    $response = Invoke-WebRequest -Uri $url -UseBasicParsing -TimeoutSec 10
    if ($response.StatusCode -eq 200) {
        Write-Host "Site accessible"
    } else {
        Write-Host "Site inaccessible, code: $($response.StatusCode)"
    }
} catch {
    Write-Host "Erreur : Site inaccessible"
}
