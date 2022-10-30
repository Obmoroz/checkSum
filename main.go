package main

/*
#cgo LDFLAGS: -ldl
//#include <dlfcn.h>
//typedef long long GoInt64;
//typedef unsigned long long GoUint64;
//typedef GoInt64 GoInt;
//typedef GoUint64 GoUint;
//typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;
static void callFromLib(void* p,_GoString_ HType,_GoString_ FilePath,_GoString_ Mail) {
    void (*fn)(_GoString_,_GoString_,_GoString_);
    *(void**)(&fn) =p ;
    fn(HType,FilePath,Mail);
}

*/
import "C"
import (
	"checkSum/config"
	"fmt"
	"os"
	"path"
	"time"
)

func main() {
	readFlags()
	if !isFlagsWasSet() {
		fmt.Println("Для работы приложения необходимо подать все аргументы\r\n" +
			"для получения справки введите флаг -help")
		os.Exit(1)
	}
	RootDir, err := config.RootDirIdentification()
	if err != nil {
		return
	}

	handle := C.dlopen(C.CString(path.Join(RootDir, "lib", "lib.so")), C.RTLD_LAZY)
	if handle == nil {
		fmt.Printf("error opening ./lib/lib.h")
		return
	}
	sdPidGetUnit := C.dlsym(handle, C.CString("MakeHash"))
	if sdPidGetUnit == nil {
		fmt.Printf("error resolving sd_pid_get_unit function")
		return
	}
	C.callFromLib(sdPidGetUnit, flagHashType, flagFilePath, flagMail)

	time.Sleep(10 * time.Second)

	fmt.Printf("RunLib is at %p\n", sdPidGetUnit)

}
