#!/bin/bash
#
# android-build.sh allows building a MIQT Go application for Android.
# For details, see the top-level README.md file.

set -Eeuo pipefail

# QT_PATH is pre-set in the Qt 5 docker container environment. Includes trailing slash
# QT_ANDROID is pre-set in the Qt 6 docker container environment
QT_PATH=${QT_PATH:-/usr/local/Qt-5.15.13/}
QT_ANDROID=${QT_ANDROID:-$QT_PATH}

export LC_ALL=C.UTF-8
	
# get_app_name returns the android app's name. This affects the default name
# in deployment-settings.json and the generated lib.so names.
# You can still customise the package name and the package ID in the xml
# files after generation.
get_app_name() {
	basename "$(pwd)"
}

get_stub_soname() {
	# libRealAppName_arm64-v8a.so
	echo "lib$(get_app_name)_arm64-v8a.so"
}

get_go_soname() {
	echo "libMiqtGolangApp_arm64-v8a.so"
}

ndk_version() {
	ls /opt/android-sdk/ndk/ | tail -n1
}

target_sdk_version() {
	ls /opt/android-sdk/platforms | tail -n1 | sed -re 's/android-//'
}

build_tools_version() {
	ls /opt/android-sdk/build-tools | tail -n1
}

# extra_libs returns a comma-separated list of extra libraries to include in
# the apk package
extra_libs() {
	if [[ -d /opt/android_openssl ]] ; then
		# Our miqt Qt5 container includes these extra .so libraries
		# However, the aqtinstall-based Qt 6 container does not use them
		echo "/opt/android_openssl/ssl_1.1/arm64-v8a/libssl_1_1.so,/opt/android_openssl/ssl_1.1/arm64-v8a/libcrypto_1_1.so"
	fi
}

# generate_template_contents produces a deployment settings JSON file that is
# understood by the androiddeployqt program.
# Available fields are documented in the template file at:
# @ref /usr/local/Qt-5.15.13/mkspecs/features/android/android_deployment_settings.prf
generate_template_contents() {
	
	cat <<EOF
{
	"_description": "Generated by miqt/android-mktemplate",
	"application-binary": "$(get_app_name)",
	"architectures": {
		"arm64-v8a" : "aarch64-linux-android"
	},
	"android-extra-libs": "$(extra_libs)",
	"android-min-sdk-version": "23",
	"android-target-sdk-version": "$(target_sdk_version)",
	"ndk": "/opt/android-sdk/ndk/$(ndk_version)",
	"ndk-host": "linux-x86_64",
	"qt": "${QT_ANDROID}",
	"sdk": "/opt/android-sdk",
	"sdkBuildToolsRevision": "$(build_tools_version)",
	"stdcpp-path": "/opt/android-sdk/ndk/$(ndk_version)/toolchains/llvm/prebuilt/linux-x86_64/sysroot/usr/lib/",
	"tool-prefix": "llvm",
	"toolchain-prefix": "llvm",
	"useLLVM": true
}
EOF
}

# pkg_config_dependencies echoes the pkg-config packages that are needed as
# a baseline for the detected Qt version.
# It's not smart enough to discover other Qt packages yet.
pkg_config_dependencies() {
	local QT_VERSION=$(detect_miqt_qt_version)	
	if [[ $QT_VERSION == 'qt5' ]] ; then
		echo Qt5Widgets
		echo Qt5AndroidExtras
		
	elif [[ $QT_VERSION == 'qt6' ]] ; then
		# There is no AndroidExtras in Qt 6
		echo Qt6Widgets 
		
	fi
}

