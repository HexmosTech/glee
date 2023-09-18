
set DOWNLOAD_URL=https://github.com/HexmosTech/glee/releases/latest/download/glee_windows.exe
set NEW_NAME=glee.exe

echo Downloading glee.exe...
:: Download glee_windows.exe as glee.exe
powershell -Command "(New-Object Net.WebClient).DownloadFile('%DOWNLOAD_URL%', '%NEW_NAME%')"

:: Check if the download was successful
if not exist "%NEW_NAME%" (
    echo Failed to download glee.exe. Please download the executable manually from:
    echo https://github.com/HexmosTech/glee/releases/latest
) else (
    echo glee.exe downloaded successfully.
)


echo Downloading configuration file...

:: Download .glee.toml file
set TOML_URL=https://raw.githubusercontent.com/HexmosTech/glee/main/.glee.toml
set TOML_DESTINATION=%USERPROFILE%\.glee.toml
powershell -Command "(New-Object Net.WebClient).DownloadFile('%TOML_URL%', '%TOML_DESTINATION%')"
echo .glee.toml file downloaded successfully.
echo Add the Ghost Configuration in %USERPROFILE%\.glee.toml file
