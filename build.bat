@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

echo 清理旧构建…
if exist dist rmdir /s /q dist
if exist build rmdir /s /q build
del /f /q *.spec 2>nul

echo 开始 PyInstaller onedir 打包…
pyinstaller ^
  --name RedBlackTree ^
  --onedir ^
  --windowed ^
  --add-data "static;static" ^
  --add-data "rb_tree;rb_tree" ^
  app.py

if errorlevel 1 (
  echo 打包失败，请检查错误信息。
  pause
  exit /b 1
)

echo 打包完成：dist\RedBlackTree\RedBlackTree.exe
pause
