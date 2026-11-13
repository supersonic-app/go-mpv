//go:build ios && !iossimulator && !mpvmediakit

package mpv

/*
#cgo CFLAGS: -I/opt/xcframework/Libmpv.xcframework/ios-arm64/Headers
#cgo LDFLAGS: -L/opt/xcframework/Libmpv.xcframework/ios-arm64 -lmpv-combined
#cgo LDFLAGS: -lc++ -lz -lbz2 -liconv
#cgo LDFLAGS: -framework AVFoundation -framework AVFAudio -framework AudioToolbox -framework CoreGraphics -framework CoreMedia -framework CoreText -framework CoreVideo -framework VideoToolbox -framework Metal
*/
import "C"
