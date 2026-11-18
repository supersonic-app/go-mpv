//go:build ios && !iossimulator && mpvmediakit

package mpv

// Uses libs from https://github.com/media-kit/libmpv-darwin-build

/*
#cgo CFLAGS: -I/opt/xcframework/media-kit-ios/Mpv.xcframework/ios-arm64/Mpv.framework/Headers
#cgo LDFLAGS: -F/opt/xcframework/media-kit-ios/Mpv.xcframework/ios-arm64 -framework Mpv
#cgo LDFLAGS: -F/opt/xcframework/media-kit-ios/Avcodec.xcframework/ios-arm64 -framework Avcodec
#cgo LDFLAGS: -F/opt/xcframework/media-kit-ios/Avformat.xcframework/ios-arm64 -framework Avformat
#cgo LDFLAGS: -F/opt/xcframework/media-kit-ios/Avfilter.xcframework/ios-arm64 -framework Avfilter
#cgo LDFLAGS: -F/opt/xcframework/media-kit-ios/Avutil.xcframework/ios-arm64 -framework Avutil
#cgo LDFLAGS: -F/opt/xcframework/media-kit-ios/Swresample.xcframework/ios-arm64 -framework Swresample
#cgo LDFLAGS: -F/opt/xcframework/media-kit-ios/Swscale.xcframework/ios-arm64 -framework Swscale
#cgo LDFLAGS: -F/opt/xcframework/media-kit-ios/Mbedtls.xcframework/ios-arm64 -framework Mbedtls
#cgo LDFLAGS: -F/opt/xcframework/media-kit-ios/Mbedcrypto.xcframework/ios-arm64 -framework Mbedcrypto
#cgo LDFLAGS: -F/opt/xcframework/media-kit-ios/Mbedx509.xcframework/ios-arm64 -framework Mbedx509
*/
import "C"
