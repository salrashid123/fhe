FROM mcr.microsoft.com/dotnet/core/sdk:3.1 as build

RUN apt-get update && apt-get install -y git cmake build-essential

WORKDIR /apps
RUN git clone https://github.com/microsoft/SEAL.git

RUN cd SEAL && cmake . -DBUILD_SHARED_LIBS=ON  -DSEAL_BUILD_SEAL_C=ON && make && make install

ADD *.cs .
ADD RideShare.csproj .
RUN mkdir lib && cp -R SEAL/lib/* lib/ && rm -rf SEAL

RUN dotnet restore
RUN dotnet publish --configuration Release


FROM mcr.microsoft.com/dotnet/core/runtime:3.1

COPY --from=build /apps/bin/Release/netcoreapp3.1/* /
COPY --from=build /apps/lib lib/
ENV LD_LIBRARY_PATH=/:lib/
ENTRYPOINT ["./RideShare"]
