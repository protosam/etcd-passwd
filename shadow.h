#include <nss.h>
#include <stddef.h>
#include <shadow.h>

enum nss_status _nss_etcd_setspent (void);
enum nss_status _nss_etcd_endspent (void);
enum nss_status _nss_etcd_getspent_r (struct spwd *result, char *buffer, size_t buflen, int *errnop);
enum nss_status _nss_etcd_getspnam_r (const char *name, struct spwd *result, char *buffer, size_t buflen, int *errnop);