# android_stub_gen generates a .so that runs an exported function from another
# so file. This is because Qt for Android always tries to run `main`, but Go's
# c-shared build mode cannot export a function named `main`.
android_stub_gen() {

	local STUBNAME="miqt-stub-$(date +%s).cpp"

	cat > $STUBNAME <<EOF
#include <android/log.h>
#include <dlfcn.h>
#include <stdlib.h>

typedef void goMainFunc_t();  

int main(int argc, char** argv) {
    __android_log_print(ANDROID_LOG_VERBOSE, "miqt_stub", "Starting up");
        
    void* handle = dlopen("$(get_go_soname)", RTLD_LAZY);
    if (handle == NULL) {
        __android_log_print(ANDROID_LOG_VERBOSE, "miqt_stub", "miqt_stub: null handle opening so: %s", dlerror());
        exit(1);
    }
    
    void* goMain = dlsym(handle, "AndroidMain");
    if (goMain == NULL) {
        __android_log_print(ANDROID_LOG_VERBOSE, "miqt_stub", "miqt_stub: null handle looking for function: %s", dlerror());
        exit(1);        
    }
    
    __android_log_print(ANDROID_LOG_VERBOSE, "miqt_stub", "miqt_stub: Found target, calling");
    
    // Cast to function pointer and call
    goMainFunc_t* f = (goMainFunc_t*)goMain;
    f();
    
    __android_log_print(ANDROID_LOG_VERBOSE, "miqt_stub", "miqt_stub: Target function returned");
    return 0;
}

EOF
	
	# Compile
	# Link with Qt libraries so that androiddeployqt detects us as being the
	# main shared library
	
	$CXX -shared \
		-ldl \
		-llog \
		-L${QT_ANDROID}/plugins/platforms -lplugins_platforms_qtforandroid_arm64-v8a \
		$(pkg-config --libs $(pkg_config_dependencies)) \
		${STUBNAME} \
		"-Wl,-soname,$(basename "$(get_stub_soname)")" \
		-o "android-build/libs/arm64-v8a/$(get_stub_soname)"
    
	rm "${STUBNAME}"
}

# require_is_main_package verifies that this is the main Go package.
require_is_main_package() {	
	if ! grep -Fq 'package main' *.go ; then
		echo "This doesn't seem to be the main package" >&2
		exit 1
	fi
}

# patch_app_main sets up the startup go files so that the go program can be
# built either as c-shared for Android or as a normal program for desktop OSes.
patch_app_main() {
	
	# Replace func main() with app_main()
		
	for srcfile in *.go ; do
		sed -i -re 's/^func main\(\) \{/func app_main() {/' "$srcfile"
	done
	
	# Add shim startup files
	
	cat <<EOF > startup_android.go
//go:build android
// +build android

package main

import "C" // Required for export support

//export AndroidMain
func AndroidMain() {
	app_main()
}

func main() {
	// Must be empty
}

EOF

	cat <<EOF > startup_other.go
//go:build !android
// +build !android

package main

func main() {
	app_main()
}

EOF

	gofmt -w startup_android.go || true
	gofmt -w startup_other.go || true
	
	# Done
}

# unpatch_app_main undoes the transformation from patch_app_main.
unpatch_app_main() {
		
	# Replace func main() with app_main()
		
	for srcfile in *.go ; do
		sed -i -re 's/^func app_main\(\) \{/func main() {/' "$srcfile"
	done

	# Remove extra startup files

	rm startup_android.go || true 
	rm startup_other.go || true
}

# build_app_so compiles the Go app as a c-shared .so.
build_app_so() {	
	go build \
		-buildmode c-shared \
		-ldflags "-s -w -extldflags -Wl,-soname,$(get_go_soname)" \
		-o "android-build/libs/arm64-v8a/$(get_go_soname)"
}

