:poor mans build script

@echo off
setlocal ENABLEDELAYEDEXPANSION


if "%1"=="" (
    echo USAGE: build target
    goto END
)
set DEST=%1
if exist %DEST% (
    echo about to delete %DEST%
    pause
    rd %DEST% /q/s
) else (
    echo what %DEST%.
)
md %DEST%
pushd %DEST%
git init
popd
set part=0
for /F "tokens=*" %%A in (index.txt) do (
    copy %%A %DEST%
    pushd %DEST%
    git add .
    git commit -m "Part !part!: %%A"
    git branch part!part!
    set /a part+=1
    popd
)
:END