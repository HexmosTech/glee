
set DOWNLOAD_URL=https://github.com/HexmosTech/glee/releases/latest/download/glee_windows.exe
set NEW_NAME=glee.exe

:: Download glee_windows.exe as glee.exe
powershell -Command "(New-Object Net.WebClient).DownloadFile('%DOWNLOAD_URL%', '%NEW_NAME%')"

:: Move glee.exe to C:\Windows\System32
move glee.exe C:\Windows\System32

:: Add your additional commands here
echo glee.exe has been moved to C:\Windows\System32


set PATH=%PATH%;%~dp0
