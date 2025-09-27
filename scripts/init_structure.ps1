# init_structure.ps1

# New-Item -ItemType Directory -Name "irit-backend"
# Set-Location -Path "irit-backend"

$directories = @(
    "cmd\api",
    "internal\auth",
    "internal\user",
    "internal\driver",
    "internal\booking",
    "pkg\database",
    "configs"
    "scripts"
)

foreach ($dir in $directories) {
    New-Item -ItemType Directory -Path $dir -Force | Out-Null
}