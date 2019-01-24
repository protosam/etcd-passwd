#include "nss.h"
#include "_cgo_export.h"

#include <string.h>
#include <grp.h>

// Thanks github.com/zro/colony for the boiler plate:
// https://github.com/zro/colony/blob/master/lib/nss/

enum nss_status _nss_etcd_setgrent(int stayopen) {
  return go_setgrent(stayopen);
}

enum nss_status _nss_etcd_endgrent(void) {
  return go_endgrent();
}

enum nss_status _nss_etcd_getgrent_r(struct group *result, char *buffer, size_t buflen, int *errnop) {
  return go_getgrent_r(result, buffer, buflen, errnop);
}

enum nss_status _nss_etcd_getgrgid_r(gid_t gid, struct group *grp, char *buffer, size_t buflen, int *errnop) {
  return go_getgrgid_r(gid, grp, buffer, buflen, errnop);
}

enum nss_status _nss_etcd_getgrnam_r(const char *name, struct group *grp, char *buffer, size_t buflen, int *errnop) {
  return go_getgrnam_r(name, grp, buffer, buflen, errnop);
}
