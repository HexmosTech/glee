:: Use GitHub API to get the latest release version
for /f "delims=" %%v in ('powershell -Command "(Invoke-RestMethod -Uri 'https://api.github.com/repos/HexmosTech/glee/releases/latest').tag_name"') do set LATEST_VERSION=%%v

:: Set URLs and file names
set DOWNLOAD_URL=https://github.com/HexmosTech/glee/releases/latest/download
set ZIP_FILE_NAME=glee-%LATEST_VERSION%-windows-%PROCESSOR_ARCHITECTURE%.zip
set MD5_FILE_NAME=%ZIP_FILE_NAME%.md5
set NEW_NAME=glee.exe

echo Downloading %ZIP_FILE_NAME%...
:: Download the appropriate zip file based on the system architecture
powershell -Command "(New-Object Net.WebClient).DownloadFile('%DOWNLOAD_URL%/%ZIP_FILE_NAME%', '%ZIP_FILE_NAME%')"

:: Check if the download was successful
if not exist "%ZIP_FILE_NAME%" (
    echo Failed to download %ZIP_FILE_NAME%. Please download the file manually from:
    echo %DOWNLOAD_URL%/
) else (
    echo %ZIP_FILE_NAME% downloaded successfully.
    
    echo Extracting %ZIP_FILE_NAME%...
    :: Unzip the downloaded file
    powershell -Command "Expand-Archive -Path '%ZIP_FILE_NAME%' -DestinationPath ."
    
    :: Get the glee.exe from the unzipped directory
    set EXTRACTED_DIR=glee-latest-windows-%PROCESSOR_ARCHITECTURE%
    set GLEE_EXE_PATH=%EXTRACTED_DIR%\glee.exe

    :: Check if glee.exe exists
    if not exist "%GLEE_EXE_PATH%" (
        echo Failed to find glee.exe in the extracted directory.
    ) else (
        echo glee.exe found in %EXTRACTED_DIR%.
        
        :: Move glee.exe to the desired location with the specified name
        move "%GLEE_EXE_PATH%" "%NEW_NAME%"
        
        echo glee.exe moved successfully.
    )

    :: Clean up: remove the downloaded zip file and the extracted directory
    del "%ZIP_FILE_NAME%"
    rmdir /s /q "%EXTRACTED_DIR%"
)

echo Downloading configuration file...

:: Download .glee.toml file
set TOML_URL=https://raw.githubusercontent.com/HexmosTech/glee/main/.glee.toml
set TOML_DESTINATION=%USERPROFILE%\.glee.toml
powershell -Command "(New-Object Net.WebClient).DownloadFile('%TOML_URL%', '%TOML_DESTINATION%')"
echo .glee.toml file downloaded successfully.
echo Add the Ghost Configuration in %USERPROFILE%\.glee.toml file
