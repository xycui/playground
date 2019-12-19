FROM mcr.microsoft.com/windows/servercore:1909

RUN powershell -Command \
    Set-ExecutionPolicy RemoteSigned -scope CurrentUser -Force;\
    iwr -useb get.scoop.sh | iex;\
    scoop install git;scoop bucket add extras;scoop update;\
    scoop install wget;

RUN powershell -Command \
    wget -O vc_redist.x86.exe "https://aka.ms/vs/16/release/vc_redist.x86.exe";\
    vc_redist.x86.exe /install /quiet /norestart;\
    Remove-Item vc_redist.x86.exe
    

RUN powershell -Command \
    wget -O vc_redist.x64.exe "https://aka.ms/vs/16/release/vc_redist.x86.exe";\
    vc_redist.x64.exe /install /quiet /norestart;\
    Remove-Item vc_redist.x64.exe

CMD [ "powershell" ]