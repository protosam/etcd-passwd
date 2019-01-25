package etcdpasswd

// #include <shadow.h>
// #include <errno.h>
import "C"
import (
	"bytes"
	"syscall"
	"unsafe"
)


//export go_setspent 
func go_setspent() nssStatus {
	err := impl.Setpwent()
	if err != nil {
		return nssStatusUnavail
	}
	return nssStatusSuccess
}

//export go_endspent
func go_endspent() nssStatus {
	err := impl.Endpwent()
	if err != nil {
		return nssStatusUnavail
	}
	return nssStatusSuccess
}

//export go_getspent_r
func go_getspent_r(spwd *C.struct_spwd, buf *C.char, buflen C.size_t, errnop *C.int) nssStatus {
	// Get all shadow entries
	p, err := impl.Getspent()
	if err == ErrNotFound {
		return nssStatusNotfound
	} else if err != nil {
		return nssStatusUnavail
	}

	return setCSpwd(p, spwd, buf, buflen, errnop)
}

//export go_getspnam_r
func go_getspnam_r(name string, spwd *C.struct_spwd, buf *C.char, buflen C.size_t, errnop *C.int) nssStatus {
	// Get shadow entry by name
	p, err := impl.Getspnam(name)
	if err == ErrNotFound {
		return nssStatusNotfound
	} else if err != nil {
		return nssStatusUnavail
	}

	return setCSpwd(p, spwd, buf, buflen, errnop)
}

func setCSpwd(p *Passwd, spwd *C.struct_spwd, buf *C.char, buflen C.size_t, errnop *C.int) nssStatus {
	if len(p.Name)+len(p.Passwd)+len(p.Gecos)+len(p.Dir)+len(p.Shell)+5 > int(buflen) {
		*errnop = C.int(syscall.EAGAIN)
		return nssStatusTryagain
	}

	gobuf := C.GoBytes(unsafe.Pointer(buf), C.int(buflen))
	b := bytes.NewBuffer(gobuf)
	b.Reset()

	spwd.sp_namp = (*C.char)(unsafe.Pointer(&gobuf[b.Len()]))
	b.WriteString(p.Name)
	b.WriteByte(0)

	spwd.sp_lstchg = C.long(17920)
	spwd.sp_min = C.long(0)
	spwd.sp_max = C.long(99999)
	spwd.sp_warn = C.long(7)

	spwd.sp_pwdp = (*C.char)(unsafe.Pointer(&gobuf[b.Len()]))
	b.WriteString(p.Passwd)
	b.WriteByte(0)

	return nssStatusSuccess
}
