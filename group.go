package etcdpasswd

// #include <grp.h>
// #include <errno.h>
import "C"
import (
	"bytes"
	"syscall"
	"unsafe"
)


// export go_setgrent
func go_setgrent(stayopen C.int) nssStatus {
	// purpose: rewinds to the beginning of the group database, to allow repeated scans
	// nssStatusTryagain || nssStatusSuccess
	return nssStatusSuccess
}

// export go_endgrent
func go_endgrent() nssStatus {
	// purpose: used to close the group database after all processing has been performed
	return nssStatusSuccess
}

// export go_getgrent_r
func go_getgrent_r(grp *C.struct_group, buf *C.char, buflen C.size_t, errnop *C.int) nssStatus {
	// Get all group entries
	p, err := impl.Getgrent()
	if err == ErrNotFound {
		return nssStatusNotfound
	} else if err != nil {
		return nssStatusUnavail
	}

	return setCGroup(p, grp, buf, buflen, errnop)
}

// export go_getgrnam_r
func go_getgrnam_r(name string, grp *C.struct_group, buf *C.char, buflen C.size_t, errnop *C.int) nssStatus {
	// Get group entry by gid
	p, err := impl.Getgrnam(name)
	if err == ErrNotFound {
		return nssStatusNotfound
	} else if err != nil {
		return nssStatusUnavail
	}

	return setCGroup(p, grp, buf, buflen, errnop)
}

// export go_getgrgid_r
func go_getgrgid_r(gid GID, grp *C.struct_group, buf *C.char, buflen C.size_t, errnop *C.int) nssStatus {
	// get group entry by name
	p, err := impl.Getgrgid(gid)
	if err == ErrNotFound {
		return nssStatusNotfound
	} else if err != nil {
		return nssStatusUnavail
	}
	return setCGroup(p, grp, buf, buflen, errnop)
}




func setCPasswd(p *Passwd, grp *C.struct_group, buf *C.char, buflen C.size_t, errnop *C.int) nssStatus {
	if len(p.Name)+len(p.Passwd)+len(p.Gecos)+len(p.Dir)+len(p.Shell)+5 > int(buflen) {
		*errnop = C.int(syscall.EAGAIN)
		return nssStatusTryagain
	}

	gobuf := C.GoBytes(unsafe.Pointer(buf), C.int(buflen))
	b := bytes.NewBuffer(gobuf)
	b.Reset()

	grp.gr_name = (*C.char)(unsafe.Pointer(&gobuf[b.Len()]))
	b.WriteString(p.Name)
	b.WriteByte(0)

	//grp.gr_passwd = (*C.char)(unsafe.Pointer(&gobuf[b.Len()]))
	//b.WriteString(p.Passwd)
	//b.WriteByte(0)

	grp.gr_gid = C.uint(p.GID)
	
	// ?? how? https://github.com/zro/colony/blob/master/lib/nss/buffer.go#L32-L35
	//grp.gr_mem = stash.WriteStringSlice(colonyGrp.Members)

	return nssStatusSuccess
}
