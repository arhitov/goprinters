package types

import "github.com/arhitov/goenum"

type OS = goenum.StringNamedMeta[osMap, osMeta]

const (
	OSUnknown OS = "unknown"
	OSLinux   OS = "linux"
	OSWindows OS = "windows"
	OSMacOS   OS = "macOS"
	OSIOS     OS = "iOS"
	OSAndroid OS = "android"
)

type osMap struct{}

type osMeta struct{}

func (l osMap) ValueMap() map[string]any {
	return map[string]any{
		string(OSUnknown): OSUnknown,
		string(OSLinux):   OSLinux,
		string(OSWindows): OSWindows,
		string(OSMacOS):   OSMacOS,
		string(OSIOS):     OSIOS,
		string(OSAndroid): OSAndroid,
	}
}

func (l osMap) NameMap() map[string]string {
	return map[string]string{
		string(OSUnknown): "Unknown",
		string(OSLinux):   "Linux",
		string(OSWindows): "Windows",
		string(OSMacOS):   "MacOS",
		string(OSIOS):     "iOS",
		string(OSAndroid): "Android",
	}
}

func (l osMap) MetaMap() map[string]osMeta {
	return map[string]osMeta{
		string(OSUnknown): {},
		string(OSLinux):   {},
		string(OSWindows): {},
		string(OSMacOS):   {},
		string(OSIOS):     {},
		string(OSAndroid): {},
	}
}
