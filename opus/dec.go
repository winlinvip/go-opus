// The MIT License (MIT)
//
// Copyright (c) 2016 winlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// The opus decoder, to decode the encoded opus frame to PCM samples.
package opus

/*
#cgo CFLAGS: -I${SRCDIR}/../opus-lib/objs/include/opus
#cgo LDFLAGS: ${SRCDIR}/../opus-lib/objs/lib/libopus.a -lm
#include "opus.h"

typedef struct {
	OpusDecoder* dec;
} opusdec_t;

static int opusdec_init(opusdec_t* h, int sample_rate, int channels) {
	h->dec = 0;

	int err = OPUS_OK;
	h->dec = opus_decoder_create(sample_rate, channels, &err);
	return err;
}

static void opusdec_close(opusdec_t* h) {
	if (h->dec) {
		opus_decoder_destroy(h->dec);
		h->dec = 0;
	}
}

static int opusdec_decode(opusdec_t* h, char* data, int len, char* pcm, int frame_size, int fec) {
	return opus_decode(h->dec, (const unsigned char*)data, (opus_int32)len,
		(opus_int16*)pcm, frame_size, fec);
}

static int opusdec_decode_float(opusdec_t* h, char* data, int len, float* pcm, int frame_size, int fec) {
	return opus_decode_float(h->dec, (const unsigned char*)data, (opus_int32)len,
		(float*)pcm, frame_size, fec);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type OpusDecoder struct {
	m C.opusdec_t
}

func NewOpusDecoder() *OpusDecoder {
	return &OpusDecoder{}
}

func (v *OpusDecoder) Init(sampleRate, channels int) (err error) {
	cSampleRate := C.int(sampleRate)
	cChannels := C.int(channels)

	r := C.opusdec_init(&v.m, cSampleRate, cChannels)

	if int(r) != 0 {
		return fmt.Errorf("init OPUS decoder failed, code is %d", int(r))
	}

	return nil
}

func (v *OpusDecoder) Close() error {
	C.opusdec_close(&v.m)
	return nil
}

func (v *OpusDecoder) Decode(data []byte, pcm []byte, fec bool) (int, error) {
	pData := (*C.char)(unsafe.Pointer(&data[0]))
	pSize := C.int(len(data))
	pPCM := (*C.char)(unsafe.Pointer(&pcm[0]))
	pFrameSize := C.int(len(pcm))
	var pDecodeFEC C.int
	if fec {
		pDecodeFEC = 1
	}

	r := C.opusdec_decode(&v.m, pData, pSize, pPCM, pFrameSize, pDecodeFEC)
	if int(r) > 0 {
		return int(r), nil
	}

	return 0, fmt.Errorf("decode err, code is %d", int(r))
}

func (v *OpusDecoder) DecodeFloat(data []byte, pcm []float32, fec bool) (int, error) {
	pData := (*C.char)(unsafe.Pointer(&data[0]))
	pSize := C.int(len(data))
	pPCM := (*C.float)(unsafe.Pointer(&pcm[0]))
	pFrameSize := C.int(len(pcm))
	var pDecodeFEC C.int
	if fec {
		pDecodeFEC = 1
	}

	r := C.opusdec_decode_float(&v.m, pData, pSize, pPCM, pFrameSize, pDecodeFEC)
	if int(r) > 0 {
		return int(r), nil
	}

	return 0, fmt.Errorf("decode err, code is %d", int(r))
}

const (
	opusOK = 0
	/** One or more invalid/out of range arguments @hideinitializer*/
	opusBadArg = -1
	/** Not enough bytes allocated in the buffer @hideinitializer*/
	opusBufferTooSmall = -2
	/** An internal error was detected @hideinitializer*/
	opusInternalError = -3
	/** The compressed data passed is corrupted @hideinitializer*/
	opusInvalidPacket = -4
	/** Invalid/unsupported request number @hideinitializer*/
	opusUnimplemented = -5
	/** An encoder or decoder structure is invalid or already freed @hideinitializer*/
	opusInvalidState = -6
	/** Memory allocation has failed @hideinitializer*/
	opusAllocFail = -7
)
