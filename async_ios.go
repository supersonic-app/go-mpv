//go:build ios && !iossimulator && !mpvmediakit

package mpv

// uses libs from https://github.com/libmpv-ios/mpv-ios

/*
#cgo CFLAGS: -I/opt/xcframework/libmpv-ios/Libmpv.xcframework/ios-arm64/Headers
#cgo LDFLAGS: -L/opt/xcframework/libmpv-ios/Libmpv.xcframework/ios-arm64 -lmpv-combined
#cgo LDFLAGS: -lc++ -lz -lbz2 -liconv
#cgo LDFLAGS: -framework AVFoundation -framework AVFAudio -framework AudioToolbox -framework CoreGraphics -framework CoreMedia -framework CoreText -framework CoreVideo -framework VideoToolbox -framework Metal
*/
import "C"
