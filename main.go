package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

/*
#cgo LDFLAGS: -ldl
//#include <windows.h>
////#include <dlfcn.h>
////typedef long long GoInt64;
////typedef unsigned long long GoUint64;
////typedef GoInt64 GoInt;
////typedef GoUint64 GoUint;
////typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;
//static void callFromLib(void* p,_GoString_ HType,_GoString_ FilePath,_GoString_ Mail) {
//   void (*fn)(_GoString_,_GoString_,_GoString_);
//   *(void**)(&fn) =p ;
//   fn(HType,FilePath,Mail);
//}

*/
import "C"

func main() {
	readFlags()
	if !isFlagsWasSet() {
		fmt.Println("Для работы приложения необходимо подать все аргументы\r\n" +
			"для получения справки введите флаг -help")
		os.Exit(1)
	}
	//RootDir, err := config.RootDirIdentification()
	//if err != nil {
	//	return
	//}
	lib, err := syscall.LoadDLL("lib.dll")
	if err != nil {
		fmt.Printf(err.Error())
	}

	proc, err := lib.FindProc("MakeHash")
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("%+v\n", proc)

	fmt.Printf("%s %s %s\n", flagHashType, flagFilePath, flagMail)
	fmt.Printf("%v %v %v\n", uintptr(unsafe.Pointer(&flagHashType)), uintptr(unsafe.Pointer(&flagFilePath)), uintptr(unsafe.Pointer(&flagMail)))

	call, _, err := proc.Call(uintptr(unsafe.Pointer(&flagHashType)), uintptr(unsafe.Pointer(&flagFilePath)),
		uintptr(unsafe.Pointer(&flagMail)))
	if err != nil {
		return
	}

	if err != nil {
		fmt.Printf("%+v \r\n", err.Error())
	}

	fmt.Printf("%+v \n", call)

	//ret, _, callErr := syscall.SyscallN(proc,
	//	3,
	//	uintptr(unsafe.Pointer(&flagHashType)),
	//	uintptr(unsafe.Pointer(&flagFilePath)),
	//	uintptr(unsafe.Pointer(&flagMail)),
	//)
	//
	//if callErr != 0 {
	//	fmt.Printf("%+v \r\n", callErr.Error())
	//}
	//
	//fmt.Printf("%+v \n", ret)

	//fmt.Printf("%+v", RootDir+"\\lib\\lib.dll")
	//
	//handle := C.LoadLibrary(C.CString(RootDir + "\\lib\\lib.dll"))
	//fmt.Printf("%+v", handle)
	//if handle == nil {
	//	fmt.Printf("error opening ./lib/lib.h")
	//	return
	//}
	//fmt.Printf("%+v", handle)
	//
	//sdPidGetUnit := C.GetProcAddress(handle, C.CString("MakeHash"))
	//if sdPidGetUnit == nil {
	//	fmt.Printf("error resolving sd_pid_get_unit function")
	//	return
	//}
	//C.callFromLib(sdPidGetUnit, flagHashType, flagFilePath, flagMail)

	//	time.Sleep(10 * time.Second)

	//	fmt.Printf("RunLib is at %p\n", sdPidGetUnit)

}
