@echo off
:poor mans build script

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
copy hello_world %DEST%
git add .
git commit -m "."
git branch part1

:END