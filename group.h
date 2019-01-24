#include <nss.h>
#include <stddef.h>
#include <grp.h>

extern enum nss_status _nss_etcd_setgrent(int stayopen);
extern enum nss_status _nss_etcd_endgrent(void);
extern enum nss_status _nss_etcd_getgrent_r(struct group *result, char *buffer, size_t buflen, int *errnop);
extern enum nss_status _nss_etcd_getgrgid_r(gid_t gid, struct group *grp, char *buffer, size_t buflen, int *errnop);
extern enum nss_status _nss_etcd_getgrnam_r(const char *name, struct group *grp, char *buffer, size_t buflen, int *errnop);
