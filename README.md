# go-opus

Golang binding for libopus(https://github.com/winlinvip/opus)

## Usage

First, get the source code:

```
go get -d github.com/winlinvip/go-opus
```

Then, compile the opus:

```
cd $GOPATH/src/github.com/winlinvip/go-opus &&
(git clone https://github.com/winlinvip/opus.git opus-lib && cd opus-lib && bash autogen.sh && ./configure --prefix=`pwd`/objs && make && make install) &&
(git clone https://github.com/winlinvip/opusfile.git && cd opusfile &&
	export PKG_CONFIG_PATH=`pwd`/../opus-lib/objs/lib/pkgconfig &&
	bash autogen.sh && ./configure --prefix=`pwd`/objs --disable-http && make && make install) &&
(git clone https://github.com/winlinvip/libopusenc.git && cd libopusenc &&
	bash autogen.sh && ./configure --prefix=`pwd`/objs && make && make install) &&
(git clone https://github.com/winlinvip/opus-tools.git && cd opus-tools &&
	export PKG_CONFIG_PATH=`pwd`/../opus-lib/objs/lib/pkgconfig:`pwd`/../opusfile/objs/lib/pkgconfig:`pwd`/../libopusenc/objs/lib/pkgconfig &&
	bash autogen.sh && ./configure --prefix=`pwd`/objs && make && make install)
```

Done, import and use the package:

* [ExampleOpusDecoder_RAW](opus/example_test.go#L24), decode the aac frame to PCM samples.

There are an example of AAC audio packets in ADTS:

* [avatar aac over ADTS](https://github.com/winlinvip/go-fdkaac/blob/master/doc/adts_data.go), user can use this file to decode to PCM then encode.

To run all examples:

```
cd $GOPATH/src/github.com/winlinvip/go-opus && go test ./...
```

Winlin 2018