sdkmanager() {
	echo /opt/android-sdk/cmdline-tools/*/bin/sdkmanager
}

# build_apk calls androiddeployqt to package the android-build directory into
# the final apk.
build_apk() {

	# Qt 6 androiddeployqt: Understands the QT_ANDROID_KEYSTORE_STORE_PASS in env
	# Qt 5 androiddeployqt: Doesn't - any use of `--sign keystore alias` here
	#  requires stdin prompt but doesn't pass androiddeployqt's stdin through
	#  to jarsigner subprocess
	# Either way, don't sign the app here, rely on separate jarsigner command

	# Work around an issue with Qt 6 sporadically failing to detect that
	# we have a valid Qt platform plugin
	# TODO why does this happen? Is it related to file sort ordering?
	# It really does fix itself after multiple attempts (usually less than 5)
	# - When it fails: Error: qmlimportscanner not found at /qmlimportscanner
	# - When it works: Error: qmlimportscanner not found at libexec/qmlimportscanner
	while ! androiddeployqt \
		--input ./deployment-settings.json \
		--output ./android-build/ \
		--release \
		2> >(tee /tmp/androiddeployqt.stderr.log >&2) ; do
		
		if grep -Fq 'Make sure the app links to Qt Gui library' /tmp/androiddeployqt.stderr.log ; then
			echo "Detected temporary problem with Qt plugin detection. Retrying..."
			sleep 1
			
		else
			# real error
			exit 1
		fi
	done
	
	local OUTAPK=$(get_app_name).apk
	rm "$OUTAPK" || true
	
	# Zipalign
	echo "Zipalign..."
	/opt/android-sdk/build-tools/*/zipalign \
		-p 4 \
		./android-build/build/outputs/apk/release/android-build-release-unsigned.apk \
		"$OUTAPK"
	
	# Sign
	echo "Signing..."
	#jarsigner \
	#	-verbose \
	#	-sigalg SHA256withRSA -digestalg SHA256 \
	#	-keystore ./android.keystore \
	#	"$OUTAPK" \
	#	"${QT_ANDROID_KEYSTORE_ALIAS}" \
	#	-storepass:env QT_ANDROID_KEYSTORE_STORE_PASS \
	#	-keypass:env QT_ANDROID_KEYSTORE_KEY_PASS
	
	/opt/android-sdk/build-tools/*/apksigner \
		sign \
		--ks ./android.keystore \
		--ks-key-alias "${QT_ANDROID_KEYSTORE_ALIAS}" \
		--ks-pass env:QT_ANDROID_KEYSTORE_STORE_PASS \
		--key-pass env:QT_ANDROID_KEYSTORE_KEY_PASS \
		"$OUTAPK"
}

# detect_env_qt_version detects the system's current Qt version.
detect_env_qt_version() {	
	if qmake --version | fgrep -q 'Qt version 5' ; then
		echo "qt5"
		return 0
	
	elif qmake --version | fgrep -q 'Qt version 6' ; then
		echo "qt6"
		return 0
		
	else 
		echo "Missing Qt tools in PATH" >&2
		exit 1
	fi
}

# detect_miqt_qt_version echoes either "qt5", "qt6", or exits bash.
detect_miqt_qt_version() {
	local IS_QT5=false
	if grep -qF '"github.com/e1z0/miqt/qt"' *.go ; then
		IS_QT5=true
	fi
	
	local IS_QT6=false
	if grep -qF '"github.com/e1z0/miqt/qt6"' *.go ; then
		IS_QT6=true
	fi
	
	if [[ $IS_QT5 == true && $IS_QT6 == true ]] ; then
		echo "Found qt5 and qt6 imports, confused about what to do next" >&2
		exit 1
		
	elif [[ $IS_QT5 == true ]] ; then
		echo "qt5"
		return 0
		
	elif [[ $IS_QT6 == true ]] ; then
		echo "qt6"
		return 0
		
	else
		echo "Found neither qt5 nor qt6 imports. Is this a MIQT Qt app?" >&2
		exit 1
	fi
}

generate_default_keystore() {
	
	local GENPASS=storepass_$(cat /dev/urandom | head -c64 | md5sum | cut -d' ' -f1)
	
	keytool \
		-genkeypair \
		-dname "cn=Miqt, ou=Miqt, o=Miqt, c=US" \
		-keyalg RSA \
		-alias miqt \
		-keypass "${GENPASS}" \
		-keystore ./android.keystore \
		-storepass "${GENPASS}" \
		-validity 20000

	echo "QT_ANDROID_KEYSTORE_PATH=./android.keystore" > android.keystore.env
	echo "QT_ANDROID_KEYSTORE_ALIAS=miqt" >> android.keystore.env
	echo "QT_ANDROID_KEYSTORE_STORE_PASS=${GENPASS}" >> android.keystore.env
	echo "QT_ANDROID_KEYSTORE_KEY_PASS=${GENPASS}" >> android.keystore.env
}

# main is the entrypoint for android-build.sh.
main() {
	
	if [[ $(detect_env_qt_version) != $(detect_miqt_qt_version) ]] ; then
		echo "The system is $(detect_env_qt_version) but the app uses $(detect_miqt_qt_version). Is this the right container?" >&2
		exit 1
	fi
	
	require_is_main_package
	
	# Rebuild deployment-settings.json
	if [[ ! -f deployment-settings.json ]] ; then
		echo "Generating deployment-settings.json..."
		generate_template_contents > deployment-settings.json
	fi
	
	mkdir -p android-build/libs/arm64-v8a
	
	if [[ ! -f android-build/libs/arm64-v8a/$(get_stub_soname) ]] ; then
		echo "Generating stub so..."
		android_stub_gen
	fi
	
	# Rebuild miqt_golang_app.so
	echo "Compiling Go app..."
	patch_app_main
	build_app_so
	unpatch_app_main
	
	# Keypair
	if [[ ! -f android.keystore || ! -f android.keystore.env ]] ; then
		echo "Signing keystore not found, generating a default one..."
		generate_default_keystore
	fi
	
	# Load keypair credentials into exported env vars
	set -o allexport
	source android.keystore.env
	set +o allexport
	
	# Generate .apk
	echo "Packaging APK..."
	build_apk
	
	echo "Complete"
}

main "$@"
