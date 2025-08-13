@echo off
setlocal EnableDelayedExpansion

REM Get go module name
for /f "tokens=*" %%i in ('go list -m') do set "goModName=%%i"

REM Set proto paths
set "currentDir=%CD%"
for %%I in ("%currentDir%\..") do set "protoImport=%%~fI"
set "protoIn=%protoImport%\proto"

REM Set proto directories
set n=0
for %%d in (%PROTO_DIRS%) do (
    set "protoDirs[!n!]=%protoIn%\%%d"
    set /a n+=1
)

echo ✅ Found proto directories:
echo   - Import path: %protoImport%
echo   - Source directory: %protoIn%
echo   - Go Mod name: %goModName%

echo Starting protoc compilation...

REM Loop through directories
for /L %%n in (0,1,2) do (
    set "dir=!protoDirs[%%n]!"
    echo Current dir: !dir!
    
    REM Find and process all .proto files
    for /R "!dir!" %%f in (*.proto) do (
        echo   Processing: %%f
        protoc --proto_path=.;%protoImport%;%protoImport%\proto ^
               --go_out=paths=source_relative:.\internal ^
               --go-grpc_out=paths=source_relative:.\internal ^
               --go-http_out=paths=source_relative:.\internal ^
               "%%f"
        
        if !ERRORLEVEL! neq 0 (
            echo ❌ Failed to compile %%f
            exit /b 1
        )
    )
)

REM Replace import paths in generated files
REM Note: Windows version of sed might not be available, using PowerShell instead
powershell -Command ^
    "Get-ChildItem -Path .\internal\proto -Filter *.go -Recurse | ForEach-Object { ^
        (Get-Content $_.FullName) | ForEach-Object { ^
            $_ -replace 'bsi/proto', '%goModName%/internal/proto' ^
        } | Set-Content $_.FullName ^
    }"

echo ✅ All proto files compiled successfully
exit /b 0 