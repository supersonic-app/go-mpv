//go:build ios && iossimulator && mpvmediakit

package mpv

// Uses libs from https://github.com/media-kit/libmpv-darwin-build

/*
#cgo CFLAGS: -I/opt/xcframework/Mpv.xcframework/ios-arm64_x86_64-simulator/Mpv.framework/Headers

#cgo LDFLAGS: -F/opt/xcframework/Mpv.xcframework/ios-arm64_x86_64-simulator -framework Mpv
#cgo LDFLAGS: -F/opt/xcframework/Avcodec.xcframework/ios-arm64_x86_64-simulator -framework Avcodec
#cgo LDFLAGS: -F/opt/xcframework/Avformat.xcframework/ios-arm64_x86_64-simulator -framework Avformat
#cgo LDFLAGS: -F/opt/xcframework/Avfilter.xcframework/ios-arm64_x86_64-simulator -framework Avfilter
#cgo LDFLAGS: -F/opt/xcframework/Avutil.xcframework/ios-arm64_x86_64-simulator -framework Avutil
#cgo LDFLAGS: -F/opt/xcframework/Swresample.xcframework/ios-arm64_x86_64-simulator -framework Swresample
#cgo LDFLAGS: -F/opt/xcframework/Swscale.xcframework/ios-arm64_x86_64-simulator -framework Swscale
#cgo LDFLAGS: -F/opt/xcframework/Mbedtls.xcframework/ios-arm64_x86_64-simulator -framework Mbedtls
#cgo LDFLAGS: -F/opt/xcframework/Mbedcrypto.xcframework/ios-arm64_x86_64-simulator -framework Mbedcrypto
#cgo LDFLAGS: -F/opt/xcframework/Mbedx509.xcframework/ios-arm64_x86_64-simulator -framework Mbedx509

*/
import "C"
