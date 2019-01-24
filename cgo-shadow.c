#include "nss.h"
#include "_cgo_export.h"

#include <string.h>
#include <shadow.h>



enum nss_status _nss_etcd_setspent (void){
	return go_setspent();
}

enum nss_status _nss_etcd_endspent (void){
	return go_endspent();
}

enum nss_status _nss_etcd_getspent_r (struct spwd *result, char *buffer, size_t buflen, int *errnop){
	return go_getspent_r(result, buffer, buflen, errnop);
}

enum nss_status _nss_etcd_getspnam_r (const char *name, struct spwd *result, char *buffer, size_t buflen, int *errnop){
	GoString goname = {name, strlen(name) };
	return go_getspnam_r(goname, spwd, buffer, buflen, errnop)
}
