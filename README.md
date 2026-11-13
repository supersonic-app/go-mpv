# go-mpv

> Go bindings for [libmpv](https://mpv.io/).

This fork has diverged from upstream (**gen2brain/go-mpv**) - it will continue focusing on CGo and maintain support for `mpv_node` APIs to meet the needs of the [Supersonic](https://github.com/dweymouth/supersonic) and [STMPS](https://github.com/spezifisch/stmps) projects.

## iOS

To use on iOS you need mpv libraries:

 * [libmpv-ios/mpv-ios](https://github.com/libmpv-ios/mpv-ios): Download Libmpv.xcframework.zip
 * [media-kit/libmpv-darwin-build](https://github.com/media-kit/libmpv-darwin-build): Download a libmpv-xcframeworks archive for your needs.

Move the extracted xcframeworks as subdirs to `/opt/xcframework/`.

For *media-kit/libmpv-darwin-build* add the tag `mpvmediakit`, e.g.:

```txt
go build -tags mpvmediakit
fyne package -os ios -tags mpvmediakit
```

Furthermore for *media-kit/libmpv-darwin-build* you need to bundle the frameworks. Makefile example:

```Makefile
IOS_XCFRAMEWORKS := Avcodec Avfilter Avformat Avutil Mbedcrypto Mbedtls Mbedx509 Mpv Swresample Swscale

.PHONY: bundledeps_ios
bundledeps_ios: package/ios/MyMediaPlayer.app
	mkdir -p package/ios/MyMediaPlayer.app/Frameworks
	for name in $(IOS_XCFRAMEWORKS); do \
		src="/opt/xcframework/$$name.xcframework/ios-arm64/$$name.framework"; \
		cp -a "$$src" "package/ios/MyMediaPlayer.app/Frameworks/"; \
	done
	install_name_tool -add_rpath '@executable_path/Frameworks' package/ios/MyMediaPlayer.app/main
	codesign --force --deep --sign - --timestamp=none package/ios/MyMediaPlayer.app
```

### Simulator

To build for iOS Simulator add the tag `iossimulator`, e.g.:

```txt
go build -tags iossimulator
fyne package -os iossimulator -tags iossimulator
```


