set BUILD_CONFIG=%1

rmdir /Q /S build && mkdir build & cd build && cmake .. -G "Ninja" -DCOMPILER=MINGW -DCMAKE_BUILD_TYPE=%BUILD_CONFIG% && cd ..
cd build && ninja
cd ..